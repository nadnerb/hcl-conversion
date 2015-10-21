package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/nadnerb/hcl-conversion/conversion"
)

var version = flag.Bool("version", false, "print version information and exit")

func main() {
	flag.Parse()
	file := flag.Arg(0)

	if *version == true {
		fmt.Printf("%s: %s\n Version: %s\n Commit: %s\n", os.Args[0], ProjectName, Version, gitCommit())
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

	values, err := conversion.LoadConfigValue(path)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	if len(values) > 0 {
		all, _ := conversion.Output(values)
		fmt.Printf("%s\n", all)
		os.Exit(0)
	} else {
		fmt.Print("No Values\n")
		os.Exit(0)
	}
}

