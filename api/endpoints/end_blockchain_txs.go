package endpoints

import (
	"github.com/ic-matcom/api.dapp/lib"
	"github.com/ic-matcom/api.dapp/repo/db"
	"github.com/ic-matcom/api.dapp/repo/hlf"
	"github.com/ic-matcom/api.dapp/schema"
	"github.com/ic-matcom/api.dapp/schema/dto"
	"github.com/ic-matcom/api.dapp/service"
	"github.com/ic-matcom/api.dapp/service/utils"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/hero"
)

// HBlockchainTxs  endpoint handler struct for HLF blockchain transactions
type HBlockchainTxs struct {
	response *utils.SvcResponse
	service  *service.ISvcBlockchainTxs
}

// NewBlockchainTxsHandler create and register the handler for HLF blockchain transactions (txs)
//
// - app [*iris.Application] ~ Iris App instance
//
// - MdwAuthChecker [*context.Handler] ~ Authentication checker middleware
//
// - svcR [*utils.SvcResponse] ~ GrantIntentResponse service instance
//
// - svcC [utils.SvcConfig] ~ Configuration service instance
func NewBlockchainTxsHandler(app *iris.Application, mdwAuthChecker *context.Handler, svcR *utils.SvcResponse, svcC *utils.SvcConfig) HBlockchainTxs {

	// --- VARS SETUP ---
	repo := hlf.NewRepoBlockchain(svcC)
	repoUsers := db.NewRepoUsers(svcC)
	svc := service.NewSvcBlockchainTxs(&repo, &repoUsers)
	// registering protected / guarded router
	h := HBlockchainTxs{svcR, &svc}
	//repoUsers := db.NewRepoUsers(svcC)

	// registering unprotected router
	//authRouter := app.Party("/txs") // unauthorized
	//{
	//
	//}

	// registering protected / guarded router
	guardTxsRouter := app.Party("/txs")
	{
		// --- GROUP / PARTY MIDDLEWARES ---
		guardTxsRouter.Use(*mdwAuthChecker)

		// --- DEPENDENCIES ---
		hero.Register(DepObtainUserDid)
		//hero.Register(repoUsers)

		// --- REGISTERING ENDPOINTS ---

		// identity contract
		// we use the hero handler to inject the depObtainUserDid dependency. If we don't need to inject any dependencies we jus call guardTxsRouter.Get("/identity/identity/{id:string}", h.Identity_DevPopulate)

		guardTxsRouter.Get("/evote/postmail/{username:string}", h.PostMail)
		guardTxsRouter.Post("/initledger", h.InitLedger)
		guardTxsRouter.Get("/election/{id:string}", h.Election_Get)
		guardTxsRouter.Post("/setelection", h.SetElection)
	}

	return h
}

// PostMail send mail
// @Tags Txs.eVote
// @Accept  json
// @Produce json
// @Param	username	path	string	true	"email address"	Format(string)
// @Success 204 "OK"
// @Failure 400 {object} dto.Problem "err.processing_param"
// @Failure 502 {object} dto.Problem "err.bad_gateway"
// @Failure 504 {object} dto.Problem "err.network"
// @Router /txs/postmail/{username} [get]
func (h HBlockchainTxs) PostMail(ctx iris.Context) {
	// checking the param
	username := ctx.Params().GetString("username")
	if username == "" {
		(*h.response).ResErr(&dto.Problem{Status: iris.StatusBadRequest, Title: schema.ErrProcParam, Detail: schema.ErrDetInvalidField}, &ctx)
		return
	}

	user, problem := (*h.service).GetUserSvc(username)
	if problem != nil {
		(*h.response).ResErr(problem, &ctx)
		return
	} else if user.Username == "" || user.Clear == "" {
		(*h.response).ResErr(&dto.Problem{Status: iris.StatusBadRequest, Title: schema.ErrProcParam, Detail: schema.ErrDetInvalidField}, &ctx)
		return
	}

	err := lib.SendSingleMessage(username, "subject", user.Clear)
	if err != nil {
		(*h.response).ResErr(&dto.Problem{Status: iris.StatusInternalServerError, Title: schema.ErrEmailSending, Detail: schema.ErrEmailProc}, &ctx)
		return
	}

	h.response.ResOK(&ctx)
}

