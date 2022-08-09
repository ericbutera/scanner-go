package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Docs(c *gin.Context) {
	// TODO openapi
	c.JSON(http.StatusOK, gin.H{
		"message":   "alive",
		"endpoints": []string{"/health"},
	})
}

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "alive",
	})
}

func Start(c *gin.Context) {
	id := uuid.New()
	c.JSON(http.StatusOK, gin.H{
		"message": "start",
		"id":      id,
	})
}

func Save(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "save",
		"id":      id,
	})
}

func Finish(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "finish",
		"id":      id,
	})
}
