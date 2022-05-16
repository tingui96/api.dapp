package service

import (
	"github.com/ic-matcom/api.dapp/repo/db"
	"github.com/ic-matcom/api.dapp/repo/hlf"
	"github.com/ic-matcom/api.dapp/schema/dto"
	"github.com/ic-matcom/api.dapp/service/utils"
	"path/filepath"
)

// region ======== SETUP =================================================================

type ISvcHlfIdentity interface {
	MkDappIdentity() (string, *dto.Problem)
	GenUsersCerts() *dto.Problem
}

type SvcHlfIdentity struct {
	repoI *hlf.RepoHlfIdentity
	repoU *db.RepoUsers

	orgPath string
	orgSpec dto.OrgSpec

	userPath  string
	caDir     string
	tlscaPath string
}

// endregion =============================================================================

// NewSvcHlfIdentity
func NewSvcHlfIdentity(pRepoI *hlf.RepoHlfIdentity, pRepoU *db.RepoUsers, conf *utils.SvcConfig) *SvcHlfIdentity { // filling the organization specification struct

	// filling important paths
	usersDir := filepath.Join(conf.OrgPath, conf.OrgDomain, "users")
	caDir := filepath.Join(conf.OrgPath, conf.OrgDomain, "ca")
	tlscaPath := filepath.Join(conf.OrgPath, conf.OrgDomain, "tlsca")

	// filling organization specification
	org := dto.OrgSpec {
		Name:          conf.OrgName,
		Domain:        conf.OrgDomain,
		EnableNodeOUs: true,
		CA:            dto.NodeSpec{CommonName: conf.CACommonName},
		Template:      dto.NodeTemplate{},
		Specs:         nil,
		Users:         dto.UsersSpec{Count: 0}, // we need to update it later
	}

	return &SvcHlfIdentity {
		pRepoI,
		pRepoU,
		conf.OrgPath,
		org,
		usersDir,
		caDir,
		tlscaPath,
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
