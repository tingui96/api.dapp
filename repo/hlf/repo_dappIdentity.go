package hlf

import (
	"errors"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	"github.com/kataras/iris/v12"
	"io/ioutil"
	"path/filepath"

	"github.com/ic-matcom/api.dapp/lib"
	"github.com/ic-matcom/api.dapp/schema"
	"github.com/ic-matcom/api.dapp/schema/dto"
	"github.com/ic-matcom/api.dapp/service/utils"
)

// region ======== SETUP =================================================================

type RepoHlfIdentity interface {
	MkDappIdentity() *dto.Problem
	MkDappAdminIdentity() *dto.Problem

	populateWallet(wallet *gateway.Wallet, isAdmin bool) error
	reProblem(title string, msg string) *dto.Problem
}

type repoHlfIdentity struct {
	MspId              string
	CryptoMaterialsDir string // WalletPath path to the wallets folders
	DappIdentityUser   string
	DappIdentityAdmin  string
}

// endregion =============================================================================

func NewRepoIdentity(SvcConf *utils.SvcConfig) RepoHlfIdentity {

	return &repoHlfIdentity {
		SvcConf.MspId,
		SvcConf.CryptoMaterialsDir,
		SvcConf.DappIdentityUser,
		SvcConf.DappIdentityAdmin,
	}
}

// region ======== METHODS ===============================================================

// MkDappIdentity crete the x509 USER dapp identity and put it in the correspondent wallet path.
// The data used for the creation is the private key and cert from an existing HLF identity,
// hence this is really and importation.
// Remember that those identities is only for authenticate the dapp to the HLF network
//
// If this method returns nil all was good
func (r *repoHlfIdentity) MkDappIdentity() *dto.Problem {

	// creating the user wallet gateway instance
	wallet, err := gateway.NewFileSystemWallet(filepath.Join(r.CryptoMaterialsDir, schema.WalletStr))
	if err != nil {
		return r.reProblem(schema.ErrCryptProc, schema.ErrDetWalletProc + " ." + err.Error())
	}

	// check the existence of normal user dapp identity
	if wallet.Exists(r.DappIdentityUser) { return nil }

	// so we need to create the x509 identity and putting in the wallet, actually is an importation from HLF identity
	e := r.populateWallet(wallet, false)
	if e != nil {
		return r.reProblem(schema.ErrCryptProc, schema.ErrDetIdentityCreate + " ." + err.Error())
	}

	// so far so good
	return nil
}

// MkDappAdminIdentity create the x509 ADMIN dapp identity and put it in the correspondent wallet path
// The data used for the creation is the private key and cert from an existing HLF identity,
// hence this is really and importation.
// Remember that those identities is only for authenticate the dapp to the HLF network
//
// If this method returns nil all was good
func (r *repoHlfIdentity) MkDappAdminIdentity() *dto.Problem {
	// creating the admin wallet gateway instance
	wallet, err := gateway.NewFileSystemWallet(filepath.Join(r.CryptoMaterialsDir, schema.WalletStr))
	if err != nil {
		return r.reProblem(schema.ErrCryptProc, schema.ErrDetWalletProc + " ." + err.Error())
	}

	// check the existence of normal admin dapp identity
	if wallet.Exists(r.DappIdentityAdmin) { return nil }

	// so we need to create the x509 identity and putting in the wallet, actually is an importation from HLF identity
	e := r.populateWallet(wallet, true)
	if e != nil {
		return r.reProblem(schema.ErrCryptProc, schema.ErrDetIdentityCreate + " ." + err.Error())
	}

	// so far so good
	return nil
}

// endregion =============================================================================

// region ======== PRIVATE AUX ===========================================================

// populateWallet create the x509 dapp user or dapp admin identity and put those in the wallet
// folder. Remember that those identities is only for authenticate the dapp to the HLF network
//
// - wallet [*gateway.Wallet] ~ wallet instance
//
// - isAdmin [isAdmin] ~ flag to know if we have to work with the admin identity
func (r *repoHlfIdentity) populateWallet (wallet *gateway.Wallet, isAdmin bool) error {

	// Getting .pem and private key files, Part 1 setup
	// non admin case
	var iPath = filepath.Join(r.CryptoMaterialsDir, "msp", r.DappIdentityUser)
	identityLabel := r.DappIdentityUser
	if isAdmin {
		iPath = filepath.Join(r.CryptoMaterialsDir, "msp", r.DappIdentityAdmin)
		identityLabel = r.DappIdentityAdmin
	}

	// Getting .pem and private key files, Part 2 walking the crypto materials dir according to the user
	var pem = lib.GetFilesByExt(filepath.Join(iPath), schema.CryptCertExt)        // getting the .pem files
	var pvs = lib.GetFilesByName(filepath.Join(iPath), schema.CryptPkFileName)	  // getting the private key files

	if len(pem) == 0 || len(pvs) == 0 {return errors.New(schema.ErrCryptProcMissing)}

	// reading cert
	cert, e := ioutil.ReadFile(filepath.Clean(pem[0]))
	if e != nil { return e }

	// read the private key
	key, err := ioutil.ReadFile(filepath.Clean(pvs[0]))
	if err != nil { return err }

	// create the identity and saving in the wallet
	identity := gateway.NewX509Identity(r.MspId, string(cert), string(key))
	return wallet.Put(identityLabel, identity)
}

// reProblem Internal wrapper method to report problem to the caller
func (r *repoHlfIdentity) reProblem(title string, msg string) *dto.Problem {
	return dto.NewProblem(iris.StatusInternalServerError, title, msg)
}

// endregion =============================================================================
