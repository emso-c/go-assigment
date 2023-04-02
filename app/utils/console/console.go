// The console package provides some useful functions for styling console output.
package console

// Format represents a console format
type Format string

// Formats
// The following formats are available:
// RED, GREEN, BLUE, RESET, BOLD, UNDERLINE, BLINK
//
// For example:
// fmt.Println(RED + "This is red text" + NC)
// fmt.Println(BOLD + "This is bold text" + NC)
const (
	RED       Format = "\033[0;31m"
	GREEN     Format = "\033[0;32m"
	BLUE      Format = "\033[0;34m"
	RESET     Format = "\033[0m"
	BOLD      Format = "\033[1m"
	UNDERLINE Format = "\033[4m"
	BLINK     Format = "\033[5m"
)
