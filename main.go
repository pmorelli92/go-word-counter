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
	"unicode"
	"unicode/utf8"
)

func main() {
	fmt.Println(wordCounterStream())
}

func wordCounter() (words map[string]int) {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	removeLineBreaks := regexp.MustCompile(`\r?\n`)
	inputText := removeLineBreaks.ReplaceAllString(string(b)," ")
	words = make(map[string]int)
	removeSpecial := regexp.MustCompile(`(?m)[^a-z]`)

	for _, w := range strings.Split(inputText, " ") {
		w = strings.ToLower(w)
		w = removeSpecial.ReplaceAllString(w, "")
		if w != "" {
			words[w]++
		}
	}

	return
}

func wordCounterConcurrent() (words map[string]int) {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	removeLineBreaks := regexp.MustCompile(`\r?\n`)
	inputText := removeLineBreaks.ReplaceAllString(string(b)," ")
	words = make(map[string]int)
	removeSpecial := regexp.MustCompile(`(?m)[^a-z]`)

	doneChan := make(chan bool)
	wordsChan := make(chan string)

	go func() {
		for {
			select {
			case w := <-wordsChan:
				words[w]++
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
			if w1 != "" {
				wordsChan <- w1
			}
		}(w)
	}

	wg.Wait()
	doneChan <- true
	return
}

func wordCounterStream() (words map[string]int) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	words = make(map[string]int)
	scanner := bufio.NewScanner(file)

	// bufio.ScanWords includes punctuation that we want to remove
	// reimplemented that method checking with unicode.IsPunct
	scanner.Split(ScanWords)

	for scanner.Scan() {
		w := scanner.Text()
		if w != "" {
			words[strings.ToLower(w)]++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Skip leading spaces.
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if !unicode.IsSpace(r) || !unicode.IsPunct(r) {
			break
		}
	}
	// Scan until space, marking end of word.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if unicode.IsSpace(r) || unicode.IsPunct(r) {
			return i + width, data[start:i], nil
		}
	}
	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return start, nil, nil
}
