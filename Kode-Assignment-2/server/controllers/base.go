package controllers

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string      `json:"message" example:"GET_SUCCESS"`
	Status  int         `json:"status" example:"200"`
	Payload interface{} `json:"payload,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func WriteJsonResponse(ctx *gin.Context, payload *Response) {
	ctx.JSON(payload.Status, payload)
}
