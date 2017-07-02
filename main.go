package main

import (
	"fmt"
	"os"
)

func main() {
	if err := gzs3Cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
