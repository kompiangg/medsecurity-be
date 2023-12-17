package rsax

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

func EncryptPrivateKey(privateKey *rsa.PrivateKey, passphrase string, salt string) ([]byte, error) {
	// Convert the private key to PEM format
	privDER := x509.MarshalPKCS1PrivateKey(privateKey)
	privBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privDER,
	}

	// Derive a key from the passphrase using PBKDF2
	key := pbkdf2.Key([]byte(passphrase), []byte(salt), 4096, 32, sha256.New)

	// Encrypt the private key using AES
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	encrypted := gcm.Seal(nonce, nonce, pem.EncodeToMemory(privBlock), nil)

	// Create an encrypted PEM block
	encBlock := &pem.Block{
		Type:  "ENCRYPTED RSA PRIVATE KEY",
		Bytes: encrypted,
	}

	return pem.EncodeToMemory(encBlock), nil
}
