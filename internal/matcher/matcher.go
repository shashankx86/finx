package matcher

import (
	"path/filepath"
)

// MatchPattern checks if the given file name matches the specified pattern.
func MatchPattern(name, pattern string) bool {
	match, _ := filepath.Match(pattern, name)
	return match
}
