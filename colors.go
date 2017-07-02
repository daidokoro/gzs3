package main

import (
	"runtime"
	"strings"

	"github.com/fatih/color"
)

// ColorString - Returns colored string
func (l *Logger) ColorString(s, col string) string {

	// If Windows, disable colorS
	if runtime.GOOS == "windows" || *l.Colors {
		return s
	}

	var result string
	switch strings.ToLower(col) {
	case "green":
		result = color.New(color.FgGreen).Add(color.Bold).SprintFunc()(s)
	case "yellow":
		result = color.New(color.FgYellow).Add(color.Bold).SprintFunc()(s)
	case "red":
		result = color.New(color.FgRed).Add(color.Bold).SprintFunc()(s)
	case "magenta":
		result = color.New(color.FgMagenta).Add(color.Bold).SprintFunc()(s)
	case "cyan":
		result = color.New(color.FgCyan).Add(color.Bold).SprintFunc()(s)
	default:
		// Unidentified, just returns the same string
		return s
	}

	return result
}
