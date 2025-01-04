package utils

import (
	"net/http"

	"github.com/eavlongs/go_backend_template/constants"
	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Errors  map[string]any `json:"errors"`
}

func RespondWithSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, SuccessResponse{
		Success: true,
		Message: constants.RequestSuccessful,
		Data:    data,
	})
}

func RespondWithSuccessAndMessage(ctx *gin.Context, message string, data interface{}) {
	ctx.JSON(http.StatusOK, SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func RespondWithError(ctx *gin.Context, status int, message string, validationError map[string]any) {
	ctx.JSON(status, ErrorResponse{
		Success: false,
		Message: message,
		Errors:  validationError,
	})
}

func RespondWithNotFoundError(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, ErrorResponse{
		Success: false,
		Message: constants.NotFoundError,
	})
}

func RespondWithUnauthorizedError(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, ErrorResponse{
		Success: false,
		Message: constants.UnauthorizedError,
	})
}

func RespondWithInternalServerError(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, ErrorResponse{
		Success: false,
		Message: constants.InternalServerError,
	})
}

func RespondWithBadRequestError(ctx *gin.Context, message string, validationError map[string]any) {
	ctx.JSON(http.StatusBadRequest, ErrorResponse{
		Success: false,
		Message: message,
		Errors:  validationError,
	})
}
