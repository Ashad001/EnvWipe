package main

import (
	"os"
	"testing"
	"time"
)

func createFakeEnvironments(t *testing.T, basedir, name string, modTime time.Time) string {
	envDir := basedir + "/" + name
	err := os.Mkdir(envDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create directory: %v\n", err)
	}

	err = os.Chtimes(envDir, modTime, modTime)

	if err != nil {
		t.Fatalf("Failed to change time: %v\n", err)
	}

	return envDir
}

func TestDeleteOldEnvironments(t *testing.T) {

	tmpDir, err := os.MkdirTemp("", "env_test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v\n", err)
	}

	defer os.RemoveAll(tmpDir)

	oldTime := time.Now().AddDate(0, 0, -40) // 40 days ago
	newTime := time.Now().AddDate(0, 0, -10) // 10 days ago

	envs := []struct {
		name     string
		modTime  time.Time
		expected bool // True if the environment SHOULD be deleted
	}{
		{"venv", oldTime, true},
		{".venv", oldTime, true},
		{"temp_env", newTime, false},
	}

	for _, env := range envs {
		createFakeEnvironments(t, tmpDir, env.name, env.modTime)
	}

	config = Config{
		ScanDirectories: []string{tmpDir},
		ThresholdDays:   30,
		ExcludeDirectories: []string{
			"temp_env",
		},
	}

	DeleteOldEnvironments()

	for _, env := range envs {
		_, err := os.Stat(tmpDir + "/" + env.name)

		if env.expected && err == nil {
			t.Fatalf("Expected %s to be deleted\n", env.name)
		}

		if !env.expected && err != nil {
			t.Errorf("Expected %s to be kept\n", env.name)
		}
	}

}
