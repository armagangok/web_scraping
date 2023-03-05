package core

import "fmt"

func GetUrl(signName string, time string) string {
	url := fmt.Sprintf("https://www.mynet.com/kadin/burclar-astroloji/%v-burcu-%v-yorumu.html", signName, time)
	return url
}

// "https://www.mynet.com/kadin/burclar-astroloji/%v-burcu-%v-yorumu.html?day=15&month=2&year=2023", signName, time

// https://www.mynet.com/kadin/burclar-astroloji/%v-burcu-%v-yorumu.html

// https://www.mynet.com/kadin/burclar-astroloji

// https://www.mynet.com/kadin/burclar-astroloji/%v-burcu-%v-yorumu.html

// https://www.mynet.com/kadin/burclar-astroloji/%v-burcu-%v-yorumu.html

func ChooseHoroscopeTime(time string) TimeType {
	time = string(time)
	switch time {
	case "0":
		return Gunluk
	case "1":
		return Haftalik
	case "2":
		return Aylik
	case "3":
		return Yillik
	}

	return ""
}

func PrintTimeMenu() {
	timeMenuText := `
** Lütfen aramak istediğiniz burcun zaman dilimini seçiniz.
[0] Günlük
[1] Haftalık
[2] Aylık
[3] Yıllık
>`
	print(timeMenuText)
}

func PrintMenu() {
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
