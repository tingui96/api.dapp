package utils

import (
	"github.com/ic-matcom/api.dapp/schema"
	"github.com/tkanos/gonfig"
	"log"
	"os"
)

// region ======== TYPES =================================================================

// conf unexported configuration schema holder struct
type conf struct {
	// Environment
	Debug    bool
	ApiDocIp string
	DappPort string

	// Cryptographic conf
	JWTSignKey string
	TkMaxAge   uint8

	// HLF Network & Crypto Materials
	MspId      string
	CppPath    string
	DbUserPath string
	DbPath     string

	CryptoMaterialsDir string
	DappIdentityUser   string
	DappIdentityAdmin  string

	AdminUI string

	// Organization & CA Specifications
	OrgPath      string
	OrgName      string
	OrgDomain    string
	CACommonName string
}

// SvcConfig exported configuration service struct
type SvcConfig struct {
	Path string `string:"Path to the config YAML file"`
	conf `conf:"Configuration object"`
}

// endregion =============================================================================

// NewSvcConfig create a new configuration service.
func NewSvcConfig() *SvcConfig {
	c := conf{}

	var configPath = os.Getenv(schema.EnvConfigPath)
	// var jwtSignKey = os.Getenv(schema.EnvJWTSignKey)
	var jwtSignKey = "45567f001601aacb761e13987cddc62ddd49c5b2"

	if configPath == "" {
		log.Println("HLF_DAPP_CONFIG: ", configPath)
		// log.Println("HLF_DAPP_JWT_SIGN_KEY: ", jwtSignKey)
		panic(schema.ErrInvalidEnvVar)
	}

	err := gonfig.GetConf(configPath, &c) // getting the conf
	if err != nil {
		panic(err)
	} // error check

	c.JWTSignKey = jwtSignKey // saving the sign key into the configuration object

	return &SvcConfig{configPath, c} // We are using struct composition here. Hence the anonymous field (https://golangbot.com/inheritance/)
}
