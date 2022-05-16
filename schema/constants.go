package schema

// region ======== i18n ERROR KEYS =======================================================
const (
	ErrAuth               = "err.authentication"
	ErrGeneric            = "err.generic"
	ErrInvalidEnvVar      = "err.invalid.enviroment.var"
	ErrFile               = "err.system_file_related"
	ErrBuntdbItemNotFound = "err.database_related.item_not_found"
	ErrBuntdb             = "err.database_related"
	ErrBuntdbIndex        = "err.database_index_related"
	ErrProcParam          = "err.processing_param"
	ErrJwtGen             = "err.jwt_generation"
	ErrWrongAuthProvider  = "err.wrong_auth_provider"
	ErrEmailSending       = "err.email_sending"
	ErrBlockchainTxs      = "err.blockchain_tx"
	ErrDecodePayloadTx    = "err.decode_payload_tx"
	ErrCryptProc          = "err.crypt_material_processing"
	ErrCryptProcMissing   = "err.crypt_material_processing.missing_files"
)

// endregion =============================================================================

// region ======== ERROR DETAILS =========================================================
const (
	ErrCredsNotFound       = "The provided credentials don't seems to be valid"
	ErrUnauthorized        = "Unauthorized user to execute that action"
	ErrDetContractNotFound = "contract function not found"
	ErrDetInvalidProvider  = "wrong or invalid provider"
	ErrDetInvalidField     = "the given field is invalid"
	ErrDetWalletProc       = "failed to create wallet"
	ErrEmailProc           = "failed to send email"
	ErrDetIdentityCreate   = "failed to create the x509 identity"
)

// endregion =============================================================================

// region ======== SOME STRINGS ==========================================================

const (
	// ENV VARS
	EnvConfigPath = "HLF_DAPP_CONFIG"
	EnvJWTSignKey = "HLF_DAPP_JWT_SIGN_KEY"

	// CRYPTO MATERIALS
	WalletStr = "wallet"

	CryptPkFileName = "priv_sk"
	CryptCertExt    = ".pem"

	// HLF network
	ChDefault = "mychannel" // default channel name

)

// endregion =============================================================================
