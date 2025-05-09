package errors

type ErrorType string

const (
	ErrTypeReportable = ErrorType("report")
	ErrTypeRaw        = ErrorType("raw")
)
