package db

import (
	"github.com/ic-matcom/api.dapp/schema/dto"
	"github.com/ic-matcom/api.dapp/service/utils"
)

// region ======== SETUP =================================================================

type RepoUsers interface {
	GetUser(field string) (*dto.User, error)
}

type repoUsers struct {
	DBUserLocation string
}
// endregion =============================================================================

// NewRepoUsers
func NewRepoUsers(SvcConf *utils.SvcConfig) RepoUsers {
	return &repoUsers{DBUserLocation: SvcConf.DbPath}
}

// region ======== METHODS ===============================================================

// GetUser get the user from the DB file that should be compliant with the dto.UserList struct
// return a list of dto.User
func (r *repoUsers) GetUser(field string) (*dto.User, error) {
	user := dto.User{
		Id:         "",
		Passphrase: "",
		Clear:      "",
		Username:   "",
		Name:       "",
	}

	return &user, nil
}

// endregion =============================================================================