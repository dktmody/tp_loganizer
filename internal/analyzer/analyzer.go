package analyzer

import (
	"math/rand"
	"os"
	"time"
)

type LogResult struct {
	LogID       string
	FilePath    string
	Status      string
	Message     string
	ErrorDetail string
}

func AnalyzeLog(id, path string) LogResult {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return LogResult{
			LogID:       id,
			FilePath:    path,
			Status:      "FAILED",
			Message:     "File not found",
			ErrorDetail: err.Error(),
		}
	}

	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Millisecond * time.Duration(50+rand.Intn(150))) // Simulate processing time

	return LogResult{
		LogID:    id,
		FilePath: path,
		Status:   "OK",
		Message:  "Analysis completed successfully",
	}
}
