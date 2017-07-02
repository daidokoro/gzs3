package main

// gzs3file constant
const gzs3file = "gzs3file"

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
	log = logger{
		DebugMode: &debug,
		Colors:    &colors,
	}
)
