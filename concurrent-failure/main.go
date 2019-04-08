package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"sync"
)

func main() {

	b, _ := ioutil.ReadFile("../input.txt")
	inputText := string(b)
	mostFrequent := make(map[string]int)
	removeSpecial := regexp.MustCompile(`(?m)[^a-z]`)

	wg := sync.WaitGroup{}

	for _, w := range strings.Split(inputText, " ") {
		wg.Add(1)
		go func(w2 string) {
			defer wg.Done()
			w2 = strings.ToLower(w2)
			w2 = removeSpecial.ReplaceAllString(w2, "")
			mostFrequent[w2] = mostFrequent[w2] + 1
		}(w)
	}

	wg.Wait()
	fmt.Println(mostFrequent)
}
