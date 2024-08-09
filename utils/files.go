package utils

import (
	"log"
	"os"
	"strings"
)

func LoadFileLines(filename string) []string {
	return strings.Split(LoadFileContents(filename), "\n")
}

func LoadFileContents(filename string) string {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return strings.ReplaceAll(string(bytes), "\r", "")
}
