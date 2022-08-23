package profile

import (
	"log"
	"scanner-go/config"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
)

func Profile(conf *config.AppConfig) {
	DataDog(conf)
}

func DataDog(conf *config.AppConfig) {
	if !conf.DataDog {
		return
	}

	log.Println("Starting profiler")
	tracer.Start(
		tracer.WithProfilerCodeHotspots(true),
		tracer.WithProfilerEndpoints(true),
		tracer.WithServiceName(conf.ServiceName),
		tracer.WithEnv(conf.Env),
		tracer.WithServiceVersion(conf.Version),
		tracer.WithGlobalTag("app", conf.AppName),
		tracer.WithAnalytics(true),
	)
	defer tracer.Stop()

	err := profiler.Start(
		profiler.WithAPIKey(conf.DataDogApiKey),
		profiler.WithService(conf.ServiceName),
		profiler.WithEnv(conf.Env),
		profiler.WithVersion(conf.Version),
		profiler.WithTags("cli"),
		profiler.WithProfileTypes(profiler.CPUProfile, profiler.HeapProfile), //profiler.BlockProfile, profiler.MutexProfile
	)
	if err != nil {
		log.Fatal(err)
	}
	defer profiler.Stop()
	log.Printf("profiler started")
}
