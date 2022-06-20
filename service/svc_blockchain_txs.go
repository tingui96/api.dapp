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
	ReadAssetSvc(id string) (interface{}, *dto.Problem)
	CreateAssetSvc(asset dto.Asset) (interface{}, *dto.Problem)
	UpdateAssetSvc(asset dto.Asset) (interface{}, *dto.Problem)
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

// region ======== METHODS ======================================================


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

func (s *svcBlockchainTxs) GetUserSvc(id string)  (*dto.User, *dto.Problem) {
	res, err := (*s.reposUser).GetUser(id)
	if err != nil {
		return nil, dto.NewProblem(iris.StatusExpectationFailed, schema.ErrBuntdb, err.Error())
	}
	return res, nil
}

func (s *svcBlockchainTxs) ReadAssetSvc(ID string) (interface{}, *dto.Problem) {
	item, err := (*s.repo).ReadAsset(ID)
	if err != nil {
		return nil, dto.NewProblem(iris.StatusExpectationFailed, schema.ErrBuntdb, err.Error())
	}

	result := lib.DecodePayload(item)

	m, ok := result.(interface{})
	if !ok {return nil, dto.NewProblem(iris.StatusExpectationFailed, schema.ErrDecodePayloadTx, err.Error())}

	return m, nil
}

func (s *svcBlockchainTxs) CreateAssetSvc(asset dto.Asset) (interface{}, *dto.Problem) {
	item, err := (*s.repo).CreateAsset(asset)
	if err != nil {
		return nil, dto.NewProblem(iris.StatusExpectationFailed, schema.ErrBuntdb, err.Error())
	}

	result := lib.DecodePayload(item)

	m, ok := result.(interface{})
	if !ok {return nil, dto.NewProblem(iris.StatusExpectationFailed, schema.ErrDecodePayloadTx, err.Error())}

	return m, nil
}

func (s *svcBlockchainTxs) UpdateAssetSvc(asset dto.Asset) (interface{}, *dto.Problem) {
	item, err := (*s.repo).UpdateAsset(asset)
	if err != nil {
		return nil, dto.NewProblem(iris.StatusExpectationFailed, schema.ErrBuntdb, err.Error())
	}

	result := lib.DecodePayload(item)

	m, ok := result.(interface{})
	if !ok {return nil, dto.NewProblem(iris.StatusExpectationFailed, schema.ErrDecodePayloadTx, err.Error())}

	return m, nil
}


// FALTA el CREATEASSSET