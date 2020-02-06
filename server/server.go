package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	examengine "github.com/nokusukun/exam/examengine"
)

type Submission struct {
	Code        string `json:"code"`
	SubmitterID string `json:"submitterID"`
	ActivityID  string `json:"activityID"`
}

func setupResponse(w http.ResponseWriter) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
}

func Hello(w http.ResponseWriter, req *http.Request) {
	setupResponse(w)
	fmt.Fprintf(w, "Sanity Check")
}

func SubmitExam(w http.ResponseWriter, req *http.Request) {
	setupResponse(w)
	decoder := json.NewDecoder(req.Body)

	var submission Submission
	err := decoder.Decode(&submission)
	if err != nil {
		fmt.Println(err, "JSON Decode", req.Body)
		return
	}

	var fileID = submission.SubmitterID + "-" + submission.ActivityID
	var filePath = "tests/" + fileID + ".js"
	var specPath = "specs/" + submission.ActivityID + ".json"

	file, err := os.Create(filePath)

	if err != nil {
		fmt.Println(err, "File Creation")
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	fileSize, err := writer.WriteString(submission.Code)

	if err != nil {
		fmt.Println(err, "File Write")
		return
	}

	fmt.Printf(fileID+" written. (%d) bytes.\n", fileSize)
	writer.Flush()

	output := examengine.RunSubmission(specPath, filePath)
	fmt.Fprintf(w, output)
}

func GetActivities(w http.ResponseWriter, req *http.Request) {
	setupResponse(w)
	file, err := ioutil.ReadFile("./server/activities.json")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(file))
}
