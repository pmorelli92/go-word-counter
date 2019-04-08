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

	doneChan := make(chan bool)
	wordsChan := make(chan string)

	go func() {
		for {
			select {
			case w := <-wordsChan:
				mostFrequent[w] = mostFrequent[w] + 1
			case <-doneChan:
				return
			}
		}
	}()

	wg := sync.WaitGroup{}

	for _, w := range strings.Split(inputText, " ") {
		wg.Add(1)
		go func(w1 string) {
			defer wg.Done()
			w1 = strings.ToLower(w1)
			w1 = removeSpecial.ReplaceAllString(w1, "")
			wordsChan <- w1
		}(w)
	}

	wg.Wait()
	doneChan <- true
	fmt.Println(mostFrequent)
}
