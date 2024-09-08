package crypto_test

import (
	"testing"

	"github.com/jantttez/crypto"
)

type CryptoCases struct {
	testName, pass, filename string
	expectedErr              error
}

var EncryptoTestCases = []CryptoCases{
	{"full valid", "password", "hello.txt", nil},
	{"empty password", "", "hello.txt", crypto.ErrEmptyPasswd},
	{"short password", "123", "hello.txt", crypto.ErrPasswdTooShort},
	{"empty filename", "password", "", crypto.ErrFileNotFound},
	{"file already encrypted", "password", "hello.txt.enc", crypto.ErrFileEncrypted},
}

var DecryptoTestCases = []CryptoCases{
	{"full valid", "password", "hello.txt.enc", nil},
	{"empty password", "", "hello.txt.enc", crypto.ErrEmptyPasswd},
	{"short password", "123", "hello.txt.enc", crypto.ErrPasswdTooShort},
	{"empty filename", "password", "", crypto.ErrFileNotFound},
	{"file already decrypted", "password", "hello.txt.dec", crypto.ErrFileDecrypted},
}

func TestEncryptFile(t *testing.T) {
	for _, tc := range EncryptoTestCases {
		t.Run(tc.testName, func(t *testing.T) {
			c := crypto.New(tc.pass)
			if err := c.EncryptFile(tc.filename, tc.pass); err != tc.expectedErr {
				t.Errorf("EncryptFile() error = %v, expectedErr %v", err, tc.expectedErr)
			}
		})
	}
}

func TestDecryptFile(t *testing.T) {
	for _, tc := range DecryptoTestCases {
		t.Run(tc.testName, func(t *testing.T) {
			c := crypto.New(tc.pass)
			if err := c.DecryptFile(tc.filename, tc.pass); err != tc.expectedErr {
				t.Errorf("DecryptFile() error = %v, expectedErr %v", err, tc.expectedErr)
			}
		})
	}
}
