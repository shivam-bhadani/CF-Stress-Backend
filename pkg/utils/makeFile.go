package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func MakeFile(path string, fileName string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	pathToCreateFile := filepath.Join(wd, path)
	fileToCreate := filepath.Join(pathToCreateFile, fileName)
	file, err := os.Create(fileToCreate)
	if err != nil {
		fmt.Println(err)
		return err
	}
	file.Close()
	return nil
}
