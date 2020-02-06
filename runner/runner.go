package runner

import (
	"flag"
	"fmt"
	"os"

	examengine "github.com/nokusukun/exam/examengine"
)

var specFile string
var testFile string
var debugEnabled bool

func init() {
	flag.StringVar(&specFile, "specFile", "", "Specification Detail Path")
	flag.StringVar(&testFile, "testFile", "", "Homework Submission")
	flag.BoolVar(&debugEnabled, "debug", false, "Enable debugging logs")
	flag.Parse()

	examengine.IsDebugEnabled = debugEnabled

	if debugEnabled {
		examengine.Log("Debug enabled")
	}

	if specFile == "" || testFile == "" {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	output := examengine.RunSubmission(specFile, testFile)
	fmt.Println(output)
}
