package main

import (
	"fmt"
	"golang.org/x/exp/slog"
	"os"
	"path/filepath"

	cp "github.com/otiai10/copy"
)

func init() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr)))
}

func Copy(src string, dst string) {
	if err := cp.Copy(src, dst); err != nil {
		slog.Error("failed to copy", err, "src", src, "dst", dst)
		os.Exit(1)
	} else {
		slog.Info("copied", "src", src, "dst", dst)
	}
}

func main() {
	wd, err := os.Getwd()
	if err != nil {
		slog.Error("failed to get current directory", err)
		os.Exit(1)
	}

	dir := wd
	for filepath.Base(dir) != "json2go" && filepath.Dir(dir) != dir {
		dir = filepath.Dir(dir)
	}

	if filepath.Dir(dir) == dir {
		slog.Error("", fmt.Errorf("failed to find json2go directory"), "working directory", wd)
		os.Exit(1)
	}

	src := filepath.Join(dir, "web/node_modules/monaco-editor/min/vs")
	dst := filepath.Join(dir, "web/public/monaco-editor/min/vs")

	_ = os.MkdirAll(dst, 0660)

	Copy(filepath.Join(src, "base"), filepath.Join(dst, "base"))
	Copy(filepath.Join(src, "basic-languages", "go"), filepath.Join(dst, "basic-languages", "go"))
	Copy(filepath.Join(src, "editor"), filepath.Join(dst, "editor"))
	Copy(filepath.Join(src, "language", "json"), filepath.Join(dst, "language", "json"))
	Copy(filepath.Join(src, "loader.js"), filepath.Join(dst, "loader.js"))
}
