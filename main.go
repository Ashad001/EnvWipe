package main

import (
	"encoding/json"
	"fmt"
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

// function with optional parameters
func CleanUpOldLogs() {
	now := time.Now()

	filepath.Walk(
		config.LogDirectory, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Println(err)
				return nil
			}

			days := now.Sub(info.ModTime()).Hours() / 24
			if int(days) > config.LogThresholdDays {
				log.Printf("deleteing: %s - %d days\n", path, int(days))
				err := os.Remove(path)
				if err != nil {
					log.Printf("Failed to delete: %s: %v\n", path, err)
				} else {
					log.Printf("Deleted: %s\n", path)
				}
			}
			return nil
		})
	
}

func main() {
	e := LoadConfigFile("config.json")
	if e != nil {
		log.Fatalf("Failed to load config file: %v\n", e)
	}

	if _, err := os.Stat(config.LogDirectory); os.IsNotExist(err) {
		// Log error and make a  directory
		log.Println("Log directory does not exist. Creating it.")
		os.Mkdir(config.LogDirectory, 0755) // 0755 is the permission for the dieectory to read, write and execute
	}

	logFile, err := os.OpenFile(
		filepath.Join(config.LogDirectory, "cleanup.log"),
		os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v\n", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	
	log.Println("Starting cleanup process")
	
	DeleteOldEnvironments()
	CleanUpOldLogs()

	log.Println("Cleanup process completed successfully")
	fmt.Println("CLeanup process completed successfully, check the log file for more details")
	


}
