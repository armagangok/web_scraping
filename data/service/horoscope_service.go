package service

import (
	"armaganModules/core"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type ServiceContract interface {
	fetchDailyHoroscope(horoscopeType string)
}

type ServiceImp struct {
	ServiceContract
}

func fetchDailyHoroscope(horoscopeType string) {

	url := core.GetUrl(horoscopeType)

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
