package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"runtime"
	"strings"
	"sync"
)

func main() {
	fmt.Println(wordCounterConcurrent())
}

func wordCounter() map[string]int {
	b, _ := ioutil.ReadFile("input.txt")
	inputText := string(b)
	mostFrequent := make(map[string]int)
	removeSpecial := regexp.MustCompile(`(?m)[^a-z]`)

	for _, w := range strings.Split(inputText, " ") {
		w = strings.ToLower(w)
		w = removeSpecial.ReplaceAllString(w, "")
		//time.Sleep(200)
		mostFrequent[w] = mostFrequent[w] + 1
	}

	return mostFrequent
}

func wordCounterConcurrent() map[string]int {
	runtime.GOMAXPROCS(4) //Make sure we use all processors
	b, _ := ioutil.ReadFile("input.txt")
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
			//time.Sleep(200)
			wordsChan <- w1
		}(w)
	}

	wg.Wait()
	doneChan <- true
	return mostFrequent
}
