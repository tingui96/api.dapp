package auth

import (
	"github.com/ic-matcom/api.dapp/lib"
	"github.com/ic-matcom/api.dapp/repo/db"
	"github.com/ic-matcom/api.dapp/schema"
	"github.com/ic-matcom/api.dapp/schema/dto"
	"github.com/kataras/iris/v12"
)

type Provider interface {
	GrantIntent(userCredential *dto.UserCredIn, data interface{}) (*dto.GrantIntentResponse, *dto.Problem)
}

// region ======== EVOTE AUTHENTICATION PROVIDER =========================================

type ProviderEvote struct {
	walletLocations string // location fo the wallet, ❗ temporal if 'cause I don't know what to do here right now
	repo            *db.RepoUsers
}

func (p *ProviderEvote) GrantIntent(uCred *dto.UserCredIn, options interface{}) (*dto.GrantIntentResponse, *dto.Problem) {
	// getting the user
	user1 := dto.User{
		Id:         "id_roronoa_zoro",
		Passphrase: "f6e248ea994f3e342f61141b8b8e3ede86d4de53257abc8d06ae07a1da73fb39",
		Clear:      "my_password",
		Username:   "zoro@matcom.uh.cu",
		Name:       "Roronoa Zoro",
	}
	checksum, _ := lib.Checksum("SHA256", []byte(uCred.Password))
	if user1.Passphrase == checksum {
		return &dto.GrantIntentResponse{Identifier: user1.Username, DID: user1.Id}, nil
	}

	return nil, dto.NewProblem(iris.StatusNotFound, schema.ErrFile, schema.ErrCredsNotFound)
}

// endregion =============================================================================
