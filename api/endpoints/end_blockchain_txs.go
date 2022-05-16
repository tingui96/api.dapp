package endpoints

import (
	"fmt"
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
		
		guardTxsRouter.Post("/initledger", h.InitLedger)
		guardTxsRouter.Get("/optativo/{id:string}", h.Asset_Get)
		guardTxsRouter.Post("/optativo/set", h.Asset_Set)
	}

	return h
}

// region ======== ENDPOINT HANDLERS DEV =================================================

// Asset_Set
// @description.markdown SetElection_Request
// @Tags Txs.eVote
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Param	Authorization	header	string 			        true 	"Insert access token" default(Bearer <Add access token here>)
// @Param	tx				body	dto.TestRequest		true	"Test data"
// @Success 200 {object} []dto.TestRequest "OK"
// @Failure 401 {object} dto.Problem "err.unauthorized"
// @Failure 400 {object} dto.Problem "err.processing_param"
// @Failure 502 {object} dto.Problem "err.bad_gateway"
// @Failure 504 {object} dto.Problem "err.network"
// @Router /txs/optativo/set [post]
func (h HBlockchainTxs) Asset_Set(ctx iris.Context) {
	var request dto.TestRequest

	// unmarshalling the json and check
	if err := ctx.ReadJSON(&request); err != nil {
		(*h.response).ResErr(&dto.Problem{Status: iris.StatusBadRequest, Title: schema.ErrProcParam, Detail: err.Error()}, &ctx)
		return
	}

	//res, problem := (*h.service).SetAssetSvc(request)
	//if problem != nil {
	//	(*h.response).ResErr(problem, &ctx)
	//	return
	//}

	(*h.response).ResOKWithData("res", &ctx)
}

// Asset_Get Get asset from the blockchain ledger. Contracts: mycc
// @Summary Get asset from the blockchain ledger.
// @Description Get asset from the blockchain ledger.
// @Tags Txs.eVote
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Param	Authorization	header	string	true 	"Insert access token" default(Bearer <Add access token here>)
// @Param	id				path	string	true	"ID"	Format(string)
// @Success 200 {object} byte "OK"
// @Failure 401 {object} dto.Problem "err.unauthorized"
// @Failure 400 {object} dto.Problem "err.processing_param"
// @Failure 502 {object} dto.Problem "err.bad_gateway"
// @Failure 504 {object} dto.Problem "err.network"
// @Router /txs/optativo/{id} [get]
func (h HBlockchainTxs) Asset_Get(ctx iris.Context) {
	// checking the param
	id := ctx.Params().GetString("id")
	if id == "" {
		(*h.response).ResErr(&dto.Problem{Status: iris.StatusBadRequest, Title: schema.ErrProcParam, Detail: schema.ErrDetInvalidField}, &ctx)
		return
	}

	result, problem := (*h.service).GetAssetSvc(id)
	if problem != nil {
		(*h.response).ResErr(problem, &ctx)
		return
	}
	fmt.Println("--> 1", result)

	(*h.response).ResOKWithData(result, &ctx)
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
