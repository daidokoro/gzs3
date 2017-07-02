package main

// HandleError - exits on error
import "os"

func handleError(err error) {
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}
