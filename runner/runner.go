package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	exam "github.com/nokusukun/exam"
)

var specFile string
var testFile string

func init() {
	flag.StringVar(&specFile, "specFile", "", "Specification Path")
	flag.StringVar(&testFile, "testFile", "", "Homework Submission")
	flag.Parse()
	if specFile == "" {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	manager := exam.InitManager("../")
	spec, err := manager.LoadSpec(specFile)
	if err != nil {
		panic(err)
	}
	result := spec.ExecuteTest(testFile)
	rb, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(rb))
}
