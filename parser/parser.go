package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	rating, err := GetTexts("https://www.f1news.ru/Championship/2025/personpoints.shtml", ".post_body")
	if err != nil {
		log.Println(err)
	}

	for _, r := range rating {
		fmt.Println("-", r)
	}
}

func GetTexts(url, selector string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var texts []string
	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		texts = append(texts, s.Text())
	})
	return texts, nil
}
