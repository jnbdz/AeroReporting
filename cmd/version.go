package cmd

import (
	"fmt"
	"github.com/jnbdz/AeroReporting/config"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of " + config.AppName,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(config.AppName + " version " + config.Version)
	},
}
