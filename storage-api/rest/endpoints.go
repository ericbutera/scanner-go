// TODO: https://github.com/swaggo/swag/blob/master/README.md#declarative-comments-format
// example: https://github.com/swaggo/swag/blob/master/example/celler/controller/accounts.go
// examples: https://github.com/swaggo/swag/tree/master/example

package rest

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const swaggerUri = "/swagger/index.html"

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
	c.Redirect(http.StatusMovedPermanently, swaggerUri)
}

type HealthResponse struct {
	Base
}

// Health check
// @Summary Health check
// @Schemes
// @Description Health check
// @Accept json
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /health [get]
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, HealthResponse{
		Base: Base{Message: "alive"},
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
		return
	}

	c.JSON(http.StatusOK, SaveResponse{
		Base:  Base{Message: "Data saved"},
		JobId: uuid.New(),
	})
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
}
