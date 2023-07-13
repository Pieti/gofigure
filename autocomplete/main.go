package main

import (
	"bufio"
	"flag"
	"fmt"
	"go-autocomplete/trie"
	"log"
	"os"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("Usage: go-autocomplete <prefix>")
		os.Exit(1)
	}
	prefix := flag.Arg(0)

	filepath := "vocabulary.txt"
	vocabulary := get_words(filepath)

	root := trie.NewTrie()
	for _, word := range vocabulary {
		root.Add(word)
	}
	n := root.FindNode(prefix)
	count := 0
	if n != nil {
		count = n.Count()
	}
	fmt.Printf("Found %d words with prefix '%s'\n\n", count, prefix)
	words := root.FindWordsWithPrefix(prefix)
	for i, word := range words {
		fmt.Println(word)
		if i >= 10 {
			fmt.Println("...")
			break
		}
	}

}

func get_words(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	words := make([]string, 0)
	for scanner.Scan() {
		word := scanner.Text()
		if len(word) < 3 {
			continue
		}
		words = append(words, scanner.Text())
	}
	return words
}
