package rest

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SaveRequest struct {
	ScanId string
	Data   any
}

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
	var data SaveRequest

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error parsing json",
			"id":      id,
		})
		return
	}

	log.Printf("saving scan %s %+v", id, data)

	c.JSON(http.StatusOK, gin.H{
		"message": "save",
		"id":      id,
		"jobId":   uuid.New(),
	})
}

func Finish(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "finish",
		"id":      id,
	})
}
