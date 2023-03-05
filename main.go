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
		core.PrintMenu()
		fmt.Scanln(&horoscopeUserInput)

		core.PrintTimeMenu()
		fmt.Scanln(&horoscopeTimeUserInput)

		timeUserInput := string(core.ChooseHoroscopeTime(horoscopeTimeUserInput))

		chooseHoroscope(horoscopeUserInput, timeUserInput)
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
