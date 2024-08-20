package main

import(
  "fmt"
  "github.com/gocolly/colly"
)

func main() {
  c := colly.NewCollector()

  c.OnRequest(func (r *colly.Request) {
    fmt.Println("Vistin: ", r.URL)
  })

  c.OnHTML("a[href]", extractSubjectName)

  c.Visit("https://alunoweb.ufba.br/SiacWWW/ListaDisciplinasEmentaPublico.do?cdCurso=112140&nuPerCursoInicial=20231")
}


func extractSubjectName(element *colly.HTMLElement) {
  fmt.Println("On extract: ", element.Name)
}
