package main

import (
	"mjpclab.dev/ehfs/src"
	"os"
)

func main() {
	ok := src.Main()
	if !ok {
		os.Exit(1)
	}
}
