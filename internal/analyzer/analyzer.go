package analyzer

import (
	"errors"
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
	// Correcting os.Stat usage to handle both return values
	if _, statErr := os.Stat(path); statErr != nil {
		if os.IsNotExist(statErr) {
			return LogResult{
				LogID:       id,
				FilePath:    path,
				Status:      "FAILED",
				Message:     "File not found",
				ErrorDetail: statErr.Error(),
			}
		}

		var pathError *os.PathError
		if errors.As(statErr, &pathError) {
			return LogResult{
				LogID:       id,
				FilePath:    path,
				Status:      "FAILED",
				Message:     "Path error detected",
				ErrorDetail: pathError.Error(),
			}
		}

		return LogResult{
			LogID:       id,
			FilePath:    path,
			Status:      "FAILED",
			Message:     "Unknown error",
			ErrorDetail: statErr.Error(),
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
