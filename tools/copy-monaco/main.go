package main

import (
	"log"
	"log/slog"
	"os"
	"path/filepath"

	cp "github.com/otiai10/copy"
)

func init() {
	log.SetFlags(0)
}

func Copy(src string, dst string) {
	if err := cp.Copy(src, dst); err != nil {
		slog.Error("failed to copy", "err", err, "src", src, "dst", dst)
		os.Exit(1)
	} else {
		slog.Info("copied", "src", src, "dst", dst)
	}
}

// findJSON2GoDir finds the json2go directory.
func findJSON2GoDir() string {
	// Get working directory
	wd, err := os.Getwd()
	if err != nil {
		slog.Error("failed to get current directory", "err", err)
		os.Exit(1)
	}

	// Find json2go directory
	dir := wd
	for filepath.Base(dir) != "json2go" && filepath.Dir(dir) != dir {
		dir = filepath.Dir(dir)
	}

	if filepath.Dir(dir) == dir {
		slog.Error("failed to find json2go directory", "working directory", wd)
		os.Exit(1)
	}

	return dir
}

func main() {
	dir := findJSON2GoDir()

	src := filepath.Join(dir, "app/ui/node_modules/monaco-editor/min/vs")
	dst := filepath.Join(dir, "app/ui/public/monaco-editor/min/vs")

	_ = os.MkdirAll(dst, 0660)

	Copy(filepath.Join(src, "base"), filepath.Join(dst, "base"))
	Copy(filepath.Join(src, "basic-languages", "go"), filepath.Join(dst, "basic-languages", "go"))
	Copy(filepath.Join(src, "editor"), filepath.Join(dst, "editor"))
	Copy(filepath.Join(src, "language", "json"), filepath.Join(dst, "language", "json"))
	Copy(filepath.Join(src, "loader.js"), filepath.Join(dst, "loader.js"))
}
