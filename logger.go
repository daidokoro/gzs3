package main

import "fmt"

// Simple logging and printing mechanisms

// logger contains logging flags, colors, debug
type logger struct {
	Colors    *bool
	DebugMode *bool
}

// Info - Prints info level log statments
func (l *logger) Info(msg string) {
	fmt.Printf("%s: %s\n", l.ColorString("info", "green"), msg)
}

// Warn - Prints warn level log statments
func (l *logger) Warn(msg string) {
	fmt.Printf("%s: %s\n", l.ColorString("warn", "yellow"), msg)
}

// Error - Prints error level log statements
func (l *logger) Error(msg string) {
	fmt.Printf("%s: %s\n", l.ColorString("error", "red"), msg)
}

// Debug - Prints debug level log statements
func (l *logger) Debug(msg string) {
	if *l.DebugMode {
		fmt.Printf("%s: %s\n", l.ColorString("debug", "magenta"), msg)
	}
}
