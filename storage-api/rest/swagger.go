package rest

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func swagger() gin.HandlerFunc {
	return ginSwagger.WrapHandler(swaggerfiles.Handler)
}
