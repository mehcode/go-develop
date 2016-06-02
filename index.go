package main

import (
	"net/url"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"
)

var sshShorthandR = regexp.MustCompile(`.*?\@(.*)$`)
var extensionR = regexp.MustCompile(`(\..*)$`)

type gitURL struct {
	Source string
	Owner  string
	Name   string
}

func parse(text string) gitURL {
	u, err := url.Parse(text)
	if err != nil {
		panic(err)
	}

	var result gitURL
	path := u.Path

	// Handle *@* for ssh shorthand
	m := sshShorthandR.FindStringSubmatch(u.Path)
	if len(m) >= 2 {
		parts := strings.SplitN(m[1], ":", 2)
		result.Source = parts[0]
		path = parts[1]
	} else {
		// Assume it's a normal URL and just piecemeal
		result.Source = u.Host
	}

	// Remove `.*` from path
	path = extensionR.ReplaceAllString(path, "")

	// Split path into owner and name
	parts := strings.Split(path, "/")
	result.Owner = parts[len(parts)-2]
	result.Name = parts[len(parts)-1]

	return result
}

func main() {
	// Parse the git URL from the command line
	text := os.Args[1]
	url := parse(text)

	// Get the destination folder (for symlink)
	var linkTarget string
	if len(os.Args) >= 3 {
		linkTarget = os.Args[2]
	} else {
		linkTarget = url.Name
	}

	// Get GOPATH
	gopath := os.Getenv("GOPATH")

	// Build the GOPATH clone target
	cloneTarget := path.Join(gopath, "src", url.Source, url.Owner, url.Name)

	// Clone the repository
	cmd := exec.Command("git", "clone", text, cloneTarget)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	// Link the repository
	err = os.Symlink(cloneTarget, linkTarget)
	if err != nil {
		panic(err)
	}
}
