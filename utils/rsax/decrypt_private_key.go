package rsax

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"

	"golang.org/x/crypto/pbkdf2"
)

func DecryptPrivateKey(encryptedKey []byte, passphrase string) (*rsa.PrivateKey, error) {
	// Decode the PEM data
	block, _ := pem.Decode(encryptedKey)
	if block == nil || block.Type != "ENCRYPTED RSA PRIVATE KEY" {
		return nil, errors.New("no valid PEM data found")
	}

	// Derive the key from the passphrase
	key := pbkdf2.Key([]byte(passphrase), []byte(passphrase), 10000, 32, sha256.New)

	// Decrypt the data
	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(block.Bytes) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := block.Bytes[:nonceSize], block.Bytes[nonceSize:]
	decryptedData, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	// Decode the PEM data
	decryptedBlock, _ := pem.Decode(decryptedData)
	if decryptedBlock == nil || decryptedBlock.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("no valid PEM data found after decryption")
	}

	// Parse the private key
	privateKey, err := x509.ParsePKCS1PrivateKey(decryptedBlock.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}
