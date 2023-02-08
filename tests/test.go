package main

import (
	"log"
	"strings"
)

func main() {
	log.Println(Split("a/b/c", "/"))
}
func Split(s, sep string) []string {
	var result []string
	i := strings.Index(s, sep)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}
	return append(result, s)
}
