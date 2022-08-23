package command

import (
	"log"
	"scanner-go/config"
	"scanner-go/profile"
	"scanner-go/provider/aws"
	"scanner-go/provider/azure"
	"scanner-go/provider/gcp"
	"scanner-go/storage"

	"github.com/spf13/cobra"
)

func NewScan(conf *config.AppConfig) *cobra.Command {
	return &cobra.Command{
		Use:   "scan",
		Short: "scan for resources",
		Long:  "scan for resources",
		Run: func(cmd *cobra.Command, args []string) {
			RunScan(cmd, args, conf)
		},
	}
}

func RunScan(cmd *cobra.Command, args []string, conf *config.AppConfig) {
	log.Print("scan command!")

	profile.Profile(conf)

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
