package matcher

import (
    "path/filepath"
    "regexp"
    "strings"
)

// MatchPattern checks if the given file name matches the specified pattern using either glob or regex
func MatchPattern(name, pattern string) bool {
    // Try regex first
    if isRegexPattern(pattern) {
        r, err := regexp.Compile(pattern)
        if err == nil {
            return r.MatchString(name)
        }
    }
    // Fallback to glob pattern
    match, _ := filepath.Match(pattern, name)
    return match
}

// isRegexPattern checks if the pattern looks like a regex
func isRegexPattern(pattern string) bool {
    regexChars := []string{"^", "$", "+", "?", "{", "}", "[", "]", "(", ")", "|", "\\"}
    for _, char := range regexChars {
        if strings.Contains(pattern, char) {
            return true
        }
    }
    return false
}
