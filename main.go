package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var version = flag.Bool("version", false, "print version information and exit")

func main() {
	flag.Parse()
	file := flag.Arg(0)

	if *version == true {
		fmt.Printf("%s version %s\n", os.Args[0], Version)
		return
	}

	// not given on the command line? try ENV.
	if file == "" {
		file = os.Getenv("HCL_CONF")
	}

	if file == "" {
		fmt.Printf("Usage: %s [options] path\n", os.Args[0])
		os.Exit(1)
	}

	path, err := filepath.Abs(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid file: %s\n", err)
		os.Exit(1)
	}

	stateFile, err := os.Open(path)
	defer stateFile.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening hcl file: %s\n", err)
		os.Exit(1)
	}

}

