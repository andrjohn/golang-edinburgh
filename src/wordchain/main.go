package main

import (
	"bufio"
	"fmt"
	"os"
)

func offCount(current, target string) int {
	curr, tart := []byte(current), []byte(target)

	// ASSUME LENGTHS MATCH
	diff := 0
	for i, _ := range curr {
		if curr[i] != tart[i] {
			diff++
		}
	}
	return diff
}

func offByOne(current, target string) bool {
	return (offCount(current, target) == 1)
}

func loadWords() map[int][]string {

	const wordlist = "/etc/dictionaries-common/words"
	words := make(map[int][]string, 0)

	file, err := os.Open(wordlist)
	if err != nil {
		panic("No dictionary")
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		lc := len(word)

		words[lc] = append(words[lc], word)
	}

	return words
}

func nextWords(all []string, match string) []string {
	var cand []string

	for _, v := range all {
		if offByOne(v, match) {
			cand = append(cand, v)
		}
	}
	return cand
}

type wordchain []string

func nextAttempts(lastAttempts []wordchain, words []string) []wordchain {

	var nextAttempts []wordchain

	for _, attempt := range lastAttempts {

		nexts := nextWords(words, attempt[len(attempt)-1])
		for _, word := range nexts {
			newAttempt := wordchain{}
			newAttempt = append(newAttempt, attempt...)
			newAttempt = append(newAttempt, word)
			nextAttempts = append(nextAttempts, newAttempt)
		}
	}
	return nextAttempts
}

func findChain(start, end string) []string {
	return []string{start, end}
}

func main() {

	start := os.Args[1]
	end := os.Args[2]

	words := loadWords()

	atts := []wordchain{[]string{start}}
	for {
		atts = nextAttempts(atts, words[len(start)])
		for _, v := range atts {
			if v[len(v)-1] == end {
				fmt.Printf("Done!\n")
				fmt.Printf("%s\n", v)
				os.Exit(0)
			}
		}
	}

	// nextAttempts([]string{start})
	fmt.Printf("Go from %s to %s!\n", os.Args[1], os.Args[2])
	os.Exit(0)
}
