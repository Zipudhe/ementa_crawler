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
	// ementa := utils.ExtractSubjectEmenta(link, element.Request.Visit(link))
	// fmt.Println("ementa: ", ementa)

	// subject = types.Subject{code, link, nome, "", period}
	subject := types.Subject{Code: code, Link: link, Name: nome, Ementa: "", Period: period}

	return subject
}
