package handlers

import "github.com/gin-gonic/gin"

type IHandler interface {
	Handle(context *gin.Context)
}
