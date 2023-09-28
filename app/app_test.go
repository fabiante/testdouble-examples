package app_test

import (
	"errors"
	"testing"

	"github.com/fabiante/testdoubleexamples/app"
	"github.com/fabiante/testdoubleexamples/app/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWithStubs(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		apis := test.NewStub()

		application := app.NewApp(apis, apis)

		apis.FnReserveFlight = func(seats int) (int, error) {
			return 1, nil
		}
		apis.FnReserveRoom = func(adults, children int) (int, error) {
			return 1, nil
		}

		adults, children := 2, 3

		reservation, err := application.ReserveVacation(adults, children)
		require.NoError(t, err)
		require.NotNil(t, reservation)

		assert.NotZero(t, reservation.RoomReservationID, "invalid room reservation")
		assert.NotZero(t, reservation.FlightReservationID, "invalid flight reservation")
	})

	t.Run("room reservation failed", func(t *testing.T) {
		apis := test.NewStub()

		application := app.NewApp(apis, apis)

		apis.FnReserveFlight = func(seats int) (int, error) {
			return 1, nil
		}
		apis.FnReserveRoom = func(adults, children int) (int, error) {
			return 0, errors.New("room unavailable")
		}

		adults, children := 2, 3

		reservation, err := application.ReserveVacation(adults, children)
		require.Error(t, err)
		require.NotNil(t, reservation)

		assert.Zero(t, reservation.RoomReservationID, "invalid room reservation")
		assert.NotZero(t, reservation.FlightReservationID, "invalid flight reservation")
	})

	t.Run("flight reservation failed", func(t *testing.T) {
		apis := test.NewStub()

		application := app.NewApp(apis, apis)

		apis.FnReserveFlight = func(seats int) (int, error) {
			return 0, errors.New("flight unavailable")
		}
		apis.FnReserveRoom = func(adults, children int) (int, error) {
			return 1, nil
		}

		adults, children := 2, 3

		reservation, err := application.ReserveVacation(adults, children)
		require.Error(t, err)
		require.NotNil(t, reservation)

		assert.NotZero(t, reservation.RoomReservationID, "invalid room reservation")
		assert.Zero(t, reservation.FlightReservationID, "invalid flight reservation")
	})
}
