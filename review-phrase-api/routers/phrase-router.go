package routers

import "github.com/gin-gonic/gin"

func setupRouter() {
	router := gin.Default()

	router.GET("phrases")
}