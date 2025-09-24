package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type LogConfig struct {
	ID   string `json:"id"`
	Path string `json:"path"`
	Type string `json:"type"`
}

var (
	logID      string
	logPath    string
	logType    string
	configFile string
)

var addLogCmd = &cobra.Command{
	Use:   "add-log",
	Short: "Ajouter une configuration de log au fichier config.json",
	Run: func(cmd *cobra.Command, args []string) {
		// Charger le fichier de configuration existant
		file, err := os.OpenFile(configFile, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			fmt.Printf("Erreur lors de l'ouverture du fichier : %v\n", err)
			return
		}
		defer file.Close()

		var configs []LogConfig
		if err := json.NewDecoder(file).Decode(&configs); err != nil && err.Error() != "EOF" {
			fmt.Printf("Erreur lors du décodage du fichier JSON : %v\n", err)
			return
		}

		// Ajouter la nouvelle configuration
		newConfig := LogConfig{
			ID:   logID,
			Path: logPath,
			Type: logType,
		}
		configs = append(configs, newConfig)

		// Réécrire le fichier avec la nouvelle configuration
		file.Seek(0, 0)
		file.Truncate(0)
		if err := json.NewEncoder(file).Encode(configs); err != nil {
			fmt.Printf("Erreur lors de l'écriture dans le fichier JSON : %v\n", err)
			return
		}

		fmt.Println("Configuration ajoutée avec succès !")
	},
}

func init() {
	rootCmd.AddCommand(addLogCmd)
	addLogCmd.Flags().StringVar(&logID, "id", "", "Identifiant du log")
	addLogCmd.Flags().StringVar(&logPath, "path", "", "Chemin du fichier de log")
	addLogCmd.Flags().StringVar(&logType, "type", "", "Type de log")
	addLogCmd.Flags().StringVar(&configFile, "file", "config.json", "Chemin du fichier de configuration")
	addLogCmd.MarkFlagRequired("id")
	addLogCmd.MarkFlagRequired("path")
	addLogCmd.MarkFlagRequired("type")
}
