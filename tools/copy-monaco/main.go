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

	// copy src/base to dst/base
	if err = cp.Copy(filepath.Join(src, "base"), filepath.Join(dst, "base")); err != nil {
		slog.Error("failed to copy base", err)
		os.Exit(1)
	} else {
		slog.Info("copied base", "src", filepath.Join(src, "base"), "dst", filepath.Join(dst, "base"))
	}

	// copy src/basic-languages/go to dst/basic-languages/go
	if err = cp.Copy(filepath.Join(src, "basic-languages", "go"), filepath.Join(dst, "basic-languages", "go")); err != nil {
		slog.Error("failed to copy basic-languages/go", err)
		os.Exit(1)
	} else {
		slog.Info("copied basic-languages/go", "src", filepath.Join(src, "basic-languages", "go"), "dst", filepath.Join(dst, "basic-languages", "go"))
	}

	// copy src/editor to dst/editor
	if err = cp.Copy(filepath.Join(src, "editor"), filepath.Join(dst, "editor")); err != nil {
		slog.Error("failed to copy editor", err)
		os.Exit(1)
	} else {
		slog.Info("copied editor", "src", filepath.Join(src, "editor"), "dst", filepath.Join(dst, "editor"))
	}

	// copy src/language/json to dst/language/json
	if err = cp.Copy(filepath.Join(src, "language", "json"), filepath.Join(dst, "language", "json")); err != nil {
		slog.Error("failed to copy language/json", err)
		os.Exit(1)
	} else {
		slog.Info("copied language/json", "src", filepath.Join(src, "language", "json"), "dst", filepath.Join(dst, "language", "json"))
	}

	// copy src/loader.js to dst/loader.js
	if err = cp.Copy(filepath.Join(src, "loader.js"), filepath.Join(dst, "loader.js")); err != nil {
		slog.Error("failed to copy loader.js", err)
		os.Exit(1)
	} else {
		slog.Info("copied loader.js", "src", filepath.Join(src, "loader.js"), "dst", filepath.Join(dst, "loader.js"))
	}
}
