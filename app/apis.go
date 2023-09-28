package app

type HotelAPI interface {
	ReserveRoom(adults, children int) (int, error)
}

type AirlineAPI interface {
	ReserveFlight(seats int) (int, error)
}
