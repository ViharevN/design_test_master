package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	BadRequest    = http.StatusBadRequest
	InternalError = http.StatusInternalServerError
)

func ApiInternalError(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(InternalError, gin.H{"error": err.Error()})
}

func ApiBadRequestError(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(BadRequest, gin.H{"error": err.Error()})
}

func ApiBadRequesErrorWithMessage(ctx *gin.Context, message string) {
	ctx.AbortWithStatusJSON(BadRequest, gin.H{"error": message})
}
