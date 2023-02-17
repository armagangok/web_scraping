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
		service.fetchDailyHoroscope()
	case "1":
		service.fetchDailyHoroscope()
	case "2":
		service.fetchDailyHoroscope(string(core.Kova))
	case "3":
		service.fetchDailyHoroscope(string(core.Oglak))
	case "4":
		service.fetchDailyHoroscope(string(core.Ikizler))
	case "5":
		service.fetchDailyHoroscope(string(core.Yengec))
	case "6":
		service.fetchDailyHoroscope(string(core.Basak))
	case "7":
		service.fetchDailyHoroscope(string(core.Oglak))
	case "8":
		service.fetchDailyHoroscope(string(core.Oglak))
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
