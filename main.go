package main

import (
	"armaganModules/core"
	"armaganModules/data/service"
	"fmt"
	// "reflect"
)

func main() {
	var horoscopeUserInput string
	var horoscopeTimeUserInput string

	for {
		printMenu()
		fmt.Scanln(&horoscopeUserInput)

		printTimeMenu()
		fmt.Scanln(&horoscopeTimeUserInput)

		a := chooseHoroscopeTime(horoscopeTimeUserInput)

		print(a)
		print("\n")

		chooseHoroscope(horoscopeUserInput, horoscopeTimeUserInput)
	}
}

func chooseHoroscope(first string, time string) {
	switch first {
	case "0":
		service.FetchDailyHoroscope(string(core.Aslan), time)
	case "1":
		service.FetchDailyHoroscope(string(core.Yay), time)
	case "2":
		service.FetchDailyHoroscope(string(core.Oglak), time)
	case "3":
		service.FetchDailyHoroscope(string(core.Kova), time)
	case "4":
		service.FetchDailyHoroscope(string(core.Ikizler), time)
	case "5":
		service.FetchDailyHoroscope(string(core.Yengec), time)
	case "6":
		service.FetchDailyHoroscope(string(core.Basak), time)
	case "7":
		service.FetchDailyHoroscope(string(core.Balik), time)
	case "8":
		service.FetchDailyHoroscope(string(core.Akrep), time)
	case "9":
		service.FetchDailyHoroscope(string(core.Terazi), time)
	case "10":
		service.FetchDailyHoroscope(string(core.Koc), time)
	case "11":
		service.FetchDailyHoroscope(string(core.Boga), time)
	}
}

func chooseHoroscopeTime(time string) core.TimeType {
	switch time {
	case "0":
		return core.Gunluk
	case "1":
		return core.Haftalik
	case "2":
		return core.Aylik
	case "3":
		return core.Yillik
	}

	return ""
}

func printMenu() {
	horoscopeMenuText := `
** Lütfen aşağıdaki burçlardan birini seçiniz.
[0] Aslan 
[1] Yay
[2] Oglak
[3] kova
[4] İkizler
[5] Yengeç
[6] Başak
[7] Balık
[8] Akrep
[9] Terazi
[10] Koç
[10] Boğa
>`

	fmt.Println(horoscopeMenuText)

}

func printTimeMenu() {
	timeMenuText := `
** Lütfen aramak istediğiniz burcun zaman dilimini seçiniz.
[0] Günlük
[1] Haftalık
[2] Aylık
[3] Yıllık
>`
	print(timeMenuText)
}

// fmt.Println("[0] Aslan ")
// fmt.Println("[1] Yay")
// fmt.Println("[2] Oglak")
// fmt.Println("[3] kova")
// fmt.Println("[4] İkizker")
// fmt.Println("[5] Yengec")
// fmt.Println("[6] Basak")
// fmt.Println("[7] Balik")
// fmt.Println("[8] Akrep")
// fmt.Println("[9] Terazi")
// fmt.Println("[10] Koc")
// fmt.Println("[10] Boga")
// fmt.Println("[Q/q/e] Exit")
