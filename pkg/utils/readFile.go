package utils

import (
	"os"
)

func ReadFile(pathOfFileToRead string) (string, error) {
	data, err := os.ReadFile(pathOfFileToRead)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
