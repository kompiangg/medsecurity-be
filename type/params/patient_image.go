package params

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"medsecurity/pkg/errors"
	"medsecurity/type/model"

	"github.com/google/uuid"
	"github.com/volatiletech/null/v9"
)

type RepositoryFindPatientImage struct {
	PatientID null.String
	DoctorID  null.String

	RepositoryPaginationParam
}

type ServiceFindPatientImage struct {
	PatientID null.String `query:"patient_id"`
	DoctorID  null.String `query:"doctor_id"`

	Role      string `param:"-"`
	AccountID string `param:"-"`

	ServicePaginationParam
}

type ServiceCreatePatientImage struct {
	PatientID      string `json:"patient_id" validate:"required,uuid4"`
	DocumentName   string `json:"document_name" validate:"required"`
	DocumentType   string `json:"document_type" validate:"required"`
	Base64Document string `json:"base64_document" validate:"required,base64"`

	DoctorID string `json:"-"`
}

func (s ServiceCreatePatientImage) EncryptBase64Document(publicKey string, keySize int) ([]byte, error) {
	block, _ := pem.Decode([]byte(publicKey))

	publicKeyParsed, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return []byte{}, err
	}

	rsaPublicKey, ok := publicKeyParsed.(*rsa.PublicKey)
	if !ok {
		return []byte{}, errors.New("invalid public key")
	}

	var encryptedByteData []byte
	chunkedSize := keySize/8 - 11

	for i := 0; i < len(s.Base64Document); i += chunkedSize {
		end := i + chunkedSize
		if end > len(s.Base64Document) {
			end = len(s.Base64Document)
		}

		encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPublicKey, []byte(s.Base64Document[i:end]))
		if err != nil {
			return []byte{}, err
		}

		encryptedByteData = append(encryptedByteData, encryptedData...)
	}

	return encryptedByteData, nil
}

func (s ServiceCreatePatientImage) ToPatientImageModel(url string) (model.PatientImage, error) {
	patientID, err := uuid.Parse(s.PatientID)
	if err != nil {
		return model.PatientImage{}, err
	}

	doctorID, err := uuid.Parse(s.DoctorID)
	if err != nil {
		return model.PatientImage{}, err
	}

	return model.PatientImage{
		ID:        uuid.New(),
		PatientID: patientID,
		DoctorID:  doctorID,
		Name:      s.DocumentName,
		Type:      s.DocumentType,
		IsValid:   true,
		URL:       url,
	}, nil
}
