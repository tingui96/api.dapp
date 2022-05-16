package auth

import (
	"github.com/ic-matcom/api.dapp/repo/db"
	"github.com/ic-matcom/api.dapp/service/utils"
)

type SvcAuthentication struct {
	AuthProviders map[string]Provider 			// similar to slices, maps are reference types.
}

// NewSvcAuthentication creates the authentication service. It provides the methods to make the
// authentication intent with the register providers.
//
// - providers [Array] ~ Maps of providers string token / identifiers
//
// - conf [*SvcConfig] ~ App conf instance pointer
func NewSvcAuthentication(providers map[string]bool, conf *utils.SvcConfig, repoUser *db.RepoUsers) *SvcAuthentication {

	k := &SvcAuthentication{AuthProviders: make(map[string]Provider)}

	for v, _ := range providers {

		if v == "default" {
			(*k).AuthProviders[v] = &ProviderEvote {
				walletLocations: conf.CryptoMaterialsDir,
				repo: repoUser,
			}
		} else if v == "other" {					// ===== DEFAULT CASE, NORMAL DATABASE LOGIN =======
			// TODO implement normal database login authentication case
		}
	}

	return k
}
