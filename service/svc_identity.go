package service

import (
	"github.com/ic-matcom/api.dapp/repo/db"
	"github.com/ic-matcom/api.dapp/repo/hlf"
	"github.com/ic-matcom/api.dapp/schema/dto"
)

// region ======== SETUP =================================================================

type ISvcHlfIdentity interface {
	MkDappIdentity() (string, *dto.Problem)
	GenUsersCerts() *dto.Problem
}

type SvcHlfIdentity struct {
	repoI *hlf.RepoHlfIdentity
	repoU *db.RepoUsers
}

// endregion =============================================================================

// NewSvcHlfIdentity
func NewSvcHlfIdentity(pRepoI *hlf.RepoHlfIdentity, pRepoU *db.RepoUsers) *SvcHlfIdentity { // filling the organization specification struct

	return &SvcHlfIdentity {
		pRepoI,
		pRepoU,
	}
}

// region ======== METHODS ===============================================================

// MkDappIdentity creates the dapp identity to authenticate all the operations in the HLF network.
// Technically, import the cert and private key from and existing HLF identity previously created
// with a CA
func (s *SvcHlfIdentity) MkDappIdentity(forAdmin bool) *dto.Problem {

	if !forAdmin {
		problem := (*s.repoI).MkDappIdentity()
		if problem != nil {
			return problem
		} else {
			return nil
		}
	} else {
		problem := (*s.repoI).MkDappAdminIdentity()
		if problem != nil {
			return problem
		} else {
			return nil
		}
	}
}

// endregion =============================================================================

// region ======== PRIVATE AUX ===========================================================
// endregion =============================================================================
