package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type Article struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func getAricles() []Article {
	var articles []Article

	doc, err := goquery.NewDocument("http://b.hatena.ne.jp/ctop/it")
	if err != nil {
		fmt.Print("url scarapping failed")
	}
	doc.Find(".entry-link").Each(func(num int, s *goquery.Selection) {
		title, _ := s.Attr("title")
		url, _ := s.Attr("href")
		article := Article{Title: title, URL: url}
		articles = append(articles, article)
	})
	return articles
}

func encodeArticlesToJSON(articles []Article) []byte {
	jsonBytes, err := json.Marshal(articles)
	if err != nil {
		fmt.Println("JSON Marshal error:", err)
		return nil
	}
	return jsonBytes
}

func main() {
	articles := getAricles()
	articlesJSON := encodeArticlesToJSON(articles)
	ioutil.WriteFile("articles.json", articlesJSON, os.ModePerm)
}
