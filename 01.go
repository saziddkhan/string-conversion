package main

import (
	"fmt"
	"regexp"
	"strings"
)

var FirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var AllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func SnakeCase(str string) string {
	snake := FirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = AllCap.ReplaceAllString(snake, "${1}_${2}")

	return strings.ToLower(snake)
}

func main() {
	fmt.Println(SnakeCase("JavaPythonGolang"))
}
