package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/linux-http/env"
)

var Env env.Settings = env.NewEnv()

type Response struct {
	Status  int
	Message string
	Data    map[string]any
}

type PanicResponse struct {
	HttpStatus int
	Response   Response
}

func PanicMiddleware(ctx *gin.Context, err any) {
	switch v := err.(type) {
	case PanicResponse:
		ctx.AbortWithStatusJSON(v.HttpStatus, v.Response)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, Response{
		Status:  500,
		Message: fmt.Sprintf("Internal Server Error %v", err),
	})
}

func NewPanic(httpStatus, status int, mesg string, data ...map[string]any) PanicResponse {
	output := map[string]any{}
	if len(data) > 0 {
		output = data[0]
	}
	return PanicResponse{
		HttpStatus: httpStatus,
		Response: Response{
			Status:  status,
			Message: mesg,
			Data:    output,
		},
	}
}
