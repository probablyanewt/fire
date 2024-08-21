package logger

import (
	"fmt"
)

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
	fmt.Printf("Debug: "+s+"\n", a...)
}

// Info is used for general logging.
// Uses the same interface as fmt.Printf.
func Info(s string, a ...any) {
	fmt.Printf(s+"\n", a...)
}

// Error is used for error logging.
// Uses the same interface as fmt.Printf.
func Error(s string, a ...any) {
	fmt.Printf("Error: "+s+"\n", a...)
}
