package constant

type CtxKey string

const (
	ProductionEnvironment = "production"

	RequestID = CtxKey("reqId")
	UserID    = CtxKey("userId")

	DB = "db"
)
