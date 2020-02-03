package main

import (
    "fmt"

    exam "github.com/nokusukun/exam"
)


func main() {
    manager := exam.InitManager("../")
    spec, err := manager.LoadSpec("specs/exam_01.json")
    if err != nil {
        panic(err)
    }
    success, result := spec.ExecuteTest("tests/01_test.js")
    fmt.Println(success, result)
}
