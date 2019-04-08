package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {

	b, _ := ioutil.ReadFile("../input.txt")
	inputText := string(b)
	mostFrequent := make(map[string]int)
	removeSpecial := regexp.MustCompile(`(?m)[^a-z]`)

	for _, w := range strings.Split(inputText, " ") {
		w = strings.ToLower(w)
		w = removeSpecial.ReplaceAllString(w, "")
		mostFrequent[w] = mostFrequent[w] + 1
	}

	fmt.Println(mostFrequent)
}
