package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func ExtractDisciplinaFromURL(url string) (string, error) {
	re := regexp.MustCompile(`[A-Z]+\d+`)

	matches := re.FindAllString(url, -1)

	if len(matches) != 1 {
		return "", errors.New("expected exactly one discipline")
	}

	return matches[0], nil
}

func ExtractSubjectName(subjectString string) (string, string, int) {
	baseURL := "https://alunoweb.ufba.br"

	subjectCode := strings.Split(subjectString, "=")
	link := baseURL + subjectString
	code := subjectCode[1][0:6]

	period, err := strconv.Atoi(subjectCode[2])
	if err != nil {
		panic(err)
	}

	return code, link, period
}

// func ExtractSubjectEmenta(subjectURL string, c *colly.Collector) (ementa string) {
//
// }

func ExtractSubjectHours(element *colly.HTMLElement) int {
	hoursText := element.Text

	hours, error := extractAndConvertNumber(hoursText)
	if error != nil {
		fmt.Println("Failed to convert hours")
	}
	return hours
}

func extractAndConvertNumber(text string) (int, error) {
	// Regular expression to match numbers
	re := regexp.MustCompile(`\d+`)

	// Find all numbers in the text
	matches := re.FindAllString(text, -1)

	// Check if we found exactly one number
	if len(matches) != 1 {
		return 0, errors.New("expected exactly one number")
	}

	// Get the matched number
	numberStr := matches[0]

	// Trim whitespace
	trimmed := strings.TrimSpace(numberStr)

	// Convert to integer
	result, err := strconv.Atoi(trimmed)
	if err != nil {
		return 0, err
	}

	return result, nil
}
