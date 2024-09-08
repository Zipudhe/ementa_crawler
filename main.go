package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Zipudhe/ementa_crawler/handlers"
	"github.com/Zipudhe/ementa_crawler/types"
	"github.com/Zipudhe/ementa_crawler/utils"
	"github.com/gocolly/colly"
)

func crawlSubjects() []types.Subject {
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

	return subjectList
}

func compareSubject(ufbaSubjects []types.Subject, hours int, ementa string) {
	fmt.Println("Will compare subjects ementa: ", ementa)
	fmt.Println("Will compare subjects hours: ", hours)
	prompt := `Dadas duas ementas de matérias diferentes, faça uma análise de cada ementa, e utilizando como critério as palavras chaves de cada ementa e a quantidade de horas de cada matéria.
  Lembre se que que caso a primeira matéria tenha uma quantidade de horas maior que a segunda em 20% isso afeta a equivalência das matérias.
  Caso a ementa 2 seja mais abrangente mas tenha tópicos equivalentes isso afeta positivamente a equivalência.

  Como saída, utilize exatamente e apenas o seguinte padrão:

  "A matéria UNB tem uma equivalência de x%" onde "x" é a porcentagem definida por você

    matéria UFBA:
    Horas: "Horas matéria UFBA"
    Ementa: "ementa matéria UFBA"

    matéria UNB:
    Horas: "Horas matéria UNB"
    Ementa: "ementa matéria UNB"
  `
}

func main() {
	fmt.Println("Exatrcting UFBA subjects....")
	reader := bufio.NewReader(os.Stdin)
	subjects := crawlSubjects()

	for {
		fmt.Print("Horas da matéria: ")
		hours_str, _ := reader.ReadString('\n')
		hours_str = strings.TrimSpace(hours_str)
		hours, error := strconv.Atoi(hours_str)

		if error != nil {
			fmt.Println("Failed to parse hours")
			return
		}

		if hours == 0 {
			break
		}

		fmt.Print("Ementa da matéria: ")
		ementa, _ := reader.ReadString('\n')
		ementa = strings.TrimSpace(ementa)

		compareSubject(subjects, hours, ementa)
	}
}
