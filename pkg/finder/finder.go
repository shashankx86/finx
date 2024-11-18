package finder

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"finx/internal/matcher"
)

// Options for searching
type Options struct {
	Type     string // 'f' for files, 'd' for directories
	MaxDepth int    // Maximum depth to search
	Verbose  bool   // Verbose output
}

// FindFiles searches for files and directories based on the given pattern and options.
func FindFiles(root, pattern string, opts Options) []string {
	var results []string
	// currentDepth := 0

	// Walk function with depth control
	var walkFn filepath.WalkFunc = func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // Skip files with errors
		}

		// Calculate depth
		depth := len(strings.Split(filepath.ToSlash(path), "/")) - len(strings.Split(filepath.ToSlash(root), "/"))
		if opts.MaxDepth >= 0 && depth > opts.MaxDepth {
			return filepath.SkipDir
		}

		// Verbose output
		if opts.Verbose {
			fmt.Println("Checking:", path)
		}

		// Type filtering
		if opts.Type == "f" && info.IsDir() {
			return nil
		}
		if opts.Type == "d" && !info.IsDir() {
			return nil
		}

		// Match pattern
		if matcher.MatchPattern(info.Name(), pattern) {
			results = append(results, path)
		}
		return nil
	}

	// Start walking the directory tree
	filepath.Walk(root, walkFn)
	return results
}
