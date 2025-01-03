package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"finx/pkg/finder"
)

func main() {
	// Define flags
	pathFlag := flag.String("path", "", "Directory path to start searching")
	nameFlag := flag.String("name", "", "Name pattern to search for (e.g., *.go)")
	typeFlag := flag.String("type", "", "Type of search: 'f' for files, 'd' for directories")
	maxDepthFlag := flag.Int("maxdepth", -1, "Maximum depth to search (e.g., -maxdepth 2)")
	verboseFlag := flag.Bool("v", false, "Enable verbose output")
	
	// Add new size filter flags
	minSizeFlag := flag.Int64("minsize", 0, "Minimum file size in bytes")
	maxSizeFlag := flag.Int64("maxsize", 0, "Maximum file size in bytes")

	// Add new permissions flag
	permsFlag := flag.String("perms", "", "Filter by file permissions (e.g., 0644, 0755)")

	// Parse flags
	flag.Parse()

	// Handle positional arguments if provided
	var path, pattern string
	if *pathFlag == "" && len(flag.Args()) >= 2 {
		path = flag.Arg(0)
		pattern = strings.Trim(flag.Arg(1), "\"")
	} else {
		path = *pathFlag
		pattern = *nameFlag
	}

	// Validate inputs
	if path == "" || pattern == "" {
		printUsage()
		os.Exit(1)
	}

	// Convert permission string to FileMode
	var perms os.FileMode
	if *permsFlag != "" {
		perm64, err := strconv.ParseUint(*permsFlag, 8, 32)
		if err != nil {
			fmt.Println("Invalid permissions format. Use octal notation (e.g., 0644)")
			os.Exit(1)
		}
		perms = os.FileMode(perm64)
	}

	// Run the finder with the provided options
	options := finder.Options{
		Type:     *typeFlag,
		MaxDepth: *maxDepthFlag,
		Verbose:  *verboseFlag,
		MinSize:  *minSizeFlag,
		MaxSize:  *maxSizeFlag,
		Perms:    perms,
	}
	results := finder.FindFiles(path, pattern, options)

	if len(results) == 0 {
		fmt.Println("No matches found.")
		return
	}

	// Print the results
	for _, result := range results {
		fmt.Println(result)
	}
}

// printUsage provides a clear usage message with examples.
func printUsage() {
    fmt.Println("Finx - An advanced find utility.")
    fmt.Println("\nUsage:")
    fmt.Println("  ./finx <directory_path> <pattern> [flags]")
    fmt.Println("  ./finx -path <directory_path> -name <pattern> [flags]")
    fmt.Println("\nFlags:")
    fmt.Println("  -type       Type of search: 'f' for files, 'd' for directories")
    fmt.Println("  -maxdepth   Maximum depth to search (e.g., -maxdepth 2)")
    fmt.Println("  -v          Enable verbose output")
    fmt.Println("  -minsize    Minimum file size in bytes")
    fmt.Println("  -maxsize    Maximum file size in bytes")
    fmt.Println("  -perms      Filter by file permissions (e.g., 0644, 0755)")
    fmt.Println("  -perms      File permissions in octal format (e.g., -perms 0644)")
    fmt.Println("\nExamples:")
    fmt.Println("  ./finx . \"*.go\"                      # Find all Go files")
    fmt.Println("  ./finx . \"^test.*\\.go$\"            # Find Go files starting with 'test'")
    fmt.Println("  ./finx . \".*_test\\.go$\"            # Find Go test files")
    fmt.Println("  ./finx . \"*.txt\" -type f -maxdepth 2")
    fmt.Println("  ./finx . \"*.log\" -minsize 1024      # Find log files at least 1KB in size")
    fmt.Println("  ./finx . \"*.txt\" -maxsize 10240     # Find text files under 10KB")
    fmt.Println("  ./finx . \"*.sh\" -perms 0755         # Find shell scripts with executable permissions")
    fmt.Println("  ./finx . \"*.sh\" -perms 0755      # Find executable shell scripts")
}
