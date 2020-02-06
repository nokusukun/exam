package examengine

import "fmt"

// IsDebugEnabled - Debug flag
var IsDebugEnabled = false

// Log - General logging function
func Log(a ...interface{}) {
	if IsDebugEnabled {
		fmt.Println(a)
	}
}
