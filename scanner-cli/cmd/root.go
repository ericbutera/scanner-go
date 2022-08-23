// TODO: https://github.com/spf13/cobra/blob/main/user_guide.md
package cmd

import (
	"log"

	"scanner-go/cmd/command"
	"scanner-go/config"

	"github.com/spf13/cobra"
)

// var (
// 	globalFlags = command.GlobalFlags{}
// )

var (
	conf    = &config.AppConfig{}
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "scan",
		Short: "Scan for resources",
		Long:  "Scan for resources",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "app.yml", "config file")
	rootCmd.AddCommand(command.NewScan(conf))
}

func initConfig() {
	err := config.Load(cfgFile, conf)
	if err != nil {
		log.Fatal(err)
	}
}
