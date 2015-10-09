package flagparse

import (
	"fmt"
	"os"
	"path"
)

// ArgumentParser is a struct for parsing command line strings
type ArgumentParser struct {
	prog string
}

// NewArgumentParser returns new ArgumentParser pointer
func NewArgumentParser(prog string) *ArgumentParser {
	ap := &ArgumentParser{}
	if prog == "" {
		prog = path.Base(os.Args[0])
	}
	ap.prog = prog
	return ap
}

func (ap *ArgumentParser) ParseArgs(args, namespace string) {
	fmt.Printf("usage: %s [-h]\n", ap.prog)
	fmt.Println(`
optional arguments:
  -h, --help  show this help message and exit`)
}
