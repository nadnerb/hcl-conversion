package main

var ProjectName = "HCL converter"

// The git commit that at compile time. Use -ldflags "-X main.GitCommit `git rev-parse HEAD`"
var GitCommit string
func gitCommit() string {
	if GitCommit != "" {
		return GitCommit
	}
	return "unknown"
}

// The version number.
const Version = "0.0.1"
