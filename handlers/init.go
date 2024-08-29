package handlers

import (
	"github.com/Zipudhe/ementa_crawler/types"
	"github.com/Zipudhe/ementa_crawler/utils"
	"github.com/gocolly/colly"
)

func HandleSubjectInfo(element *colly.HTMLElement, collector *colly.Collector) types.Subject {
	subjectUri := element.Attr("href")

	nome := element.Text
	code, link, period := utils.ExtractSubjectName(subjectUri)
	element.Request.Visit(link)

	subject := types.Subject{Code: code, Link: link, Name: nome, Ementa: "", Period: period}

	return subject
}

func HandleSubjectHoras(element *colly.HTMLElement) int {
	subjectHoras := utils.ExtractSubjectHours(element)

	return subjectHoras
}
