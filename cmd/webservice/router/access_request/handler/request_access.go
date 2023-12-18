package handler

import (
	"medsecurity/pkg/errors"
	"medsecurity/type/params"
	"medsecurity/utils/httpx"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h accessRequestHandler) RequestAccess() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req params.ServiceGivingPermission
		err := c.Bind(&req)
		if err != nil {
			return httpx.WriteErrorResponse(c, err, "bad request")
		}

		jwtClaim, err := httpx.GetJWTClaimsFromContext(c)
		if err != nil {
			return httpx.WriteErrorResponse(c, err, "error when getting jwt claim")
		}

		var ok bool
		req.PatientID, ok = jwtClaim["sub"].(string)
		if !ok {
			return httpx.WriteErrorResponse(c, errors.ErrInternalServer, "error when getting jwt claim")
		}

		err = h.patientImageService.GivingPermission(c.Request().Context(), req)
		if err != nil {
			return httpx.WriteErrorResponse(c, err, "error when giving permission")
		}

		return httpx.WriteResponse(c, http.StatusCreated, "created")
	}
}
