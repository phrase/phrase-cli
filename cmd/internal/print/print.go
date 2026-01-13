package print

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/phrase/phrase-cli/cmd/internal/shared"

	ct "github.com/daviddengcn/go-colortext"
)

const phrase_logo = `
   *****@@@@@@@@@@@@@@@@@@@@.
   .********@@@@@@@@@@@@@@@@@@@@
   .***********             @@@@@
   .**************          ,@@@@
   .**************          ,@@@@
   .**************          ,@@@@
   .**************          ,@@@@
   .**************          ,@@@@
   .**************          ,@@@@
   .**************          ,@@@@
   .**************          @@@@@
   .**************  (@@@@@@@@@@@
   .**************  (@@@@@@@@@.
   **************
      ***********
         ********
            *****
`

type ErrorMessage struct {
	Error string `json:"error"`
}

type InfoMessage struct {
	Message string `json:"message"`
}

func PhraseLogo() {
	WithColor(ct.Cyan, phrase_logo)
}

// Prints the argument if the command is started in non-batch mode
func NonBatchPrintf(msg string, args ...interface{}) {
	if !shared.BatchMode {
		fmt.Printf(msg, args...)
	}
}

func NonBatchPrint(msg string) {
	if !shared.BatchMode {
		fmt.Print(msg)
	}
}

func Success(msg string, args ...interface{}) {
	if shared.BatchMode {
		jsonStruct := &InfoMessage{Message: fmt.Sprintf(msg, args...)}
		encodedJson, _ := json.Marshal(jsonStruct)
		fmt.Println(string(encodedJson))
	} else {
		WithColor(ct.Green, msg, args...)
	}
}

func Warn(msg string, args ...interface{}) {
	if shared.BatchMode {
		jsonStruct := &InfoMessage{Message: fmt.Sprintf(msg, args...)}
		encodedJson, _ := json.Marshal(jsonStruct)
		fmt.Println(string(encodedJson))
	} else {
		WithColor(ct.Yellow, msg, args...)
	}
}

func Failure(msg string, args ...interface{}) {
	WithColor(ct.Red, msg, args...)
}

func WithColor(color ct.Color, msg string, args ...interface{}) {
	fprintWithColor(os.Stdout, color, msg, args...)
}

func Error(err error) {
	if shared.BatchMode {
		jsonStuct := &ErrorMessage{Error: fmt.Sprint(err)}
		encodedJson, _ := json.Marshal(jsonStuct)
		fmt.Println(string(encodedJson))
	} else {
		fprintWithColor(os.Stderr, ct.Red, "ERROR: %s", err)
	}
}

func fprintWithColor(w io.Writer, color ct.Color, msg string, args ...interface{}) {
	ct.Foreground(color, true)
	fmt.Fprintf(w, msg, args...)
	fmt.Fprintln(w)
	ct.ResetColor()
}
