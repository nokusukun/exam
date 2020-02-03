package main

import (
	"encoding/json"
	"fmt"

	exam "github.com/nokusukun/exam"
)

func main() {
	manager := exam.InitManager("../")
	spec, err := manager.LoadSpec("specs/exam_01.json")
	if err != nil {
		panic(err)
	}
	result := spec.ExecuteTest("tests/01_test.js")
	rb, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(rb))
}
