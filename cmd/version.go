package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tailwarden/komiser/internal"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const logo = `
 _  __               _
| |/ /___  _ __ ___ (_)___  ___ _ __
| ' // _ \| '_ \ _ \| / __|/ _ \ '__|
| . \ (_) | | | | | | \__ \  __/ |
|_|\_\___/|_| |_| |_|_|___/\___|_|
`

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show tool version",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(logo)
		showLine(cmd, "Version", internal.Version)
		showLine(cmd, "Go version", internal.GoVersion)
		showLine(cmd, "Commit", internal.Commit)
		showLine(cmd, "OS/Arch", fmt.Sprintf("%s/%s", internal.Os, internal.Arch))
		showLine(cmd, "Built", internal.Buildtime)
	},
}

func showLine(cmd *cobra.Command, title string, value string) {
	cmd.Printf(" %-16s %s\n", fmt.Sprintf("%s:", cases.Title(language.Und, cases.NoLower).String(title)), value)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
