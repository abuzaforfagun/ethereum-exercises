package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"os"
)

func Encrypt(plainText string) (string, error) {
	aesKey := os.Getenv("wallet_secret_key")
	if aesKey == "" {
		return "", errors.New("unable to get secret key")
	}

	block, err := aes.NewCipher([]byte(aesKey))
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	encrypted := aesGCM.Seal(nonce, nonce, []byte(plainText), nil)
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func Decrypt(encryptedText string) (string, error) {
	aesKey := os.Getenv("wallet_secret_key")
	if aesKey == "" {
		return "", errors.New("unable to get secret key")
	}

	encryptedData, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	// Create AES block cipher
	block, err := aes.NewCipher([]byte(aesKey))
	if err != nil {
		return "", err
	}

	// Use GCM (Galois/Counter Mode)
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Separate nonce and ciphertext
	nonceSize := aesGCM.NonceSize()
	if len(encryptedData) < nonceSize {
		return "", errors.New("encrypted data too short")
	}
	nonce, ciphertext := encryptedData[:nonceSize], encryptedData[nonceSize:]

	// Decrypt the private key bytes
	plainText, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}
