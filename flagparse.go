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
	var argList []string
	if args == "" {
		argList = os.Args[1:]
	}
	fmt.Printf("usage: %s [-h]\n", ap.prog)

	if 0 < len(argList) && (argList[0] == "-h" || argList[0] == "--help") {
		fmt.Println(`
optional arguments:
  -h, --help  show this help message and exit`)
	} else {
		fmt.Printf("%s: error: unrecognized arguments %s\n", ap.prog, argList[0])
	}
}
