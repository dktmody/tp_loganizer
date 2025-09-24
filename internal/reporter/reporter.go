package reporter

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Report []interface{}

func ExportReport(report interface{}, filePath string) error {
	// Créer les répertoires nécessaires si ils n'existent pas
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(report)
}
