package main

// gzs3file constant
const config = "gzs3file"

var (
	gituser string
	gitpass string
	gituri  string
	gitrsa  string
	profile string
	region  string
	debug   bool
	colors  bool

	// define logger
	log = Logger{
		DebugMode: &debug,
		Colors:    &colors,
	}
)
