package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

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

func ExtractSubjectEmenta(subjectURL string, c *colly.Collector) (ementa string) {
	c.Visit(subjectURL)

	ementa = ""
	return
}

func ExtractSubjectHours(element *colly.HTMLElement) (ementa string) {
	hoursText := element.Attr("innerText")
	fmt.Println("****** \nOn: ", element.Request.URL)
	fmt.Println("text: ", hoursText)
	fmt.Println("******")

	ementa = ""
	return
}
