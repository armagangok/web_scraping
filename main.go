package main

import (
	"fmt"
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"armaganModules/core"
)

func main() {
	fetchDailyHoroscope("yay")
}

func fetchDailyHoroscope(signName string) {

	url := core.GetUrl(signName)

	response, httpError := http.Get(url)

	if response.StatusCode != 200 {
		fmt.Println("ERROR: ", response.StatusCode)
	} else {
		document, readerError := goquery.NewDocumentFromReader(response.Body)
		document.Find(".detail-content-inner").Each(func(i int, selection *goquery.Selection) {
			title := selection.Find("p").Text()
			fmt.Println(title)

			print(readerError)
		})
	}

	print(httpError)
}

