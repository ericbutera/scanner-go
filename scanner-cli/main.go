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
	app_conf, err := config.NewAppConfig(config_path)
	if err != nil {
		log.Fatal(err)
	}

	profile(app_conf)

	client := storage.NewClient()

	scan_id, err := storage.Start(client)
	if err != nil {
		log.Fatal(err)
	}

	span := tracer.StartSpan("scanner")
	defer span.Finish()

	gcp.Scan(app_conf)
	aws.Scan(app_conf)
	azure.Scan(app_conf)

	span.Finish(tracer.WithError(err))

	storage.End(client, scan_id)
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
	)
	defer tracer.Stop()

	err := profiler.Start(
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
