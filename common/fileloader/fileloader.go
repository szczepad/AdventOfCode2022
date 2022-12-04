package fileloader

import (
	"bufio"
	"log"
	"os"
)

func LoadFile(filePath string) []string {
	result := []string{}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}
