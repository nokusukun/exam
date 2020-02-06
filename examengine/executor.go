package examengine

import (
	"encoding/json"
)

func RunSubmission(specFilePath string, testFilePath string) string {
	manager := InitManager(".")
	spec, err := manager.LoadSpec(specFilePath)
	if err != nil {
		panic(err)
	}
	result := spec.ExecuteTest(testFilePath)
	finalTestOutput, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(finalTestOutput)
}
