package cmd

import (
	"errors"

	"github.com/mlabouardy/komiser/internal"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Run Komiser server",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
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

		noTracking, _ := cmd.Flags().GetBool("no-tracking")
		if noTracking {
			log.Info("Tracking has been disabled")
		}

		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			return err
		}

		err = internal.Exec(port, file, noTracking, regions, cmd)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.PersistentFlags().Int("port", 3000, `Port to start server on, default:"3000".`)
	startCmd.PersistentFlags().StringArray("regions", []string{}, "Restrict Komiser inspection to list of regions.")
	startCmd.PersistentFlags().String("config", "config.toml", "Path to configuration file.")
	startCmd.PersistentFlags().Bool("verbose", true, "Show verbose debug information.")
	startCmd.PersistentFlags().Bool("no-tracking", false, "Disable user analytics.")
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
