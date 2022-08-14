package main

import (
	"flag"
	"log"
	"scanner-go/config"
	"scanner-go/provider/aws"
	"scanner-go/provider/azure"
	"scanner-go/provider/gcp"
	storage "scanner-go/storage"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
)

var (
	config_path = flag.String("config_path", ".", "path to config")
)

func main() {
	conf, err := config.NewAppConfig(config_path)
	if err != nil {
		log.Fatal(err)
	}

	profile(conf)

	store := storage.NewStorage(conf.StorageApiUrl)

	if err := store.Start(); err != nil {
		log.Fatal(err)
	}

	// TODO: learn how to make a wrapper https://github.com/DataDog/dd-trace-go/blob/v1.40.1/contrib/google.golang.org/api/api.go#L47
	// span := tracer.StartSpan("scanner")
	// defer span.Finish()
	// span.Finish(tracer.WithError(err))

	if conf.Gcp {
		gcp.Scan(conf, store)
	}

	if conf.Aws {
		aws.Scan(conf, store)
	}

	if conf.Azure {
		azure.Scan(conf, store)
	}

	store.End()
}

func profile(conf config.AppConfig) {
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
		profiler.WithProfileTypes(profiler.CPUProfile, profiler.HeapProfile, profiler.BlockProfile, profiler.MutexProfile),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer profiler.Stop()
}
