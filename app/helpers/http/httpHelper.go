package http

import "github.com/labstack/echo/v4"

type genericResponse struct {
	Error   bool        `json:"error"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Ok(c echo.Context, httpStatus int, msg string, data interface{}) error {
	response := genericResponse{
		Error:   false,
		Status:  httpStatus,
		Message: msg,
		Data:    data,
	}

	return c.JSON(httpStatus, response)
}

func Error(c echo.Context, httpStatus int, err error) error {
	response := genericResponse{
		Error:   true,
		Status:  httpStatus,
		Message: err.Error(),
	}

	return c.JSON(httpStatus, response)
}
