package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "loganalyzer",
	Short: "LogAnalyzer is a CLI tool for analyzing logs",
	Long:  `LogAnalyzer helps system administrators analyze log files from various sources in parallel.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// Here you will define your flags and configuration settings.
}