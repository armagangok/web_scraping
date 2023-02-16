package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	fetchTop250MoviesFromImd()

}

func fetchTop250MoviesFromImd() {
	response, httpError := http.Get("https://www.imdb.com/chart/top")

	if response.StatusCode != 200 {
		fmt.Println("ERROR: ", response.StatusCode)
	} else {
		document, readerError := goquery.NewDocumentFromReader(response.Body)
		document.Find(".lister-list").Each(func(i int, selection *goquery.Selection) {
			title := selection.Find("a").Text()
			fmt.Println(title)

			print(readerError)
		})
	}

	print(httpError)
}
