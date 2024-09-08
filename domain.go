package crypto

import "errors"

const (
	ENC = "enc"
	DEC = "dec"
)

var (
	ErrEmptyPasswd    = errors.New("password cannot be empty")
	ErrFileNotFound   = errors.New("file not found")
	ErrFileEncrypted  = errors.New("file already encrypted")
	ErrPasswdTooShort = errors.New("password too short")
	ErrFileDecrypted  = errors.New("file already decrypted")
)
