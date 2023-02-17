package core

import "fmt"

func GetUrl(signName string) string {
	url := fmt.Sprintf("https://www.mynet.com/kadin/burclar-astroloji/%v-burcu-gunluk-yorumu.html?day=15&month=2&year=2023", signName)
	return url
}