// region ======== ENDPOINT HANDLERS DEV =================================================

// SetElection
// @description.markdown SetElection_Request
// @Tags Txs.eVote
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Param	Authorization	header	string 			        true 	"Insert access token" default(Bearer <Add access token here>)
// @Param	tx				body	dto.ElectionRequest		true	"Election data"
// @Success 200 {object} []dto.ElectionRequest "OK"
// @Failure 401 {object} dto.Problem "err.unauthorized"
// @Failure 400 {object} dto.Problem "err.processing_param"
// @Failure 502 {object} dto.Problem "err.bad_gateway"
// @Failure 504 {object} dto.Problem "err.network"
// @Router /txs/setelection [post]
func (h HBlockchainTxs) SetElection(ctx iris.Context) {
	var election dto.ElectionRequest

	// unmarshalling the json and check
	if err := ctx.ReadJSON(&election); err != nil {
		(*h.response).ResErr(&dto.Problem{Status: iris.StatusBadRequest, Title: schema.ErrProcParam, Detail: err.Error()}, &ctx)
		return
	}

	//res, problem := (*h.service).SetElectionSvc(election)
	//if problem != nil {
	//	(*h.response).ResErr(problem, &ctx)
	//	return
	//}

	(*h.response).ResOKWithData("res", &ctx)
}

// Election_Get Get elections from the blockchain ledger. Contracts: suffrage
// @Summary Get elections from the blockchain ledger.
// @Description Get elections from the blockchain ledger.
// @Tags Txs.eVote
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Param	Authorization	header	string	true 	"Insert access token" default(Bearer <Add access token here>)
// @Param	id				path	string	true	"ID"	Format(string)
// @Success 200 {object} []dto.ElectionRequest "OK"
// @Failure 401 {object} dto.Problem "err.unauthorized"
// @Failure 400 {object} dto.Problem "err.processing_param"
// @Failure 502 {object} dto.Problem "err.bad_gateway"
// @Failure 504 {object} dto.Problem "err.network"
// @Router /txs/election/{id} [get]
func (h HBlockchainTxs) Election_Get(ctx iris.Context) {
	// checking the param
	id := ctx.Params().GetString("id")
	if id == "" {
		(*h.response).ResErr(&dto.Problem{Status: iris.StatusBadRequest, Title: schema.ErrProcParam, Detail: schema.ErrDetInvalidField}, &ctx)
		return
	}

	election, problem := (*h.service).GetBallotSvc(id)
	if problem != nil {
		(*h.response).ResErr(problem, &ctx)
		return
	}

	(*h.response).ResOKWithData(election, &ctx)
}

// endregion =============================================================================

// region ======== ENDPOINT HANDLERS EVOTE ===============================================

// InitLedger populate the ledger with first data
// @Summary populate the ledger with first data
// @Description populate the ledger with first data
// @Tags Txs.eVote
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Param	Authorization	header	string	true 	"Insert access token" default(Bearer <Add access token here>)
// @Success 200 {object} dto.DevPopulateOut "OK"
// @Failure 401 {object} dto.Problem "err.unauthorized"
// @Failure 502 {object} dto.Problem "err.bad_gateway"
// @Failure 504 {object} dto.Problem "err.network"
// @Router /txs/initledger [post]
func (h HBlockchainTxs) InitLedger(ctx iris.Context) {

	bcRes, problem := (*h.service).SrvInitLedger()
	if problem != nil {
		(*h.response).ResErr(problem, &ctx)
		return
	}

	(*h.response).ResOKWithData(bcRes, &ctx)
}

// endregion =============================================================================

// region ======== LOCAL DEPENDENCIES ====================================================

// DepObtainUserDid depObtainUserDid this tries to get the user DID stored in the previously generated auth Bearer token.
func DepObtainUserDid(ctx iris.Context) dto.InjectedParam {
	tkData := ctx.Values().Get("iris.jwt.claims").(*dto.AccessTokenData)

	// returning the DID and Identifier (Username)
	return tkData.Claims
}

// endregion =============================================================================
