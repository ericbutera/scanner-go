package rest

import (
	"log"
	appconfig "storage-api/config" // TODO: ericbutera/scanner-go/storage-api
	_ "storage-api/docs"
	"time"

	// "github.com/gin-contrib/requestid"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
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
	profile(config, r)

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

func profile(config appconfig.AppConfig, r *gin.Engine) {
	if !config.DataDog {
		log.Print("DataDog is not enabled")
		return
	}

	log.Println("DataDog is enabled")
	tracer.Start(
		tracer.WithProfilerCodeHotspots(true),
		tracer.WithProfilerEndpoints(true),
		tracer.WithServiceName(config.ServiceName),
		tracer.WithEnv(config.Env),
		tracer.WithServiceVersion(config.Version),
		tracer.WithGlobalTag("app", config.AppName),
		tracer.WithAnalytics(true),
	)
	defer tracer.Stop()

	err := profiler.Start(
		profiler.WithAPIKey(config.DataDogApiKey),
		profiler.WithService(config.ServiceName),
		profiler.WithEnv(config.Env),
		profiler.WithVersion(config.Version),
		profiler.WithTags("api"),
		profiler.WithProfileTypes(profiler.CPUProfile, profiler.HeapProfile), //profiler.BlockProfile, profiler.MutexProfile

	)
	if err != nil {
		log.Fatal(err)
	}

	r.Use(gintrace.Middleware(config.AppName))
}

func routes(r *gin.Engine, app *App) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/", app.Docs)
	r.GET("/health", app.Health)
	r.POST("/scan/start", app.Start)
	r.POST("/scan/:id/save", app.Save)
	r.POST("/scan/:id/finish", app.Finish)

}
