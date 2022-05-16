package hlf

import (
	"github.com/ic-matcom/api.dapp/schema/ccFuncNames"
	"github.com/ic-matcom/api.dapp/schema/dto"
	"errors"
	jsoniter "github.com/json-iterator/go"
	"path/filepath"

	"github.com/ic-matcom/api.dapp/schema"
	"github.com/ic-matcom/api.dapp/service/utils"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

// region ======== SETUP =================================================================

type RepoBlockchain interface {
	InitLedger() ([]byte, error)
	Get(request dto.GetRequest) ([]byte, error)
}

type repoBlockchain struct {
	ChannelName       string          // ChannelName HLF channel name
	CppPath           string          // CppPath path to the connection profile
	WalletPath        string          // WalletPath path to the wallets folders
	Wallet            *gateway.Wallet // Wallet with admin privilege identity for admins ops on the network
	DappIdentityUser  string          // DappIdentityUser dapp user identity to authenticate normal dapp ops in the HLF network
	DappIdentityAdmin string          // DappIdentityAdmin dapp admin identity to authenticate admin dapp ops in the HLF network
}

// endregion =============================================================================

func NewRepoBlockchain(SvcConf *utils.SvcConfig) RepoBlockchain {

	wallet, err := gateway.NewFileSystemWallet(filepath.Join(SvcConf.CryptoMaterialsDir, schema.WalletStr))
	if err != nil {
		panic(schema.ErrDetWalletProc + " ." + err.Error())
	}

	return &repoBlockchain{
		schema.ChDefault,
		SvcConf.CppPath,
		SvcConf.CryptoMaterialsDir,
		wallet,
		SvcConf.DappIdentityUser,
		SvcConf.DappIdentityAdmin,
	}
}

// region ======== METHODS EVOTE =========================================================

// Suffrage_InitLedger
func (r *repoBlockchain) InitLedger() ([]byte, error) {

	// getting components instance
	gw, _, contract, e := r.getSDKComponents(r.ChannelName, ccfuncnames.ContractNameCC1 , true)
	if e != nil {
		return nil, e
	}
	defer gw.Close()

	// Creating the initial data in the ledger
	issuer, e := contract.SubmitTransaction(ccfuncnames.CC1InitLedger, "[]")
	if e != nil {
		return nil, e
	}

	return issuer, nil
}

func (r *repoBlockchain) Get(request dto.GetRequest) ([]byte, error) {
	// getting components instance
	gw, _, contract, e := r.getSDKComponents(r.ChannelName, ccfuncnames.ContractNameCC1 , false)
	if e != nil {
		return nil, e
	}
	defer gw.Close()

	strArgs, _ := jsoniter.Marshal(request)
	res, e := contract.EvaluateTransaction(ccfuncnames.CC2ReadAsset, string(strArgs))
	//res, e := contract.SubmitTransaction(ccfuncnames.CC2ReadAsset, string(strArgs))
	if e != nil {
		return nil, e
	}

	return res, nil
}

// endregion =============================================================================

// region ======== PRIVATE AUX ===========================================================

// getSDKComponents create the instances for the main components of HLF SDK: gateway, network and contract
//
// - channel [string] ~ HLF / Channel name
//
// - contractName [string] ~ chaincode contract name to invoke
//
// - withAdminIdentity [bool] ~ do we need to use the administration dapp identity fo the transaction ?
func (r *repoBlockchain) getSDKComponents(channel, contractName string, withAdminIdentity bool) (*gateway.Gateway, *gateway.Network, *gateway.Contract, error) {
	var identityLabel = r.DappIdentityUser

	if withAdminIdentity {
		identityLabel = r.DappIdentityAdmin
	}

	// trying to get an instance of HLF SDK network gateway, from the connection profile
	gw, e := gateway.Connect( // gt = gateway
		gateway.WithConfig(config.FromFile(filepath.Clean(r.CppPath))),
		gateway.WithIdentity(r.Wallet, identityLabel))
	if e != nil {
		return nil, nil, nil, e
	}

	// trying to get an instance of the gateway network
	nt, e := gw.GetNetwork(channel) // nt == network
	if e != nil {
		return nil, nil, nil, e
	}

	// trying to get the contract
	contract := nt.GetContract(contractName)
	//contract := nt.GetContractWithName(schema.CceVote, contractName) // contractName = chaincode
	if contract == nil {
		return nil, nil, nil, errors.New(schema.ErrDetContractNotFound)
	}

	// so far so good, returning the instance pointers
	return gw, nt, contract, nil
}

// endregion =============================================================================