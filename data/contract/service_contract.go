package contract

type ServiceContract interface {
	fetchDailyHoroscope(horoscopeType string)
}