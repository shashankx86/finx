package main

import (
	"flag"
	"fmt"
	"os"
	// "path/filepath"
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

	// Run the finder with the provided options
	options := finder.Options{
		Type:     *typeFlag,
		MaxDepth: *maxDepthFlag,
		Verbose:  *verboseFlag,
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
	fmt.Println("\nExamples:")
	fmt.Println("  ./finx /path/to/search \"*.go\"")
	fmt.Println("  ./finx /path/to/search *.txt -type f -maxdepth 2 -v")
	fmt.Println("  ./finx -path /path/to/search -name \"*.md\" -type d")
}
