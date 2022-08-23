// TODO: https://github.com/spf13/cobra/blob/main/user_guide.md
package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	testFlagName string

	rootCmd = &cobra.Command{
		Use:   "scan",
		Short: "Scan for resources",
		Long:  "Scan for resources",
	}

	scanCmd = &cobra.Command{
		Use:   "scan",
		Short: "scan for resources",
		Long:  "scan for resources",
		Run: func(cmd *cobra.Command, args []string) {
			log.Print("scan command!")
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(scanCmd)
}

func initConfig() {
}
