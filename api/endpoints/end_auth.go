package endpoints

import (
	"github.com/ic-matcom/api.dapp/lib"
	"github.com/ic-matcom/api.dapp/repo/db"
	"github.com/ic-matcom/api.dapp/repo/hlf"
	"github.com/ic-matcom/api.dapp/schema"
	"github.com/ic-matcom/api.dapp/schema/dto"
	"github.com/ic-matcom/api.dapp/schema/mapper"
	"github.com/ic-matcom/api.dapp/service"
	"github.com/ic-matcom/api.dapp/service/auth"
	"github.com/ic-matcom/api.dapp/service/utils"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/hero"
)

type HAuth struct {
	response  *utils.SvcResponse
	appConf   *utils.SvcConfig
	//providers map[string]bool
}

// NewAuthHandler create and register the authentication handlers for the App. For the moment, all the
// auth handlers emulates the Oauth2 "password" grant-type using the "client-credentials" flow.
//
// - app [*iris.Application] ~ Iris App instance
//
// - MdwAuthChecker [*context.Handler] ~ Authentication checker middleware
//
// - svcR [*utils.SvcResponse] ~ GrantIntentResponse service instance
//
// - svcC [utils.SvcConfig] ~ Configuration service instance
func NewAuthHandler(app *iris.Application, MdwAuthChecker *context.Handler, svcR *utils.SvcResponse, svcC *utils.SvcConfig) HAuth {

	// --- VARS SETUP ---
	h := HAuth{svcR, svcC} //, make(map[string]bool)}
	// filling providers
	// h.providers["default"] = true
	// h.providers["another_provider"] = true

	repoHlfIdentity := hlf.NewRepoIdentity(svcC)
	repoUsers := db.NewRepoUsers(svcC)


	svcIdentity := service.NewSvcHlfIdentity(&repoHlfIdentity, &repoUsers) // instantiating HLF identity Service
	svcAuth := auth.NewSvcAuthentication(svcC, &repoUsers)          // instantiating authentication Service

	// registering unprotected router
	authRouter := app.Party("/auth") // authorize
	{
		// --- GROUP / PARTY MIDDLEWARES ---

		// --- DEPENDENCIES ---
		hero.Register(depObtainUserCred)
		hero.Register(svcAuth) // as an alternative, we can put this dependencies as property in the HAuth struct, as we are doing in the rest of the endpoints / handlers


		// --- REGISTERING ENDPOINTS ---
		authRouter.Post("/", hero.Handler(h.authIntent))
	}

	// registering protected router
	guardAuthRouter := app.Party("/auth")
	{
		// --- GROUP / PARTY MIDDLEWARES ---
		guardAuthRouter.Use(*MdwAuthChecker) // registering access token checker middleware

		// --- DEPENDENCIES ---
		hero.Register(svcIdentity)
		hero.Register(DepObtainUserDid)
		hero.Register(repoUsers)

		// --- REGISTERING ENDPOINTS ---
		guardAuthRouter.Get("/logout", h.logout)
		guardAuthRouter.Get("/user", hero.Handler(h.userGet))

		// guardAuthRouter.Post("/gencerts", hero.Handler(h.GenUsersCerts))
	}

	return h
}

// region ======== ENDPOINT HANDLERS =====================================================

// authIntent Intent to grant authentication using the provider user's credentials
// @Summary Auth the user
// @Description Intent to grant authentication using the provider user's credentials
// @Tags Auth
// @Accept multipart/form-data
// @Produce json
// @Param 	credential 	body 	dto.UserCredIn 	true	"User Login Credential"
// @Success 200 "OK"
// @Failure 401 {object} dto.Problem "err.unauthorized"
// @Failure 400 {object} dto.Problem "err.wrong_auth_provider"
// @Failure 504 {object} dto.Problem "err.network"
// @Failure 500 {object} dto.Problem "err.json_parse"
// @Router /auth [post]
func (h HAuth) authIntent(ctx iris.Context, uCred *dto.UserCredIn, svcAuth *auth.SvcAuthentication) {

	provider := "default"

	authGrantedData, problem := svcAuth.AuthProviders[provider].GrantIntent(uCred, nil) // requesting authorization to (provider) mechanisms in this case
	if problem != nil {                                                                 // check for errors
		(*h.response).ResErr(problem, &ctx)
		return
	}

	// TODO pass this to the service
	// if so far so good, we are going to create the auth token
	tokenData := mapper.ToAccessTokenDataV(authGrantedData)
	accessToken, err := lib.MkAccessToken(tokenData, []byte(h.appConf.JWTSignKey), h.appConf.TkMaxAge)
	if err != nil {
		(*h.response).ResErr(&dto.Problem{Status: iris.StatusInternalServerError, Title: schema.ErrJwtGen, Detail: err.Error()}, &ctx)
		return
	}

	(*h.response).ResOKWithData(string(accessToken), &ctx)
}

// logout this endpoint invalidated a previously granted access token
// @Description This endpoint invalidated a previously granted access token
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert access token" default(Bearer <Add access token here>)
// @Tags Auth
// @Produce  json
// @Success 204 "OK"
// @Failure 401 {object} dto.Problem "err.unauthorized"
// @Failure 500 {object} dto.Problem "err.generic
// @Router /auth/logout [get]
func (h HAuth) logout(ctx iris.Context) {

	err := ctx.Logout()

	if err != nil {
		(*h.response).ResErr(&dto.Problem{Status: iris.StatusInternalServerError, Title: schema.ErrGeneric, Detail: err.Error()}, &ctx)
		return
	}

	// so far so good
	(*h.response).ResOK(&ctx)
}

// userGet Get user from the BD.
// @Summary Retorna el usuario autenticado
// @Description Retorna informacion del usuario autenticado en el sistema
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert access token" default(Bearer <Add access token here>)
// @Tags Auth
// @Produce  json
// @Success 200 {object} []dto.GetUser "OK"
// @Failure 401 {object} dto.Problem "err.unauthorized"
// @Failure 500 {object} dto.Problem "err.generic
// @Router /auth/user [get]
func (h HAuth) userGet(ctx iris.Context, params dto.InjectedParam, r db.RepoUsers) {
	user, err := r.GetUser(params.Did)
	if err != nil {
		(*h.response).ResErr(dto.NewProblem(iris.StatusInternalServerError, schema.ErrBuntdb, err.Error()), &ctx)
		return
	}
	user.Clear = ""
	l := map[string]*dto.User{"user": user}
	(*h.response).ResOKWithData(l, &ctx)
}

// endregion =============================================================================

// region ======== LOCAL DEPENDENCIES ====================================================

// depObtainUserCred is used as dependencies to obtain / create the user credential from request body (multipart/form-data).
// It return a dto.UserCredIn struct
func depObtainUserCred(ctx iris.Context) dto.UserCredIn {
	cred := dto.UserCredIn{}

	// Getting data
	cred.Username = ctx.PostValue("username")
	cred.Password = ctx.PostValue("password")

	// TIP: We can do some validation here if we want
	return cred
}

// endregion =============================================================================