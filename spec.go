// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    spec, err := UnmarshalSpec(bytes)
//    bytes, err = spec.Marshal()

package exam

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

func UnmarshalSpec(data []byte) (Spec, error) {
	var r Spec
	err := json.Unmarshal(data, &r)
	return r, err
}

func (s *Spec) Marshal() ([]byte, error) {
	return json.Marshal(s)
}

type Spec struct {
	Manager *Manager
	// Environment to use
	Env string `json:"env"`

	// HTTP Only: Source 'master' code to run the values
	Source string `json:"source"`

	// Entry type defines what sort of application it is
	// argv - Values gets passed through the argument variables
	// console - Gets passed through STDIN
	// http - Values are passed through as rest APIs
	Entry string `json:"entry"`

	// Inputs determine the values to be passed on the source/test codes
	// They can be static values or generated randomly
	Data []Input `json:"data"`

	// Determines how many passes to run the program
	Passes int64 `json:"passes"`

	// Timeout in milliseconds, stops the code from running for more than x number of times.
	Timeout int64 `json:"timeout"`

	// HTTP Only: Endpoints to test
	HTTPEndpoints []HTTPEndpoint `json:"endpoints"`
}

type HTTPEndpoint struct {
	Endpoint    string      `json:"endpoint"`
	Body        interface{} `json:"body"`
	Method      string      `json:"method"`
	RequestBody string      `json:"requestBody"`
	Expected    string      `json:"expected"`
}

type Input struct {
	Arguments []string `json:"arguments"`
	Expected  string   `json:"expected"`
}

type Test struct {
	Passed       bool     `json:"passed"`
	Inputs       []string `json:"inputs"`
	SourceOutput string   `json:"src_output"`
	TestOutput   string   `json:"test_output"`
}

func (s *Spec) execArgv(testPath string) []Test {
	env, found := s.Manager.env.Environments[s.Env]
	if !found {
		panic(fmt.Errorf("enviroment '%v' not found for spec '%v'", s.Env))
	}

	var timeout int64 = 1000
	if s.Timeout != 0 {
		timeout = s.Timeout
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(timeout))
	defer cancel()

	var tests []Test

	for _, i := range s.Data {
		var test Test
		var args = i.Arguments

		resultTest, err := env.Run(ctx, testPath, args)
		if err != nil {
			panic(err)
		}

		resOut, err := resultTest.CombinedOutput()
		if err != nil {
			panic(err)
		}

		var normalizedTestOutput = strings.TrimSpace(string(resOut))

		test.Passed = string(i.Expected) == normalizedTestOutput
		test.Inputs = args
		test.SourceOutput = string(i.Expected)
		test.TestOutput = normalizedTestOutput

		tests = append(tests, test)
	}

	return tests
}

func (s *Spec) execHTTP() []Test {
	var tests []Test
	for _, i := range s.HTTPEndpoints {
		var tmpTest Test
		var actual string
		var url = "http://" + s.Source + i.Endpoint
		if i.Method == "GET" {
			actual = SendGet(url)
		} else if i.Method == "POST" {
			actual = SendPost(url, i.RequestBody)
		} else {
			panic("Unknown HTTP Method: " + i.Method)
		}
		tmpTest.Passed = actual == i.Expected
		tmpTest.Inputs = []string{url, i.RequestBody}
		tmpTest.SourceOutput = i.Expected
		tmpTest.TestOutput = actual
		tests = append(tests, tmpTest)
	}

	return tests
}

// ExecuteTest runs the requested tests depending on test type (argv || http)
func (s *Spec) ExecuteTest(testPath string) []Test {
	var tests []Test

	switch s.Entry {
	case "argv":
		if testPath == "" {
			panic("Test path is empty")
		}
		tests = s.execArgv(testPath)
	case "http":
		tests = s.execHTTP()
	default:
		panic(fmt.Errorf("cannot run test entry through: '%v'", s.Entry))
	}

	return tests
}
