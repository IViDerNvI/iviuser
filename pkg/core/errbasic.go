package core

// Basic error codes and messages for the application.
var (
	// Error codes for common HTTP errors.
	ErrOK = NewErrCode(200, 0, "OK")
	// Error codes for Unknwon errors.
	ErrUnknownError = NewErrCode(500, 999999, "Unknown error")
)
