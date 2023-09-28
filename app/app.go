package app

import (
	"errors"
	"fmt"
)

// App implements domain specific use cases.
type App struct {
	hotel   HotelAPI
	airline AirlineAPI
}

func NewApp(hotel HotelAPI, airline AirlineAPI) *App {
	return &App{
		hotel:   hotel,
		airline: airline,
	}
}

// ReserveVacation reserves a vacation: Both a flight and a hotel room will be reserved.
func (app *App) ReserveVacation(adults, children int) (*VacationReservation, error) {
	reservation := new(VacationReservation)

	errs := make([]error, 0, 2)

	roomReservation, err := app.hotel.ReserveRoom(adults, children)
	if err != nil {
		errs = append(errs, fmt.Errorf("%w: %w", ErrReservationFailed, err))
	}

	reservation.RoomReservationID = roomReservation

	flightReservation, err := app.airline.ReserveFlight(adults + children)
	if err != nil {
		errs = append(errs, fmt.Errorf("%w: %w", ErrReservationFailed, err))
	}

	reservation.FlightReservationID = flightReservation

	return reservation, errors.Join(errs...)
}

type VacationReservation struct {
	RoomReservationID   int
	FlightReservationID int
}

var ErrReservationFailed = errors.New("reservation failed")
