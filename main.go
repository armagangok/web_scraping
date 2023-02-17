package main

import (
	"armaganModules/core"
	"fmt"
	"net/http"

	// "reflect"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	var first string

	for {

		printMenu(first)

		switch first {
		case "0":
			fetchDailyHoroscope(string(core.Aslan))
		case "1":
			fetchDailyHoroscope(string(core.Yay))
		case "2":
			fetchDailyHoroscope(string(core.Kova))
		case "3":
			fetchDailyHoroscope(string(core.Oglak))
		case "4":
			fetchDailyHoroscope(string(core.Ikizler))
		case "5":
			fetchDailyHoroscope(string(core.Yengec))
		case "6":
			fetchDailyHoroscope(string(core.Basak))
		case "7":
			fetchDailyHoroscope(string(core.Oglak))
		case "8":
			fetchDailyHoroscope(string(core.Oglak))
		}
	}

	// fetchDailyHoroscope("yay")
}

func printMenu(userChoice string) {
	fmt.Println("[0] Aslan ")
	fmt.Println("[1] Yay")
	fmt.Println("[2] Oglak")
	fmt.Println("[3] kova")
	fmt.Println("[4] İkizker")
	fmt.Println("[5] Yengec")
	fmt.Println("[6] Basak")
	fmt.Println("[7] Balik")
	fmt.Println("[8] Akrep")
	fmt.Println("[9] Terazi")
	fmt.Println("[10] Koc")
	fmt.Println("[10] Boga")
	fmt.Println("[Q/q/e] Exit")
	getUserInput(userChoice)
}

func getUserInput(first string) {
	fmt.Scanln(&first)
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
