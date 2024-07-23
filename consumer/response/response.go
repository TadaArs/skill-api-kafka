package response

import (
	"net/http"
	"consumer/errs"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func Success(ctx *gin.Context, statusCode int, data any) {
	ctx.JSON(statusCode, Response{
		Status: "Success",
		Data:   data,
	})
}

func Error(ctx *gin.Context, err error) {
	switch e := err.(type) {
	case errs.Err:
		ctx.JSON(e.StatusCode, Response{
			Status:  "error",
			Message: e.Message,
		})
	case error:
		ctx.JSON(http.StatusInternalServerError, Response{
			Status:  "error",
			Message: err.Error(),
		})

	}
}
