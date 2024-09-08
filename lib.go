package crypto

import "strings"

func ValidateFilename(filename, ops string) error {
	if filename == "" {
		return ErrFileNotFound
	}

	if strings.Contains(filename, ".enc") {
		if ops == DEC {
			return nil
		}

		return ErrFileEncrypted
	}

	if !strings.Contains(filename, ".enc") && ops == DEC {
		return ErrFileDecrypted
	}

	return nil
}

func ValidatePassword(password string) error {
	if password == "" {
		return ErrEmptyPasswd
	}
	if len(password) < 4 {
		return ErrPasswdTooShort
	}
	return nil
}
