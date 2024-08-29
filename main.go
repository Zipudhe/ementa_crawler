package main

import (
	"fmt"
	"strings"

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
		subjectHoras := handlers.HandleSubjectHoras(element)
		disciplina, err := utils.ExtractDisciplinaFromURL(element.Request.URL.RequestURI())
		if err != nil {
			fmt.Println("Should handle error")
			fmt.Println(err)
		}

		for i := 0; i < len(subjectList); i++ {
			if subjectList[i].Code == disciplina {
				subjectList[i].CargaHoraria = subjectHoras
				break
			}
		}
	})

	c.OnHTML("body > :nth-child(3) :nth-child(7)", func(element *colly.HTMLElement) {
		// Checar se a url é de uma matéria
		// Pesquisar matéria da ementa e settar a ementa
		if !strings.Contains(element.Request.URL.RequestURI(), "Lista") {
			disciplina, err := utils.ExtractDisciplinaFromURL(element.Request.URL.RequestURI())
			if err != nil {
				fmt.Println("Should handle error")
				fmt.Println(err)
			}

			for i := 0; i < len(subjectList); i++ {
				if subjectList[i].Code == disciplina {
					subjectList[i].Ementa = element.Text
					break
				}
			}
		}
	})

	c.Visit("https://alunoweb.ufba.br/SiacWWW/ListaDisciplinasEmentaPublico.do?cdCurso=112140&nuPerCursoInicial=20231")

	c.Wait()
	for i := 0; i < len(subjectList); i++ {
		fmt.Println(subjectList[i])
	}
}
