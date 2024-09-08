package crypto

import "errors"

const (
	ENC = "enc"
	DEC = "dec"
)

var ErrEmptyPasswd = errors.New("password cannot be empty")
var ErrFileNotFound = errors.New("file not found")
var ErrFileEncrypted = errors.New("file already encrypted")
var ErrPasswdTooShort = errors.New("password too short")
var ErrFileDecrypted = errors.New("file already decrypted")
