package main

import (
	"fmt"
	"net/http"

	server "github.com/nokusukun/exam/server"
)

func main() {
	http.HandleFunc("/hello", server.Hello)
	http.HandleFunc("/exam/submit", server.SubmitExam)
	http.HandleFunc("/activities", server.GetActivities)

	fmt.Println("Serving on port 8888")
	http.ListenAndServe(":8888", nil)
}
