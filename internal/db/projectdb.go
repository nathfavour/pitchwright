package db

import (
	"database/sql"
	"encoding/json"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nathfavour/pitchwright/internal/config"
)

// ProjectDBPath returns the path to the per-project DB based on project path hash.
func ProjectDBPath(projectPath string) (string, error) {
	dir, err := config.EnsureConfigDir()
	if err != nil {
		return "", err
	}
	projectsDir := filepath.Join(dir, "projects")
	if _, err := os.Stat(projectsDir); os.IsNotExist(err) {
		if err := os.MkdirAll(projectsDir, 0755); err != nil {
			return "", err
		}
	}
	// Use a hash of the project path for unique DB name
	hash := config.HashString(projectPath)
	return filepath.Join(projectsDir, hash+".db"), nil
}

// InitProjectDB initializes the per-project metadata database.
func InitProjectDB(projectPath string, metadata map[string]interface{}) (*sql.DB, error) {
	dbPath, err := ProjectDBPath(projectPath)
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS project_metadata (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		project_name TEXT,
		project_path TEXT,
		repo_url TEXT,
		key_files TEXT,
		last_updated DATETIME DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		return nil, err
	}
	// Insert or update metadata
	keyFiles, _ := json.Marshal(metadata["key_files"])
	_, err = db.Exec(`INSERT OR REPLACE INTO project_metadata (id, project_name, project_path, repo_url, key_files) VALUES (1, ?, ?, ?, ?)`,
		metadata["project_name"], metadata["project_path"], metadata["repo_url"], string(keyFiles))
	if err != nil {
		return nil, err
	}
	return db, nil
}

// LoadProjectMetadata loads metadata from the per-project DB.
func LoadProjectMetadata(projectPath string) (map[string]interface{}, error) {
	dbPath, err := ProjectDBPath(projectPath)
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	row := db.QueryRow(`SELECT project_name, project_path, repo_url, key_files FROM project_metadata WHERE id=1`)
	var name, path, repo, keyFiles string
	err = row.Scan(&name, &path, &repo, &keyFiles)
	if err != nil {
		return nil, err
	}
	var files []string
	json.Unmarshal([]byte(keyFiles), &files)
	return map[string]interface{}{
		"project_name": name,
		"project_path": path,
		"repo_url":     repo,
		"key_files":    files,
	}, nil
}
