package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

const baseURL = "https://it.wikipedia.org/wiki/Settore_scientifico-disciplinare"

func main() {
	doc, err := goquery.NewDocument(baseURL)
	if err != nil {
		log.Fatal(err)
	}

	var query string

	doc.Find("table.wikitable > tbody > tr").Each(func(i int, tr *goquery.Selection) {
		tds := tr.Find("td")
		if tds.Length() == 0 {
			return
		}

		t := tds.Last().Text()
		parts := strings.Split(t, " - ")
		code := strings.TrimSpace(parts[0])
		name := strings.TrimSpace(parts[1])

		query = query + fmt.Sprintf("%v||%v##", name, code)
	})

	fmt.Println(query)
}
