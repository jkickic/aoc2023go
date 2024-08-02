package utils

import (
	"log"
	"os"
)

func LoadFileContents(filename string) string {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}
