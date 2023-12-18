package patient_image

import (
	"context"
	"medsecurity/pkg/errors"
	"medsecurity/type/params"
	"medsecurity/type/result"
	"medsecurity/utils/aesx"
	"medsecurity/utils/rsax"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/volatiletech/null/v9"
)

func (s service) PatientRequestGetImage(ctx context.Context, param params.ServicePatientRequestGetImage) (result.ServicePatientRequestGetImage, error) {
	err := s.validator.Validate(param)
	if err != nil {
		return result.ServicePatientRequestGetImage{}, err
	}

	imageID, err := uuid.Parse(param.ImageID)
	if err != nil {
		return result.ServicePatientRequestGetImage{}, errors.Wrap(err, "error when parsing image id")
	}

	patientID, err := uuid.Parse(param.PatientID)
	if err != nil {
		return result.ServicePatientRequestGetImage{}, errors.Wrap(err, "error when parsing patient id")
	}

	patient, err := s.patientRepository.Find(ctx, params.RepoFindPatient{
		ID: null.NewString(param.PatientID, true),
	})
	if errors.Is(err, errors.ErrRecordNotFound) {
		return result.ServicePatientRequestGetImage{}, errors.ErrAccountNotFound
	} else if err != nil {
		return result.ServicePatientRequestGetImage{}, errors.Wrap(err, "error when finding patient")
	}

	err = param.ComparePassword(patient.Password)
	if errors.Is(err, errors.ErrIncorrectPassword) {
		return result.ServicePatientRequestGetImage{}, errors.ErrIncorrectPassword
	} else if err != nil {
		return result.ServicePatientRequestGetImage{}, err
	}

	patientImage, err := s.patientImageRepository.Find(ctx, params.RepositoryFindPatientImage{
		ID:        imageID,
		PatientID: null.NewString(param.PatientID, true),
		IsValid:   true,
	})
	if errors.Is(err, errors.ErrRecordNotFound) {
		return result.ServicePatientRequestGetImage{}, errors.ErrRecordNotFound
	} else if err != nil {
		return result.ServicePatientRequestGetImage{}, errors.Wrap(err, "error when finding patient image")
	}

	uuidToken := strings.Replace(uuid.New().String(), "-", "", -1)
	token := uuidToken[:10]
	validInMinute := 5

	encryptedPassword, err := aesx.Encrypt([]byte(s.config.AES.Secret), []byte(param.Password))
	if err != nil {
		return result.ServicePatientRequestGetImage{}, errors.Wrap(err, "error when encrypting password")
	}

	err = s.patientImageRepository.InsertPatientRequestGetImageToken(ctx, params.RepositoryInsertRequestPatientImageToken{
		PatientID:     patientID.String(),
		ImageID:       patientImage.ID.String(),
		ValidInMinute: validInMinute,
		Password:      string(encryptedPassword),
		Token:         token,
	})
	if err != nil {
		return result.ServicePatientRequestGetImage{}, errors.Wrap(err, "error when inserting token")
	}

	accesHistoryModel, err := param.ToAccessHistoryModel()
	if err != nil {
		return result.ServicePatientRequestGetImage{}, errors.Wrap(err, "error when converting to access history model")
	}

	err = s.accessHistoryRepository.Insert(ctx, accesHistoryModel)
	if err != nil {
		return result.ServicePatientRequestGetImage{}, errors.Wrap(err, "error when inserting access history")
	}

	return result.ServicePatientRequestGetImage{
		Token:     token,
		ExpiredAt: time.Now().Add(time.Minute * time.Duration(validInMinute)).Unix(),
	}, nil
}

