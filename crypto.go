package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

type Crypto struct {
	Password string
}

func New(p string) *Crypto {
	return &Crypto{p}
}

func (c *Crypto) generateKey(password string) ([]byte, error) {
	if err := ValidatePassword(password); err != nil {
		return nil, err
	}
	hash := sha256.Sum256([]byte(password))
	return hash[:], nil
}

func (c *Crypto) EncryptText(plaintext, password string) (string, error) {
	key, genErr := c.generateKey(password)
	if genErr != nil {
		return "", genErr
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func (c *Crypto) DecryptText(ciphertext, password string) (string, error) {
	key, genErr := c.generateKey(password)
	if genErr != nil {
		return "", genErr
	}

	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := data[:nonceSize], string(data[nonceSize:])

	plaintext, err := gcm.Open(nil, nonce, []byte(ciphertext), nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func (c *Crypto) EncryptFile(filename, password string) error {
	vErr := ValidateFilename(filename, ENC)
	if vErr != nil {
		return vErr
	}

	plaintext, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	ciphertext, err := c.EncryptText(string(plaintext), password)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename+".enc", []byte(ciphertext), 0644)
}

func (c *Crypto) DecryptFile(filename, password string) error {
	vErr := ValidateFilename(filename, DEC)
	if vErr != nil {
		return vErr
	}

	ciphertext, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	plaintext, err := c.DecryptText(string(ciphertext), password)
	if err != nil {
		return err
	}

	newfilename := strings.Split(filename, ".enc")[0] + ".dec"

	return ioutil.WriteFile(newfilename, []byte(plaintext), 0644)
}
