package model

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"reflect"
	"time"

	"github.com/google/uuid"
)

type PatientSecret struct {
	ID         uuid.UUID `db:"id"`
	PatientID  uuid.UUID `db:"patient_id"`
	PrivateKey string    `db:"private_key"`
	PublicKey  string    `db:"public_key"`
	KeySize    int       `db:"key_size"`
	Salt       string    `db:"salt"`
	IsValid    bool      `db:"is_valid"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

func (p PatientSecret) IsPrivateKeyPublicKeyMatch(privateKey *rsa.PrivateKey) (bool, error) {
	block, _ := pem.Decode([]byte(p.PublicKey))

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false, err
	}

	publicKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return false, err
	}

	if reflect.DeepEqual(privateKey.PublicKey, *publicKey) {
		return true, nil
	}

	return false, nil
}