func (s service) PatientGetImage(ctx context.Context, param params.ServicePatientGetImage) (result.ServicePatientGetImage, error) {
	err := s.validator.Validate(param)
	if err != nil {
		return result.ServicePatientGetImage{}, err
	}

	patientID, err := uuid.Parse(param.PatientID)
	if err != nil {
		return result.ServicePatientGetImage{}, errors.Wrap(err, "error when parsing patient id")
	}

	cache, err := s.patientImageRepository.FindPatientRequestGetImageToken(ctx, params.RepositoryFindRequestPatientImageToken{
		PatientID: patientID,
	})
	if errors.Is(err, errors.ErrRecordNotFound) {
		return result.ServicePatientGetImage{}, errors.ErrRecordNotFound
	} else if err != nil {
		return result.ServicePatientGetImage{}, errors.Wrap(err, "error when finding token")
	}

	if cache.Token != param.Token {
		return result.ServicePatientGetImage{}, errors.ErrUnauthorized
	}

	decryptedPassword, err := aesx.Decrypt([]byte(s.config.AES.Secret), []byte(cache.Password))
	if err != nil {
		return result.ServicePatientGetImage{}, errors.Wrap(err, "error when decrypting password")
	}

	_, err = s.patientRepository.Find(ctx, params.RepoFindPatient{
		ID: null.NewString(param.PatientID, true),
	})
	if errors.Is(err, errors.ErrRecordNotFound) {
		return result.ServicePatientGetImage{}, errors.ErrAccountNotFound
	} else if err != nil {
		return result.ServicePatientGetImage{}, errors.Wrap(err, "error when finding patient")
	}

	patientSecret, err := s.patientSecretRepository.FindByPatientID(ctx, patientID)
	if errors.Is(err, errors.ErrRecordNotFound) {
		return result.ServicePatientGetImage{}, errors.ErrRecordNotFound
	} else if err != nil {
		return result.ServicePatientGetImage{}, errors.Wrap(err, "error when finding patient secret")
	}

	decryptedSalt, err := aesx.Decrypt([]byte(s.config.AES.Secret), []byte(patientSecret.Salt))
	if err != nil {
		return result.ServicePatientGetImage{}, errors.Wrap(err, "error when decrypting salt")
	}

	imageID, err := uuid.Parse(cache.ImageID)
	if err != nil {
		return result.ServicePatientGetImage{}, errors.Wrap(err, "error when parsing image id")
	}

	param.ImageID = imageID

	patientImage, err := s.patientImageRepository.Find(ctx, params.RepositoryFindPatientImage{
		ID:        imageID,
		PatientID: null.NewString(param.PatientID, true),
		IsValid:   true,
	})
	if errors.Is(err, errors.ErrRecordNotFound) {
		return result.ServicePatientGetImage{}, errors.ErrRecordNotFound
	} else if err != nil {
		return result.ServicePatientGetImage{}, errors.Wrap(err, "error when finding patient image")
	}

	image, err := s.cloudinaryRepository.DownloadFile(ctx, patientImage.URL)
	if errors.Is(err, errors.ErrNotFound) {
		return result.ServicePatientGetImage{}, errors.ErrNotFound
	} else if err != nil {
		return result.ServicePatientGetImage{}, errors.Wrap(err, "error when downloading image")
	}

	privateKey, err := rsax.DecryptPrivateKey(patientSecret.PrivateKey, string(decryptedPassword), string(decryptedSalt))
	if err != nil {
		return result.ServicePatientGetImage{}, errors.Wrap(err, "error when decrypting private key")
	}

	match, err := patientSecret.IsPrivateKeyPublicKeyMatch(privateKey)
	if err != nil {
		return result.ServicePatientGetImage{}, errors.Wrap(err, "error when checking private key and public key match")
	} else if !match {
		return result.ServicePatientGetImage{}, errors.ErrInternalServer
	}

	res := result.ServicePatientGetImage{
		DocumentName: patientImage.Name,
		DocumentType: patientImage.Type,
	}
	err = res.DecryptImage(privateKey, patientSecret.KeySize, image)
	if err != nil {
		return result.ServicePatientGetImage{}, errors.Wrap(err, "error when decrypting image")
	}

	accessHistoryModel, err := param.ToAccessHistoryModel()
	if err != nil {
		return result.ServicePatientGetImage{}, errors.Wrap(err, "error when converting to access history model")
	}

	err = s.accessHistoryRepository.Insert(ctx, accessHistoryModel)
	if err != nil {
		return result.ServicePatientGetImage{}, errors.Wrap(err, "error when inserting access history")
	}

	return res, nil
}
