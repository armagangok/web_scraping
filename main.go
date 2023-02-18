package main

import (
	"armaganModules/core"
	"armaganModules/data/service"
	"fmt"
	// "reflect"
)

func main() {
	var first string

	for {
		printMenu()
		fmt.Scanln(&first)

		chooseHoroscope(first)
	}

}

func chooseHoroscope(first string) {
	switch first {
	case "0":
		service.FetchDailyHoroscope(string(core.Aslan))
	case "1":
		service.FetchDailyHoroscope(string(core.Yay))
	case "2":
		service.FetchDailyHoroscope(string(core.Oglak))
	case "3":
		service.FetchDailyHoroscope(string(core.Kova))
	case "4":
		service.FetchDailyHoroscope(string(core.Ikizler))
	case "5":
		service.FetchDailyHoroscope(string(core.Yengec))
	case "6":
		service.FetchDailyHoroscope(string(core.Basak))
	case "7":
		service.FetchDailyHoroscope(string(core.Balik))
	case "8":
		service.FetchDailyHoroscope(string(core.Akrep))
	case "9":
		service.FetchDailyHoroscope(string(core.Terazi))
	case "10":
		service.FetchDailyHoroscope(string(core.Koc))
	case "11":
		service.FetchDailyHoroscope(string(core.Boga))
	}
}

func printMenu() {
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

}
