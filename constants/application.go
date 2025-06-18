package constants

type AppKey string
type ctxKey string

const (
	ApplicationCtx     ctxKey = "application"
	AuthTokenCtx       ctxKey = "auth_token"
	AuthTokenParsedCtx ctxKey = "auth_token_parsed"
)

type localsKey string

const (
	ConfigLocals    localsKey = "config"
	JwtUserLocals   localsKey = "jwt_user"
	DBLocals        localsKey = "db"
	RequestDBLocals localsKey = "request_db"
)
