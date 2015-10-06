package flagparse

// HelpFormatter for generating usage messages and argument help strings.
type HelpFormatter struct {
	prog              string
	indentIncrement   int
	maxHelpPosition   int
	width             string
	currentIndent     int
	level             int
	actionMaxLength   int
	rootSection       string
	currentSection    string
	whitespaceMatcher string
	longBreakMatcher  string
}

// NewHelpFormatter returns new HelpFormatter struct.
func NewHelpFormatter(prog string, indentIncrement, maxHelpPosition, width int) HelpFormatter {
	hf := HelpFormatter{}
	hf.prog = prog
	hf.indentIncrement = indentIncrement
	//hf.maxHelpPosition = math.Min(maxHelpPosition, math.Max(width-20, indentIncrement*2))

	return hf
}
