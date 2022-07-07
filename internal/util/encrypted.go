package util

import (
	"crypto/aes"
	gocipher "crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"strings"

	"github.com/feditools/democrablock/internal/config"
	"github.com/spf13/viper"
)

var (
	ErrDataTooSmall = errors.New("data too small")
)

func Decrypt(b []byte) ([]byte, error) {
	gcm, err := getCrypto()
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(b) < nonceSize {
		return nil, ErrDataTooSmall
	}

	nonce, ciphertext := b[:nonceSize], b[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func Encrypt(b []byte) ([]byte, error) {
	gcm, err := getCrypto()
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, b, nil), nil
}

func getCrypto() (gocipher.AEAD, error) {
	key := []byte(strings.ToLower(viper.GetString(config.Keys.EncryptionKey)))
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := gocipher.NewGCM(cipher)
	if err != nil {
		return nil, err
	}

	return gcm, nil
}
