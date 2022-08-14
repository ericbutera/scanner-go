// TODO: https://github.com/swaggo/swag/blob/master/README.md#declarative-comments-format
// example: https://github.com/swaggo/swag/blob/master/example/celler/controller/accounts.go

package rest

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Base struct {
	Message string `json:"message"`
}

type StartResponse struct {
	Base
	Id uuid.UUID `json:"id"`
}

type ErrorResponse struct {
	Base
	Code int
}

type SaveRequest struct {
	ScanId string
	Data   any `required:"true" description:"Raw vendor data to save"`
}

type SaveResponse struct {
	Base
	JobId uuid.UUID `description:"Unique identifier for the saved data" json:"jobId"`
}

func Docs(c *gin.Context) {
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

// Start a scan session
// @Summary Start a scan session
// @Schemes
// @Description Start a scan session
// @Accept json
// @Produce json
// @Success 200 {object} StartResponse
// @Router /scan/start [post]
func Start(c *gin.Context) {
	id := uuid.New()
	c.JSON(http.StatusOK, &StartResponse{
		Base: Base{Message: "Session started"},
		Id:   id,
	})
}

// Save scan data
// @Summary Save scan data
// @Schemes
// @Description Save scan data
// @Accept json
// @Produce json
// @Param 	id 		path 	string 		true 	"Scan ID"
// @Param 	data 	body 	SaveRequest true 	"Scan data"
// @Success 200 {object} Base
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /scan/{id}/save [post]
func Save(c *gin.Context) {
	id := c.Param("id")
	var data SaveRequest

	log.Printf("saving scan %s %+v", id, data)

	// TODO: 404 id not found

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Base: Base{Message: "Invalid request"},
		})
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"message": "error parsing json",
		// 	"id":      id,
		// })
		return
	}

	c.JSON(http.StatusOK, SaveResponse{
		Base:  Base{Message: "Data saved"},
		JobId: uuid.New(),
	})
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "save",
	// 	"id":      id,
	// 	"jobId":   uuid.New(),
	// })
}

// Finish a scan session
// @Summary Finish a scan session
// @Schemes
// @Description Finish a scan session
// @Accept json
// @Produce json
// @Success 200 {Base} Base
// @Failure 404 {object} ErrorResponse
// @Router /scan/{id}/finish [post]
func Finish(c *gin.Context) {
	id := c.Param("id")
	log.Printf("finishing scan %s", id)

	// TODO: 404 id not found

	c.JSON(http.StatusOK, Base{
		Message: "Session finished",
	})
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "finish",
	// 	"id":      id,
	// })
}
