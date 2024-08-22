package logger

import (
	"fmt"
	"os"
)

const reset = "\033[0m"
const red = "\033[38;2;208;0;0m"
const orange = "\033[38;2;232;93;4m"
const yellow = "\033[38;2;250;163;7m"

func LogLogo() {
	logo := `  __ _
 / _(_)_ _ ___ 
|  _| | '_/ -_)
|_| |_|_| \___|
`
	println(logo)
}

// Debug is used for debug logs and controlled by the flag --debug.
// Uses the same interface as fmt.Printf.
func Debug(s string, a ...any) {
	fmt.Printf(orange+"Debug: "+s+"\n"+reset, a...)
}

// Info is used for general logging.
// Uses the same interface as fmt.Printf.
func Info(s string, a ...any) {
	fmt.Printf(yellow+s+"\n"+reset, a...)
}

// Error is used for error logging.
// Uses the same interface as fmt.Printf.
func Error(s string, a ...any) {
	fmt.Printf(red+"Error: "+s+"\n"+reset, a...)
}

// Fatal is used much in the same way as log.Fatal
// Uses the same interface as fmt.Printf.
func Fatal(s string, a ...any) {
	Error(s, a...)
	os.Exit(1)
}
