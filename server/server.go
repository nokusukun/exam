package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	examengine "github.com/nokusukun/exam/examengine"
)

type Submission struct {
	Code        string `json:"code"`
	SubmitterID string `json:"submitterID"`
	ActivityID  string `json:"activityID"`
}

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Sanity Check")
}

func SubmitExam(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var submission Submission
	err := decoder.Decode(&submission)
	if err != nil {
		panic(err)
	}

	var fileID = submission.SubmitterID + "-" + submission.ActivityID
	var filePath = "tests/" + fileID + ".js"
	var specPath = "specs/" + submission.ActivityID + ".json"

	file, err := os.Create(filePath)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	fileSize, err := writer.WriteString(submission.Code)

	if err != nil {
		panic(err)
	}

	fmt.Printf(fileID+" written. (%d) bytes.\n", fileSize)
	writer.Flush()

	output := examengine.RunSubmission(specPath, filePath)
	fmt.Fprintf(w, output)
}
