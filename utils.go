package main

import (
	"bufio"
	"os"
)

func read(path string) ([]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words, scanner.Err()
}

func isPlayerInList(list []string, playerToCheck string) bool {
	for _, player := range list {
		if player == playerToCheck {
			return true
		}
	}
	return false
}
