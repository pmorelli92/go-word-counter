package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
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

func wordCounterStream() (words map[string]int) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	words = make(map[string]int)

	scanner := bufio.NewScanner(file)

	// TODO we actually need to implement our own split function here since
	// bufio.ScanWords includes symbols like commas
	// split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// 	return
	// }

	// Just use the default for now
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		w := scanner.Text()
		words[strings.ToLower(w)]++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// fmt.Println(words)
	return
}
