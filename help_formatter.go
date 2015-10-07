package flagparse

import (
	"math"
	"regexp"
)

// HelpFormatter for generating usage messages and argument help strings.
type HelpFormatter struct {
	prog              string
	indentIncrement   int
	maxHelpPosition   int
	width             int
	currentIndent     int
	level             int
	actionMaxLength   int
	rootSection       *section
	currentSection    *section
	whitespaceMatcher *regexp.Regexp
	longBreakMatcher  *regexp.Regexp
}

type section struct{}

// NewHelpFormatter returns new HelpFormatter struct.
func NewHelpFormatter(prog string, indentIncrement, maxHelpPosition, width int) HelpFormatter {
	hf := HelpFormatter{}
	hf.prog = prog
	hf.indentIncrement = indentIncrement
	hf.maxHelpPosition = int(math.Min(float64(maxHelpPosition), math.Max(float64(width-20), float64(indentIncrement*2))))
	hf.width = width
	hf.currentIndent = 0
	hf.level = 0
	hf.actionMaxLength = 0
	hf.rootSection = &section{}
	hf.currentSection = hf.rootSection
	hf.whitespaceMatcher = regexp.MustCompile(`\s+`)
	hf.longBreakMatcher = regexp.MustCompile(`\n\n\n+`)

	return hf
}
