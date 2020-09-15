package cmd

import (
	"log"
	"os"
	"path/filepath"
)

func getConfPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(homeDir, ".pathtab")
}

func getConfFile() *os.File {
	filePath := getConfPath()
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	return file
}
