package mapper

import (
	"github.com/ic-matcom/api.dapp/schema/dto"
	"strings"
)

// TIP ref https://hellokoding.com/crud-restful-apis-with-go-modules-wire-gin-gorm-and-mysql/
// Is we need it, this method can perform validation and return two values: the mapped struct and the error

// ToAccessTokenDataV region ======== AUTHORIZATION =========================================================
// dto.GrantIntentResponse to dto.AccessTokenData
// TODO ground the rol idea, according to the Evote app logic
func ToAccessTokenDataV(obj *dto.GrantIntentResponse) *dto.AccessTokenData {
	// claims := dto.Claims{ Sub: obj.Identifier, Rol: "undefined" }
	claims := dto.InjectedParam{ Did: obj.DID, Username: obj.Identifier }

	return &dto.AccessTokenData{ Scope: strings.Fields("api.dapp"), Claims: claims }
}

// endregion =============================================================================
