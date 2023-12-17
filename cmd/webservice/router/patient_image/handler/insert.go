package handler

import (
	"medsecurity/pkg/errors"
	"medsecurity/type/constant"
	"medsecurity/type/params"
	"medsecurity/utils/httpx"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h patientImageHandler) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		var param params.ServiceCreatePatientImage
		err := c.Bind(&param)
		if err != nil {
			return httpx.WriteErrorResponse(c, errors.ErrBadRequest, "bad request")
		}

		role, err := httpx.GetRoleAccountFromContext(c)
		if err != nil {
			return httpx.WriteErrorResponse(c, err, "error at parsing role")
		}

		if role != constant.DoctorRole {
			return httpx.WriteErrorResponse(c, errors.ErrUnauthorized, "you are not authorized to access this resource")
		}

		jwtClaim, err := httpx.GetJWTClaimsFromContext(c)
		if err != nil {
			return httpx.WriteErrorResponse(c, err, nil)
		}

		var ok bool
		param.DoctorID, ok = jwtClaim["sub"].(string)
		if !ok {
			return httpx.WriteErrorResponse(c, errors.ErrInternalServer, nil)
		}

		err = h.patientImageService.Insert(c.Request().Context(), param)
		if err != nil {
			return httpx.WriteErrorResponse(c, err, nil)
		}

		return httpx.WriteResponse(c, http.StatusCreated, "created")
	}
}
