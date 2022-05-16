package lib

import (
	"github.com/kataras/iris/v12/middleware/jwt"
	"time"

	"github.com/ic-matcom/api.dapp/schema/dto"
)

// MkAccessToken create a signed JTW token with the specified data. This could be used for authentication purpose by a middleware
func MkAccessToken(data *dto.AccessTokenData, sigKey []byte, TkAge uint8) ([]byte, error) {

	// https://github.com/kataras/iris/blob/master/_examples/auth/jwt/middleware/main.go | https://github.com/iris-contrib/examples/blob/master/auth/jwt/basic/main.go
	tk, err := jwt.Sign(jwt.HS256, sigKey, data, jwt.MaxAge(time.Duration(TkAge)*time.Minute))
	if err != nil { return nil, err }

	return tk, err
}
