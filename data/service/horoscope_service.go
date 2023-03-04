package service

import (
	"armaganModules/core"
	"armaganModules/data/contract"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type ServiceImp struct {
	contract.ServiceContract
}

func FetchDailyHoroscope(horoscopeType string, time string) {

	url := core.GetUrl(horoscopeType, time)

	response, httpError := http.Get(url)

	if response.StatusCode != 200 {
		fmt.Println("ERROR: ", response.StatusCode)
	} else {
		document, readerError := goquery.NewDocumentFromReader(response.Body)

		document.Find(".medyanet-content").Each(func(i int, selection *goquery.Selection) {
			title := selection.Find("p").Text()

			fmt.Println(title)

			print(readerError)
		})
	}

	print(httpError)
}
