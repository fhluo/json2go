package main

import (
	"github.com/lmittmann/tint"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	cp "github.com/otiai10/copy"
)

func init() {
	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stderr, &tint.Options{
			Level:      slog.LevelInfo,
			TimeFormat: time.TimeOnly,
		}),
	))
}

func Copy(src string, dst string) {
	err := cp.Copy(src, dst)
	if err != nil {
		slog.Error("failed to copy", "err", err, "src", src, "dst", dst)
		os.Exit(1)
	}

	slog.Info("copied", "src", src, "dst", dst)
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

var Paths = [...]string{
	"base",
	"basic-languages/go",
	"editor",
	"language/json",
	"loader.js",
}

func main() {
	dir := findJSON2GoDir()

	src := filepath.Join(dir, "json2go-web/node_modules/monaco-editor/min/vs")
	dst := filepath.Join(dir, "json2go-web/public/monaco-editor/min/vs")

	_ = os.MkdirAll(dst, 0660)

	for _, path := range Paths {
		Copy(filepath.Join(src, path), filepath.Join(dst, path))
	}
}
