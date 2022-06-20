package endpoints

import (
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

		guardTxsRouter.Post("/init_ledger", h.InitLedger)
		guardTxsRouter.Get("/read_asset/{id:string}", h.ReadAsset)
		guardTxsRouter.Patch("/update_asset", h.UpdateAsset)
		guardTxsRouter.Post("/create_asset", h.CreateAsset)
	}

	return h
}

// InitLedger populate the ledger with first data
// @Summary populate the ledger with first data
// @Description populate the ledger with first data
// @Tags Txs.mycc
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Param	Authorization	header	string	true 	"Insert access token" default(Bearer <Add access token here>)
// @Success 200 {object} dto.DevPopulateOut "OK"
// @Failure 401 {object} dto.Problem "err.unauthorized"
// @Failure 502 {object} dto.Problem "err.bad_gateway"
// @Failure 504 {object} dto.Problem "err.network"
// @Router /txs/init_ledger [post]
func (h HBlockchainTxs) InitLedger(ctx iris.Context) {

	bcRes, problem := (*h.service).SrvInitLedger()
	if problem != nil {
		(*h.response).ResErr(problem, &ctx)
		return
	}

	(*h.response).ResOKWithData(bcRes, &ctx)
}

// ReadAsset Get asset from the blockchain ledger. Contracts: mycc
// @Summary Get asset from the blockchain ledger.
// @description.markdown ReadAsset_Request
// @Tags Txs.mycc
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
// @Router /txs/read_asset/{id} [get]
func (h HBlockchainTxs) ReadAsset(ctx iris.Context) {
	// checking the param
	id := ctx.Params().GetString("id")
	if id == "" {
		(*h.response).ResErr(&dto.Problem{Status: iris.StatusBadRequest, Title: schema.ErrProcParam, Detail: schema.ErrDetInvalidField}, &ctx)
		return
	}

	result, problem := (*h.service).ReadAssetSvc(id)
	if problem != nil {
		(*h.response).ResErr(problem, &ctx)
		return
	}

	(*h.response).ResOKWithData(result, &ctx)
}

// UpdateAsset updates an existing asset in the world state with provided parameters.
// @Summary updates an existing asset in the world state with provided parameters.
// @Tags Txs.mycc
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Param	Authorization	header	string	true 	"Insert access token" default(Bearer <Add access token here>)
// @Param 	Asset	 		body 	dto.Asset 	true	"Asset Data"
// @Success 204 "OK"
// @Failure 401 {object} dto.Problem "err.unauthorized"
// @Failure 400 {object} dto.Problem "err.processing_param"
// @Failure 502 {object} dto.Problem "err.bad_gateway"
// @Failure 504 {object} dto.Problem "err.network"
// @Router /txs/update_asset [patch]
func (h HBlockchainTxs) UpdateAsset(ctx iris.Context) {
	// getting asset data
	var assetParams dto.Asset
	if err := ctx.ReadJSON(&assetParams); err != nil {
		(*h.response).ResErr(&dto.Problem{Status: iris.StatusBadRequest, Title: schema.ErrProcParam, Detail: err.Error()}, &ctx)
		return
	}

	_, problem := (*h.service).UpdateAssetSvc(assetParams)
	if problem != nil {
		(*h.response).ResErr(problem, &ctx)
		return
	}

	(*h.response).ResOK(&ctx)
}

// CreateAsset create an asset in the world state with provided parameters.
// @Summary create an asset in the world state with provided parameters.
// @Tags Txs.mycc
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Param	Authorization	header	string	true 	"Insert access token" default(Bearer <Add access token here>)
// @Param 	Asset	 		body 	dto.Asset 	true	"Asset Data"
// @Success 204 "OK"
// @Failure 401 {object} dto.Problem "err.unauthorized"
// @Failure 400 {object} dto.Problem "err.processing_param"
// @Failure 502 {object} dto.Problem "err.bad_gateway"
// @Failure 504 {object} dto.Problem "err.network"
// @Router /txs/create_asset [post]
func (h HBlockchainTxs) CreateAsset(ctx iris.Context) {
	// getting asset data
	var assetParams dto.Asset
	if err := ctx.ReadJSON(&assetParams); err != nil {
		(*h.response).ResErr(&dto.Problem{Status: iris.StatusBadRequest, Title: schema.ErrProcParam, Detail: err.Error()}, &ctx)
		return
	}

	_, problem := (*h.service).CreateAssetSvc(assetParams)
	if problem != nil {
		(*h.response).ResErr(problem, &ctx)
		return
	}

	(*h.response).ResOK(&ctx)
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
