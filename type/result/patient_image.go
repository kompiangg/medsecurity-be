package result

import (
	"crypto/rand"
	"crypto/rsa"
	"medsecurity/type/model"
	"time"

	"github.com/google/uuid"
)

type PatientImageBriefInformation struct {
	ID          uuid.UUID `json:"id"`
	PatientID   uuid.UUID `json:"patient_id"`
	PatientName string    `json:"patient_name"`
	DoctorID    uuid.UUID `json:"doctor_id"`
	DoctorName  string    `json:"doctor_name"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	IsValid     bool      `json:"is_valid"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (p *PatientImageBriefInformation) FromPatientModel(
	model model.PatientImage,
	patient model.Patient,
	doctor model.Doctor,
) {
	p.ID = model.ID
	p.PatientID = model.PatientID
	p.PatientName = patient.FullName
	p.DoctorID = model.DoctorID
	p.DoctorName = doctor.FullName
	p.Name = model.Name
	p.Type = model.Type
	p.IsValid = model.IsValid
	p.CreatedAt = model.CreatedAt
	p.UpdatedAt = model.UpdatedAt
}

type ServicePatientRequestGetImage struct {
	Token     string `json:"token"`
	ExpiredAt int64  `json:"expired_at"`
}

type ServicePatientGetImage struct {
	DocumentName string `json:"document_name"`
	DocumentType string `json:"document_type"`
	Base64Image  string `json:"base64_image"`
}

func (s *ServicePatientGetImage) DecryptImage(privateKey *rsa.PrivateKey, keySize int, encryptedImage []byte) error {
	var decryptedData []byte
	chunkedDataLength := keySize / 8

	for i := 0; i < len(encryptedImage); i += chunkedDataLength {
		end := i + chunkedDataLength
		if end > len(encryptedImage) {
			end = len(encryptedImage)
		}

		decryptedBytes, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptedImage[i:end])
		if err != nil {
			return err
		}

		decryptedData = append(decryptedData, decryptedBytes...)
	}

	s.Base64Image = string(decryptedData)
	return nil
}

type RepositoryGetRequestPatientImageToken struct {
	ImageID  string `redis:"image_id"`
	Token    string `redis:"token"`
	Password string `redis:"password"`
}
