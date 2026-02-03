package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var stdin = bufio.NewReader(os.Stdin)

// P prints msg, then reads a line of user input. The input line is then scanned into the args using fmt.Sscan().
//
// This doesn't use fmt.Scanln() because prompt() is often called in a loop (running until user input is valid)
// and Scanln returns two separate errors for example when scanning into one integer and "a\n" is read from stdin,
// resulting in the prompt message being printed twice.
func P(msg string, args ...interface{}) error {
	fmt.Print(msg + " ")

	line, err := stdin.ReadString('\n')
	if err != nil {
		return err
	}

	_, err = fmt.Sscan(line, args...)
	return err
}

// Line prints msg, then reads a line of user input. The input line is then cleaned of whitespace and assigned to val.
func Line(msg string, val *string) error {
	fmt.Print(msg + " ")

	line, err := stdin.ReadString('\n')
	if err != nil {
		return err
	}

	*val = strings.TrimSpace(line)
	if *val == "" {
		return fmt.Errorf("empty input")
	}
	return nil
}

// WithDefault prints msg, then parses a line of user input into
func WithDefault(msg string, arg *string, defaultValue string) error {
	err := Line(msg+" "+fmt.Sprintf("[default %v]", defaultValue), arg)

	if err != nil {
		*arg = defaultValue
		return nil
	}

	return err
}
