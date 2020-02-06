package examengine

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// SendGet sends GET request, returns the response body
func SendGet(path string) string {
	resp, err := http.Get(path)
	if err != nil {
		return "Error in communicating with path"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Error in parsing body"
	}
	return string(body)
}

// SendPost sends POST request with the JSON object, returns the response body
func SendPost(path string, params string) string {
	var jsonString = []byte(params)
	resp, err := http.Post(path, "application/json", bytes.NewBuffer(jsonString))
	if err != nil {
		return "Error in communicating with path"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Error in parsing body"
	}
	return string(body)
}
