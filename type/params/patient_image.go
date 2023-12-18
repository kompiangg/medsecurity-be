package params

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"medsecurity/pkg/errors"
	"medsecurity/type/model"
	"time"

	"github.com/google/uuid"
	"github.com/volatiletech/null/v9"
	"golang.org/x/crypto/bcrypt"
)

type RepositoryFindAllPatientImage struct {
	ImageID   null.String
	PatientID null.String
	DoctorID  null.String

	RepositoryPagination
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

func (s ServiceCreatePatientImage) ToAccessHistoryModel(imageID uuid.UUID) model.AccessHistory {
	return model.AccessHistory{
		ID:             uuid.New(),
		PatientImageID: imageID,
		PatientID:      null.NewString("", false),
		DoctorID:       null.NewString(s.DoctorID, true),
		Purpose:        "Created the image",
	}
}

type ServicePatientRequestGetImage struct {
	ImageID  string `json:"image_id" validate:"required,uuid4"`
	Password string `json:"password" validate:"required"`

	PatientID string `json:"-" validate:"required,uuid4"`
}

func (s ServicePatientRequestGetImage) ToAccessHistoryModel() (model.AccessHistory, error) {
	imageID, err := uuid.Parse(s.ImageID)
	if err != nil {
		return model.AccessHistory{}, err
	}

	return model.AccessHistory{
		ID:             uuid.New(),
		PatientImageID: imageID,
		PatientID:      null.NewString(s.PatientID, true),
		DoctorID:       null.NewString("", false),
		Purpose:        "Requested to get the image",
	}, nil
}

func (s ServicePatientRequestGetImage) ComparePassword(encryptedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(s.Password))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return errors.ErrIncorrectPassword
	} else if err != nil {
		return err
	}

	return nil
}

type RepositoryFindPatientImage struct {
	ID        uuid.UUID
	PatientID null.String
	DoctorID  null.String
	IsValid   bool
}

type ServicePatientGetImage struct {
	Token string `param:"token" validate:"required"`

	PatientID string    `json:"-"`
	ImageID   uuid.UUID `json:"-"`
}

func (s ServicePatientGetImage) ToAccessHistoryModel() (model.AccessHistory, error) {
	return model.AccessHistory{
		ID:             uuid.New(),
		PatientImageID: s.ImageID,
		PatientID:      null.NewString(s.PatientID, true),
		DoctorID:       null.NewString("", false),
		Purpose:        "Downloaded the image",
	}, nil
}

type RepositoryInsertRequestPatientImageToken struct {
	PatientID     string `redis:"-"`
	ValidInMinute int    `redis:"-"`

	ImageID  string `redis:"image_id"`
	Password string `redis:"password"`
	Token    string `redis:"token"`
}

type RepositoryFindRequestPatientImageToken struct {
	PatientID uuid.UUID `redis:"-"`
}

type ServiceGivingPermission struct {
	DoctorID string `json:"doctor_id" validate:"required,uuid4"`
	ImageID  string `json:"image_id" validate:"required,uuid4"`
	Password string `json:"password" validate:"required"`

	PatientID string `json:"-"`
}

func (s ServiceGivingPermission) ToAccessRequestModel(allowedUntilInDays int) (model.AccessRequest, error) {
	patientID, err := uuid.Parse(s.PatientID)
	if err != nil {
		return model.AccessRequest{}, err
	}

	doctorID, err := uuid.Parse(s.DoctorID)
	if err != nil {
		return model.AccessRequest{}, err
	}

	imageID, err := uuid.Parse(s.ImageID)
	if err != nil {
		return model.AccessRequest{}, err
	}

	return model.AccessRequest{
		ID:           uuid.New(),
		PatientID:    patientID,
		DoctorID:     doctorID,
		ImageID:      imageID,
		Purpose:      "Giving permission to access the image",
		IsAllowed:    null.BoolFrom(true),
		AllowedUntil: time.Now().AddDate(0, 0, allowedUntilInDays),
	}, nil
}

func (s ServiceGivingPermission) CompareHashAndPassword(encryptedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(s.Password))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return errors.ErrIncorrectPassword
	} else if err != nil {
		return err
	}

	return nil
}

type RepositoryInsertRequestToRedis struct {
	RequestID       string `redis:"-"`
	KeepAliveInDays int    `redis:"-"`

	Password string `redis:"password"`
}
