package clier

import (
	"fmt"
	"os"
	"path"
	"strings"
)

type Namespace struct{}
type actionsContainer struct{}
type attributeHolder struct{}

// def add_argument(self, *args, **kwargs):
// args = 任意のタプルになる
// kwargs = 任意個のディクショナリ
func (ac *actionsContainer) AddArgument() {
	fmt.Println("hoghoge")
}

// ArgumentParser is a struct for parsing command line strings
type ArgumentParser struct {
	prog string
	*actionsContainer
}

type ArgumentParserCommand struct {
	prog                string
	usage               string
	description         string
	epilog              string
	parents             []interface{} //ArgumentParser TODO: interface
	formatter           interface{}
	prefixChars         string
	fromfilePrefixChars interface{}
	argumentDefault     interface{}
	conflictHandler     string
	addHelp             bool
}

// NewArgumentParser returns new ArgumentParser pointer
func NewArgumentParser(apc *ArgumentFactory) *ArgumentParser {
	ap := &ArgumentParser{}
	if prog == "" {
		prog = path.Base(os.Args[0])
	}
	ap.prog = prog
	return ap
}

func (ap *ArgumentParser) ParseArgs(args []string, ns *Namespace) *Namespace {
	namespace, argList := ap.ParseKnownArgs(args, ns)
	if 0 < len(argList) {
		msg := fmt.Sprintf("unrecognized arguments: %s", strings.Join(argList, " "))
		ap.Error(msg)
	}
	return namespace
}

func (ap *ArgumentParser) ParseKnownArgs(args []string, namespace *Namespace) (*Namespace, []string) {
	namespace, args = ap.parseKnownArgs(args, namespace)
	return namespace, args
}

func (ap *ArgumentParser) parseKnownArgs(args []string, namespace *Namespace) (*Namespace, []string) {
	return namespace, args
}

/*
** ======================
** Help-printing methods
** ======================
 */
func (ap *ArgumentParser) PrintUsage(file *os.File) {
	//ap.printMessage(ap.FormatUsage(), file)
}

func (ap *ArgumentParser) PrintHelp(file *os.File) {
	//ap.printMessage(ap.FormatHelp(), file)
}

func (ap *ArgumentParser) printMessage(message string, file *os.File) {
	fmt.Fprint(file, message)
}

/*
** =================
** Exiting methods
** =================
 */

func (ap *ArgumentParser) Exit(status int, message string) {
	ap.printMessage(message, os.Stderr)
	os.Exit(status)
}

/* Error prints a usage message incorporating the message to stderr and exits.
** If you override this in a struct, it should not return -- it should either exit or raise an exception.
 */
func (ap *ArgumentParser) Error(message string) {
	ap.PrintUsage(os.Stderr)
	ap.Exit(2, fmt.Sprintf("%s: error: %s\n", ap.prog, message))
}
