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
	"math/rand"
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

	Env     string  `json:"env"`
	Source  string  `json:"source"`
	Entry   string  `json:"entry"`
	Inputs  []Input `json:"inputs"`
	Passes  int64   `json:"passes"`
	Timeout int64   `json:"timeout"`
}

type Input struct {
	Type       string   `json:"type"`
	RangeStart int      `json:"rangeStart,omitempty"`
	RangeEnd   int      `json:"rangeEnd,omitempty"`
	RangeList  []string `json:"rangeList,omitempty"`
}

func (i *Input) Generate() (string, error) {
	switch i.Type {
	case "int":
		rand.Seed(time.Now().UnixNano())
		min := i.RangeStart
		max := i.RangeEnd
		return fmt.Sprintf("%v", rand.Intn(max-min+1)+min), nil
	case "string":
		return i.RangeList[rand.Intn(len(i.RangeList))], nil
	}
	return "", fmt.Errorf("cannot generate values for type: '%v'", i.Type)
}

type Test struct {
	Passed       bool     `json:"passed"`
	Inputs       []string `json:"inputs"`
	SourceOutput string   `json:"src_output"`
	TestOutput   string   `json:"test_output"`
}

// TODO: Integrate execConsole (STDIN/STDOUT) execution for the spec

func (s *Spec) execArgv(testPath string) (bool, Test) {
	src := s.Source
	var args []string

	for _, i := range s.Inputs {
		v, err := i.Generate()
		if err != nil {
			panic(err)
		}
		args = append(args, v)
	}
	fmt.Printf("Running spec '%v' with args '%v'\n", src, args)
	env, found := s.Manager.env.Environments[s.Env]
	if !found {
		panic(fmt.Errorf("enviroment '%v' not found for spec '%v'", s.Env, src))
	}

	var timeout int64 = 1000
	if s.Timeout != 0 {
		timeout = s.Timeout
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(timeout))
	defer cancel()

	resultSrc, err := env.Run(ctx, src, args)
	if err != nil {
		panic(err)
	}

	srcOut, err := resultSrc.CombinedOutput()
	if err != nil {
		panic(err)
	}

	resultTest, err := env.Run(ctx, testPath, args)
	if err != nil {
		panic(err)
	}

	resOut, err := resultTest.CombinedOutput()
	if err != nil {
		panic(err)
	}

	return resultSrc == resultTest, Test{
		Passed:       resultSrc == resultTest,
		Inputs:       args,
		SourceOutput: string(srcOut),
		TestOutput:   string(resOut),
	}
}

func (s *Spec) ExecuteTest(testPath string) (bool, []Test) {
	// success := false
	var tests []Test
	var success bool

	switch s.Entry {
	case "argv":
		for i := int64(0); i < s.Passes; i++ {
			s, test := s.execArgv(testPath)
			success = s
			tests = append(tests, test)
		}
	default:
		panic(fmt.Errorf("cannot run test entry through: '%v'", s.Entry))
	}

	return success, tests
}
