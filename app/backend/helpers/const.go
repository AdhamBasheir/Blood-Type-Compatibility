package helpers

const (
	TimeFormat = "2006-01-02 15-04-05 Mon"
)

// Session and security related constants
const (
	SessionDuration    = 3600 * 24 * 7 // 7 days
	SessionTokenLength = 32
	CSRFTokenLength    = 32
)

// User related constants
const (
	MaxLoginAttempts   = 5
	LoginAttemptWindow = 15 // in minutes
	PasswordMinLength  = 8
	PasswordMaxLength  = 64
)
