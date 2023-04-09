package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var wordMap = make(map[string]int)

func main() {

	if len(os.Args) > 3 || len(os.Args) < 2 {
		return
	}

	path := os.Args[1]
	num := os.Args[2]

	if len(path) > 0 {
		open(path)
	}

	//need to check whether there are enough args first
	if len(num) > 0 {
		getMostUsed(wordMap, num)
	}
}

func open(path string) {

	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	wordCount := 0

	for scanner.Scan() {

		wordCount++

		wordsToCheck := []string{scanner.Text()}

		mapWords(wordsToCheck, wordMap)
	}

	fmt.Println("Total number of words is ", wordCount, "\n")

}

func mapWords(words []string, wordMap map[string]int) {
	punctRule := regexp.MustCompile(`[\p{P}\p{S}]`)
	for _, word := range words {
		word = strings.ToLower(punctRule.ReplaceAllString(word, ""))
		_, exists := wordMap[word]
		if exists {
			wordMap[word] += 1
		} else {
			wordMap[word] = 1
		}
	}
}

func getMostUsed(m map[string]int, numAsString string) {
	n, err := strconv.Atoi(numAsString)
	if err != nil {
		fmt.Print("err")
	}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})

	sameFreqWords := make(map[int][]string)
	for _, word := range keys {
		freq := m[word]
		sameFreqWords[freq] = append(sameFreqWords[freq], word)
	}

	uniqueFreqs := len(sameFreqWords)
	if n > uniqueFreqs {
		n = uniqueFreqs
	}

	if len(keys) > 0 {
		switch n {
		case 1:
			fmt.Printf("The most used word in this text is '%s' (%d)\n", keys[0], m[keys[0]])
		default:
			fmt.Printf("The most used words in this text are:\n")
			uniqueValues := 0

			for i := 0; uniqueValues < n; i++ {
				freq := m[keys[i]]
				words := sameFreqWords[freq]
				if freq == m[keys[i+1]] {
					continue
				} else {
					uniqueValues++
				}

				fmt.Printf("%s (%d)\n", strings.Join(words, ", "), freq)
				delete(sameFreqWords, freq)
			}
		}
	}
}
