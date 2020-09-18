package cmd

import (
	"io/ioutil"
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

func writeConf(text string) {
	filePath := getConfPath()
	data := []byte(text)
	ioutil.WriteFile(filePath, data, 0644)
}

func readConf() string {
	filePath := getConfPath()
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		data = []byte("")
	}
	return string(data)
}
