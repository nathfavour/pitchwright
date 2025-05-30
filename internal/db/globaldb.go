package db

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nathfavour/pitchwright/internal/config"
)

const (
	GlobalDBName = "pitch_history.db"
	ProjectDBDir = "projects"
)

// InitGlobalDB initializes the global pitch history database.
func InitGlobalDB() (*sql.DB, error) {
	dir, err := config.EnsureConfigDir()
	if err != nil {
		return nil, err
	}
	dbPath := filepath.Join(dir, GlobalDBName)
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS pitch_history (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		project_path TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		pitch_content TEXT
	)`)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// FlushGlobalDB deletes all pitch history (for --cache flush).
func FlushGlobalDB() error {
	dir, err := config.ConfigDir()
	if err != nil {
		return err
	}
	dbPath := filepath.Join(dir, GlobalDBName)
	return os.Remove(dbPath)
}
