package cmd

import (
	"fmt"
	"gocoop/internal"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:               "gocoop",
	Short:             "GoCoop helps you manage your chicken coop",
	PersistentPreRunE: persistentPreRunE,
	Run:               internal.Run,
}

func persistentPreRunE(cmd *cobra.Command, args []string) error {
	logging, err := cmd.Flags().GetString("logging")
	if err != nil {
		return fmt.Errorf("error with the logging flag: %s", err)
	}

	// Parse the logging level
	level, err := logrus.ParseLevel(logging)
	if err != nil {
		return fmt.Errorf("error while parsing the logging level: %s", err)
	}

	// Set
	logrus.SetLevel(level)

	// Set the TextFormatter
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
	})

	return nil
}

func init() {
	rootCmd.PersistentFlags().StringP("logging", "l", "info", "Logging level")
	rootCmd.Flags().StringP("config", "c", "config.yml", "Configuration file")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
