package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateAndWriteFile(path string, fileName string, clusterToWrite string) error {
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
	_, err = file.WriteString(clusterToWrite)
	if err != nil {
		fmt.Println(err)
		return err
	}
	file.Close()
	return nil
}
