package valueobject

type TokenType string

var (
	TokenTypeAccessToken  = TokenType("access_token")
	TokenTypeRefreshToken = TokenType("refresh_token")
)
