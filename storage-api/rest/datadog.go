package rest

import (
	"log"
	appconfig "storage-api/config"

	"github.com/gin-gonic/gin"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
)

func initDataDog(config appconfig.AppConfig, r *gin.Engine) {
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
