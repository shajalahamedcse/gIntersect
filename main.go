package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var howToUse = `usage : intersect [File A] [File B]
Prints out the lines that are both in [File A] and [File B]`

//Exit with error message
func exit(msg interface{}) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

//Get the Word Map
func getWordMap(fileName string) (map[string]bool, error) {

	//This map will load every word as key and true as value
	var wordMap = make(map[string]bool)
	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// word by word parsing
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		wordMap[scanner.Text()] = true
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return wordMap, err
}

func intersect(wordMap map[string]bool, inputFile string) error {
	file, err := os.Open(inputFile)
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		if wordMap[word] {
			fmt.Println(word)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func main() {

	flag.Parse()

	//Check two file included or not
	if len(flag.Args()) < 2 {
		exit(howToUse)
	}

	wordMap, err := getWordMap(flag.Arg(0))

	if err != nil {
		exit(err)
	}

	if err := intersect(wordMap, flag.Arg(1)); err != nil {
		exit(err)
	}
}
