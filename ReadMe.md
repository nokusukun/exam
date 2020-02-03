# Kingsland Automated Exam Checker

## Dependencies

* NodeJS v12 or Higher
* Golang v1.12 or Higher

## Running

In the [runner folder]("./runner"):

1. Add the base case of the exam in which all tests will be run against on the [exams folder]("./runner/exams").
2. Add student submissions on the [tests folder]("./runner/tests").
3. Add parameter specifications on the [specs folder]("./runner/specs").
4. Run: `go run runner.go`