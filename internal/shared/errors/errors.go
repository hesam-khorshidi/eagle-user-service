package errors

import "github.com/pkg/errors"

type Error error

var (
	ErrEntityNotFound      = errors.New("entity not found")
	ErrEntityAlreadyExists = errors.New("entity already exists")
)

var (
	ErrInvalidToken           = errors.New("invalid token")
	ErrTokenExpired           = errors.New("refresh token expired")
	ErrTokenNotValidYet       = errors.New("refresh token not valid yet")
	ErrInvalidTokenSignature  = errors.New("invalid refresh token signature")
	ErrTokenBlacklisted       = errors.New("refresh token is blacklisted")
	ErrMissingTokenJTI        = errors.New("refresh token missing JTI claim")
	ErrMissingTokenUserID     = errors.New("refresh token missing user ID claim")
	ErrUserNotFoundForToken   = errors.New("user associated with token not found or inactive")
	ErrInvalidTokenType       = errors.New("invalid token")
	ErrUserIDExtractionFailed = errors.New("failed to extract user ID from token")
)

var (
	ErrCacheNotFound = errors.New("cache not found")
)
