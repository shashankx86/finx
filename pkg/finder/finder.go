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
	Type       string      // 'f' for files, 'd' for directories
	MaxDepth   int         // Maximum depth to search
	Verbose    bool        // Verbose output
	MinSize    int64       // Minimum file size in bytes
	MaxSize    int64       // Maximum file size in bytes
	Perms      os.FileMode // File permissions to match
}

// FindFiles searches for files and directories based on the given pattern and options.
func FindFiles(root, pattern string, opts Options) []string {
	var results []string

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

		// Size filtering (only for files)
		if !info.IsDir() {
			size := info.Size()

			// Only apply minSize filter if minSize is positive
			if opts.MinSize > 0 && size < opts.MinSize {
				if opts.Verbose {
					fmt.Printf("Skipping %s: size %d < minimum %d\n", path, size, opts.MinSize)
				}
				return nil
			}

			// Only apply maxSize filter if maxSize is positive
			if opts.MaxSize > 0 && size > opts.MaxSize {
				if opts.Verbose {
					fmt.Printf("Skipping %s: size %d > maximum %d\n", path, size, opts.MaxSize)
				}
				return nil
			}
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
