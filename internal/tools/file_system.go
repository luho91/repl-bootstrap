package tools

import (
	"os"
)

func ReadFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	return string(data), err
}
