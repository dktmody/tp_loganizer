package cmd

import (
	"fmt"
	"sync"
	"time"

	"github.com/dktmody/go_loganizer/internal/analyzer"
	"github.com/dktmody/go_loganizer/internal/config"
	"github.com/dktmody/go_loganizer/internal/reporter"

	"github.com/spf13/cobra"
)

var configPath string
var outputPath string

// analyzeCmd represents the analyze command
var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze log files based on a configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		// Charger la configuration
		configs, err := config.LoadConfig(configPath)
		if err != nil {
			fmt.Printf("Failed to load config: %v\n", err)
			return
		}

		// Canal pour collecter les résultats
		results := make(chan analyzer.LogResult, len(configs))
		var wg sync.WaitGroup

		// Analyser les logs de manière concurrente
		for _, logConfig := range configs {
			wg.Add(1)
			go func(cfg config.LogConfig) {
				defer wg.Done()
				result := analyzer.AnalyzeLog(cfg.ID, cfg.Path)
				results <- result
			}(logConfig)
		}

		// Attendre que toutes les goroutines se terminent
		wg.Wait()
		close(results)

		// Collecter les résultats
		var report []analyzer.LogResult
		for result := range results {
			report = append(report, result)
		}

		// Exporter le rapport
		currentTime := time.Now().Format("2006-01-02_15-04-05")
		outputPath = fmt.Sprintf("%s_%s", outputPath, currentTime)
		if err := reporter.ExportReport(report, outputPath); err != nil {
			fmt.Printf("Failed to export report: %v\n", err)
			return
		}

		fmt.Println("Analysis completed successfully!")
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
	analyzeCmd.Flags().StringVarP(&configPath, "config", "c", "", "Path to the configuration file")
	analyzeCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Path to the output report file")
	analyzeCmd.MarkFlagRequired("config")
	analyzeCmd.MarkFlagRequired("output")
}
