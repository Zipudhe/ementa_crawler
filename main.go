package main

import (
	"fmt"

	"github.com/Zipudhe/ementa_crawler/handlers"
	"github.com/Zipudhe/ementa_crawler/types"
	"github.com/Zipudhe/ementa_crawler/utils"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.MaxDepth(2),
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 2}) // Controle de goroutine através do colly

	subjectList := make([]types.Subject, 0)

	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Vistin: ", r.URL)
	})

	c.OnHTML("a[href]", func(element *colly.HTMLElement) {
		subject := handlers.HandleSubjectInfo(element, c)
		subjectList = append(subjectList, subject)
	})

	c.OnHTML("th[colspan='3']", func(element *colly.HTMLElement) {
		utils.ExtractSubjectHours(element)
	})

	// c.OnHTML(".even > td[colspan='5']", func(element *colly.HTMLElement) {
	// 	subject := handlers.HandleSubjectInfo(element, c)
	// 	subjectList = append(subjectList, subject)
	// })

	fmt.Println("TESTANDO")
	c.Visit("https://alunoweb.ufba.br/SiacWWW/ListaDisciplinasEmentaPublico.do?cdCurso=112140&nuPerCursoInicial=20231")

	c.Wait()
}
