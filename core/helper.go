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