package rest

import (
	"log"
	appconfig "storage-api/config" // TODO: ericbutera/scanner-go/storage-api

	"github.com/gin-gonic/gin"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
)

func Serve(config appconfig.AppConfig) {

	r := gin.Default()

	if config.DataDog {
		profile(config)
		r.Use(gintrace.Middleware(config.AppName))
	}

	r.GET("/", Docs)
	r.GET("/health", Health)
	r.POST("/scan/start", Start)
	r.POST("/scan/:id/save", Save)
	r.POST("/scan/:id/finish", Finish)

	// listen and serve on 0.0.0.0:8080
	r.Run()
}

func profile(config appconfig.AppConfig) {
	log.Println("starting profiler")
	tracer.Start(
		tracer.WithProfilerCodeHotspots(true),
		tracer.WithProfilerEndpoints(true),
		tracer.WithServiceName(config.ServiceName),
		tracer.WithEnv(config.Env),
		tracer.WithServiceVersion(config.Version),
		tracer.WithGlobalTag("app", config.AppName),
	)
	defer tracer.Stop()

	err := profiler.Start(
		profiler.WithService(config.ServiceName),
		profiler.WithEnv(config.Env),
		profiler.WithVersion(config.Version),
		profiler.WithTags("api"),
		profiler.WithProfileTypes(profiler.CPUProfile, profiler.HeapProfile, profiler.BlockProfile, profiler.MutexProfile),
	)
	if err != nil {
		log.Fatal(err)
	}
}
