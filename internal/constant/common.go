package constant

type CtxKey string

const (
	ProductionEnvironment = "production"

	RequestID = CtxKey("reqId")
	UserID    = CtxKey("userId")
	TokenID   = CtxKey("jti")

	DB = "db"

	AccessTokenType  = "accesToken"
	RefreshTokenType = "refreshToken"
)
