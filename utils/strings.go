package utils

import (
	"regexp"
	"strings"
)

func NotEmpty(s string) bool {
	return s != ""
}

func CaptureNonDigits(s string) (capturedText string, isText bool, capturedNumbers string) {
	textRegex := regexp.MustCompile(`\D+`)
	numericRegex := regexp.MustCompile(`\d[\d\s]+`)
	foundText := strings.ReplaceAll(textRegex.FindString(s), " ", "")
	foundNumbers := numericRegex.FindString(s)
	return foundText, len(foundText) > 0, foundNumbers
}
