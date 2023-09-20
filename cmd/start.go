package cmd

import (
	"errors"
	"time"

	"github.com/getsentry/sentry-go"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tailwarden/komiser/internal"
	"github.com/tailwarden/komiser/utils"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Run Komiser server",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		defer func() {
			err := recover()
			if err != nil {
				sentry.CurrentHub().Recover(err)
				sentry.Flush(time.Second * 5)
			}
		}()

		file, err := cmd.Flags().GetString("config")
		if err != nil {
			return err
		}
		if file == "" {
			return errors.New("you must specify a config file with '--config PATH'")
		}

		regions, err := cmd.Flags().GetStringArray("regions")
		if err != nil {
			return err
		}

		verbose, _ := cmd.Flags().GetBool("verbose")
		setupLogging(verbose)

		analytics := utils.Analytics{}

		telemetry, _ := cmd.Flags().GetBool("telemetry")
		if !telemetry {
			log.Info("Telemetry has been disabled")
		} else {
			analytics.Init()
		}

		address, err := cmd.Flags().GetString("listen-address")
		if err != nil {
			return err
		}

		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			return err
		}

		err = internal.Exec(address, port, file, telemetry, analytics, regions, cmd)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.PersistentFlags().StringP("listen-address", "a", "0.0.0.0", `Listen address to start server on.`)
	startCmd.PersistentFlags().Int("port", 3000, `Port to start server on, default:"3000".`)
	startCmd.PersistentFlags().StringArray("regions", []string{}, "Restrict Komiser inspection to list of regions.")
	startCmd.PersistentFlags().String("config", "config.toml", "Path to configuration file.")
	startCmd.PersistentFlags().Bool("verbose", true, "Show verbose debug information.")
	startCmd.PersistentFlags().Bool("telemetry", true, "Disable user analytics.")
}

func setupLogging(verbose bool) {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	if verbose {
		log.SetLevel(log.DebugLevel)
		log.Info("Debug logging is enabled")
	} else {
		log.SetLevel(log.InfoLevel)
	}
}
