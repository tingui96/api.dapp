package auth

import (
	"github.com/ic-matcom/api.dapp/repo/db"
	"github.com/ic-matcom/api.dapp/service/utils"
)

type SvcAuthentication struct {
	AuthProviders map[string]Provider // similar to slices, maps are reference types.
}

// NewSvcAuthentication creates the authentication service. It provides the methods to make the
// authentication intent with the register providers.
//
// - providers [Array] ~ Maps of providers string token / identifiers
//
// - conf [*SvcConfig] ~ App conf instance pointer
func NewSvcAuthentication(conf *utils.SvcConfig, repoUser *db.RepoUsers) *SvcAuthentication {

	k := &SvcAuthentication{AuthProviders: make(map[string]Provider)}

	(*k).AuthProviders["default"] = &ProviderEvote{
		walletLocations: conf.CryptoMaterialsDir,
		repo:            repoUser,
	}

	return k
}
