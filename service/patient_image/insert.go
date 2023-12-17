package patient_image

import (
	"bytes"
	"context"
	"medsecurity/pkg/errors"
	"medsecurity/type/params"

	"github.com/google/uuid"
	"github.com/volatiletech/null/v9"
)

func (s service) Insert(ctx context.Context, param params.ServiceCreatePatientImage) error {
	err := s.validator.Validate(param)
	if err != nil {
		return err
	}

	patientID, err := uuid.Parse(param.PatientID)
	if err != nil {
		return errors.ErrInternalServer
	}

	patientSecret, err := s.patientSecretRepository.FindByPatientID(ctx, patientID)
	if errors.Is(err, errors.ErrRecordNotFound) {
		return errors.ErrAccountNotFound
	} else if err != nil {
		return errors.Wrap(err, "error at find patient secret")
	}

	_, err = s.doctorRepository.Find(ctx, params.RepoFindDoctor{
		ID: null.NewString(param.DoctorID, true),
	})
	if errors.Is(err, errors.ErrRecordNotFound) {
		return errors.ErrDoctorNotFound
	} else if err != nil {
		return errors.Wrap(err, "error at find doctor")
	}

	encryptedDocument, err := param.EncryptBase64Document(patientSecret.PublicKey, patientSecret.KeySize)
	if err != nil {
		return errors.Wrap(err, "error at encrypt base64 document")
	}

	byteReader := bytes.NewReader(encryptedDocument)
	uploadResult, err := s.cloudinaryRepository.UploadEncryptedFile(ctx, byteReader)
	if err != nil {
		return errors.Wrap(err, "error at upload encrypted file")
	}

	patientImage, err := param.ToPatientImageModel(uploadResult.SecureURL)
	if err != nil {
		return errors.Wrap(err, "error at convert to patient image model")
	}

	tx, err := s.patientImageRepository.Insert(ctx, patientImage)
	if err != nil {
		err := s.cloudinaryRepository.Remove(ctx, uploadResult.PublicID)
		if err != nil {
			return errors.Wrap(err, "error at remove encrypted file")
		}
		tx.Rollback()
		return errors.Wrap(err, "error at insert patient image")
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "error at commit transaction")
	}

	return nil
}
