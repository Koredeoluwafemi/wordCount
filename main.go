package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Pair struct {
	Key   string
	Value int
}

type Wordstring []Pair

func (w Wordstring) Len() int           { return len(w) }
func (w Wordstring) Swap(i, j int)      { w[i], w[j] = w[j], w[i] }
func (w Wordstring) Less(i, j int) bool { return w[i].Value > w[j].Value }

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter Text: ")
		// reads user input until \n by default
		scanner.Scan()
		// Holds the string that was scanned
		text := scanner.Text()

		if text != "" {
			//convert string to array
			sortArray := highlighter(text)

			for i, item := range sortArray {
				fmt.Println(i+1, " ", item.Key, "-", item.Value)
			}
		}

		if len(text) != 0 {
			fmt.Println(text)
		} else {
			// exit if user entered an empty string
			break
		}

	}

	// handle error
	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}
}

func highlighter(text string) []Pair {
	array := strings.Split(text, " ")
	newArray := make(map[string]int)
	if len(array) > 0 {
		for _, item := range array {
			newArray[item] += 1
		}
	}

	sortArray := make(Wordstring, len(newArray))

	i := 0
	for key, value := range newArray {
		sortArray[i] = Pair{key, value}
		i++
	}

	sort.Sort(sortArray)
	arrayLen := len(sortArray)
	if arrayLen > 10 {
		return sortArray[0:10]
	}

	return sortArray[0:arrayLen]
}
