package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Config struct {
	ScanDirectories  []string `json:"scanDirectories"`
	ThresholdDays    int      `json:"thresholdDays"`
	LogDirectory     string   `json:"logDirectory"`
	LogThresholdDays int      `json:"logThresholdDays"`
}

var config Config

func LoadConfigFile(confiFile string) error {
	file, err := ioutil.ReadFile(confiFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

func DeleteOldEnvironments() {
	now := time.Now()

	for _, dir := range config.ScanDirectories {
		filepath.Walk(
			dir, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					log.Println(err)
					return nil
				}

				if info.IsDir() && (filepath.Base(path) == "venv" || filepath.Base(path) == ".venv") {
					days := now.Sub(info.ModTime()).Hours() / 24
					if int(days) > config.ThresholdDays {
						log.Printf("deleteing: %s - %d days\n", path, int(days))
						err := os.RemoveAll(path)
						if err != nil {
							log.Printf("Failed to delete: %s: %v\n", path, err)
						} else {
							log.Printf("Deleted: %s\n", path)
						}
					}
				}
				return nil
			})
	}
}
