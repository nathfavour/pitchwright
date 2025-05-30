package config

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	ConfigDirName  = ".pitchwright"
	ConfigFileName = "configs.json"
)

// ConfigDir returns the path to the .pitchwright directory in the user's home.
func ConfigDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ConfigDirName), nil
}

// EnsureConfigDir ensures the .pitchwright directory and configs.json exist.
func EnsureConfigDir() (string, error) {
	dir, err := ConfigDir()
	if err != nil {
		return "", err
	}
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return "", err
		}
	}
	configPath := filepath.Join(dir, ConfigFileName)
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		defaultConfig := map[string]interface{}{
			"keyFolders": []string{"src", "app", "lib"},
		}
		f, err := os.Create(configPath)
		if err != nil {
			return "", err
		}
		defer f.Close()
		json.NewEncoder(f).Encode(defaultConfig)
	}
	return dir, nil
}

// HashString returns a short SHA1 hash of a string (for DB naming).
func HashString(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))[:12]
}
