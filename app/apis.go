package app

type HotelAPI interface {
	ReserveRoom() (int, error)
}

type AirlineAPI interface {
	ReserveFlight() (int, error)
}
