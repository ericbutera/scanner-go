package rest

import (
	"github.com/gin-gonic/gin"
)

func Serve() {
	r := gin.Default()

	r.GET("/", Docs)
	r.GET("/health", Health)
	r.POST("/scan/start", Start)
	r.POST("/scan/:id/save", Save)
	r.POST("/scan/:id/finish", Finish)

	// listen and serve on 0.0.0.0:8080
	r.Run()
}
