package httpx

import (
	"errors"
	"log"
	"net/http"
	"strings"

	x "medsecurity/pkg/errors"
	httppkg "medsecurity/pkg/http"
	"medsecurity/type/constant"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func WriteResponse(c echo.Context, code int, data interface{}) error {
	if data == nil {
		data = http.StatusText(code)
	}

	err := c.JSON(code, httppkg.HTTPBaseResponse{
		Error: nil,
		Data:  data,
	})
	if err != nil {
		log.Println("[WriteResponse] FATAL ERROR on send response to client:", err)
		return err
	}

	return nil
}

func WriteErrorResponse(c echo.Context, errParam error, detail interface{}) error {
	e := httppkg.GetResponseErr(errParam)

	if x.Is(errParam, x.ErrValidation) {
		e.Message = x.ErrBadRequest.Error()
		e.HTTPErrorCode = echo.ErrBadRequest.Code

		// To getting the Unwrap method from private object joinError in "errors" package
		var joinErr interface{ Unwrap() []error }
		if errors.As(errParam, &joinErr) {
			errs := joinErr.Unwrap()[1].Error()
			detail = strings.Split(errs, "\n --- ")[1:]
		}
	} else if !x.Is(errParam, x.ErrInternalServer) {
		x.ErrorStack(errParam)
		detail = nil
	} else {
		detail = nil
	}

	err := c.JSON(e.HTTPErrorCode, httppkg.HTTPBaseResponse{
		Error: &httppkg.HTTPErrorBaseResponse{
			Message: e.Message,
			Detail:  detail,
		},
		Data: nil,
	})

	if err != nil {
		log.Println("[WriteErrorResponse] FATAL ERROR on send response to client:", err)
		return err
	}

	return nil
}

func GetJWTClaimsFromContext(c echo.Context) (jwt.MapClaims, error) {
	token, ok := c.Get(constant.ContextKey).(*jwt.Token)
	if !ok {
		return nil, errors.New("cannot get jwt token from context")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("cannot get jwt claims from context")
	}

	return claims, nil
}

func GetRoleAccountFromContext(c echo.Context) (string, error) {
	claims, ok := c.Get(constant.ContextKeyRole).(string)
	if !ok {
		return "", errors.New("cannot get jwt claims from context")
	}

	return claims, nil
}
