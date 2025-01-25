package main

import (
	"github.com/jnbdz/AeroReporting/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     config.AppName + "-cli",
	Short:   config.AppTitleName + " CLI application",
	Version: config.Version,
}
