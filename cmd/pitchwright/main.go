package main

import (
	"fmt"
	"os"

	"github.com/nathfavour/pitchwright/internal/config"
	"github.com/nathfavour/pitchwright/internal/db"
)

func main() {
	// Ensure config and global DB are initialized
	_, err := config.EnsureConfigDir()
	if err != nil {
		fmt.Println("Error initializing config:", err)
		os.Exit(1)
	}
	_, err = db.InitGlobalDB()
	if err != nil {
		fmt.Println("Error initializing global DB:", err)
		os.Exit(1)
	}
	// ...existing code for CLI flags, commands, and project DB init...
	fmt.Println("Pitchwright CLI initialized. (Add CLI logic here)")
}
