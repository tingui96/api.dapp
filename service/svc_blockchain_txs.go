package service

import (
	"github.com/ic-matcom/api.dapp/lib"
	"github.com/ic-matcom/api.dapp/repo/db"
	"github.com/ic-matcom/api.dapp/repo/hlf"
	"github.com/ic-matcom/api.dapp/schema"
	"github.com/ic-matcom/api.dapp/schema/dto"
	"github.com/kataras/iris/v12"
)

// region ======== SETUP =================================================================

// ISvcBlockchainTxs Blockchain transactions service interface
type ISvcBlockchainTxs interface {
	SrvInitLedger() ([]byte, *dto.Problem)
	GetUserSvc(id string) (*dto.User, *dto.Problem)
	GetBallotSvc(id string) (interface{}, *dto.Problem)
}

type svcBlockchainTxs struct {
	repo *hlf.RepoBlockchain
	reposUser *db.RepoUsers
}

// endregion =============================================================================

// NewSvcBlockchainTxs instantiate the HLF blockchains transactions services
func NewSvcBlockchainTxs(pRepo *hlf.RepoBlockchain, reposUser *db.RepoUsers) ISvcBlockchainTxs {
	return &svcBlockchainTxs{pRepo, reposUser }
}

// region ======== METHODS IDENTITY ======================================================


// Identity_GetIssuer Get an issuer from the blockchain ledger according with the given id.
//func (s *svcBlockchainTxs) Identity_GetIssuer(id string) ([]byte, *dto.Problem) {
//
//	issuer, e := (*s.repo).Identity_GetIssuer(id)
//	if e != nil {
//		return nil, dto.NewProblem(iris.StatusBadGateway, schema.ErrBlockchainTxs, e.Error())
//	}
//
//	return issuer, nil
//}

// endregion =============================================================================

// region ======== METHODS SUFFRAGE ======================================================

// SrvInitLedger Is a initial method for creating the necessary data for the suffrage contract
// and saving it to the ledger.
// contract: suffrage, initLedger
func (s *svcBlockchainTxs) SrvInitLedger() ([]byte, *dto.Problem) {

	data, e := (*s.repo).InitLedger()
	if e != nil {
		return nil, dto.NewProblem(iris.StatusBadGateway, schema.ErrBlockchainTxs, e.Error())
	}

	return data, nil
}

// endregion =============================================================================

// region ======== PRIVATE AUX ===========================================================
// endregion =============================================================================


func (s *svcBlockchainTxs) GetUserSvc(id string)  (*dto.User, *dto.Problem) {
	res, err := (*s.reposUser).GetUser(id)
	if err != nil {
		return nil, dto.NewProblem(iris.StatusExpectationFailed, schema.ErrBuntdb, err.Error())
	}
	return res, nil
}

func (s *svcBlockchainTxs) GetBallotSvc(id string) (interface{}, *dto.Problem) {
	r := dto.GetRequest{
		ID:         "ID1",
		ElectionID: "EID1",
	}
	item, err := (*s.repo).Get(r)
	if err != nil {
		return nil, dto.NewProblem(iris.StatusExpectationFailed, schema.ErrBuntdb, err.Error())
	}
	result := lib.DecodePayload(item)

	m, ok := result.([]interface{})
	if !ok {return nil, dto.NewProblem(iris.StatusExpectationFailed, schema.ErrDecodePayloadTx, err.Error())}
	if len(m) > 0 {return m[0], nil}

	return dto.ElectionRequest{}, nil
}