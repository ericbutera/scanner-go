// TODO:
// https://github.com/axiaoxin-com/ratelimiter
// https://github.com/gin-gonic/contrib

package rest

import (
	"log"
	appconfig "storage-api/config" // TODO: ericbutera/scanner-go/storage-api
	_ "storage-api/docs"
	"time"

	// "github.com/gin-contrib/requestid"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type App struct {
	Logger *zap.Logger
	Sugar  *zap.SugaredLogger
}

// @title Storage API
// @BasePath /
func Serve(config appconfig.AppConfig) {

	r := gin.Default()
	// TODO: gin.SetMode(gin.DebugMode)

	app := &App{}

	// r.Use(requestid.New())
	logger(config, r, app)
	initDataDog(config, r)
	initOpenTel(config, r)

	app.Logger.Info("Starting server")
	app.Sugar.Infow("starting",
		"app name", config.AppName,
		"version", config.Version,
	)

	routes(r, app)

	r.Run()
}

func logger(config appconfig.AppConfig, r *gin.Engine, app *App) {
	// TODO: NewProduction()
	logger, err := zap.NewDevelopment()
	defer logger.Sync()

	if err != nil {
		log.Fatalf("log error %v", err)
	}

	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	app.Logger = logger
	app.Sugar = logger.Sugar()
}

func routes(r *gin.Engine, app *App) {
	r.GET("/swagger/*any", swagger())
	r.GET("/", app.Docs)
	r.GET("/health", app.Health)
	r.POST("/scan/start", app.Start)
	r.POST("/scan/:id/save", app.Save)
	r.POST("/scan/:id/finish", app.Finish)

}
