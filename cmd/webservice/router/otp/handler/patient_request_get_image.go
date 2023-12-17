package handler

import (
	"medsecurity/pkg/errors"
	"medsecurity/type/params"
	"medsecurity/utils/httpx"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h otpHandler) PatientRequestGetImage(c echo.Context) error {
	var req params.ServicePatientRequestGetImage
	err := c.Bind(&req)
	if err != nil {
		return httpx.WriteErrorResponse(c, err, nil)
	}

	jwtClaim, err := httpx.GetJWTClaimsFromContext(c)
	if err != nil {
		return httpx.WriteErrorResponse(c, err, nil)
	}

	var ok bool
	req.PatientID, ok = jwtClaim["sub"].(string)
	if !ok {
		return httpx.WriteErrorResponse(c, errors.ErrInternalServer, nil)
	}

	res, err := h.patientImageService.PatientRequestGetImage(c.Request().Context(), req)
	if err != nil {
		return httpx.WriteErrorResponse(c, err, nil)
	}

	return httpx.WriteResponse(c, http.StatusCreated, res)
}
