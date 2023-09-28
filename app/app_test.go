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

		apis.FnReserveFlight = func() (int, error) {
			return 1, nil
		}
		apis.FnReserveRoom = func() (int, error) {
			return 1, nil
		}

		reservation, err := application.ReserveVacation()
		require.NoError(t, err)
		require.NotNil(t, reservation)

		assert.NotZero(t, reservation.RoomReservationID, "invalid room reservation")
		assert.NotZero(t, reservation.FlightReservationID, "invalid flight reservation")
	})

	t.Run("room reservation failed", func(t *testing.T) {
		apis := test.NewStub()

		application := app.NewApp(apis, apis)

		apis.FnReserveFlight = func() (int, error) {
			return 1, nil
		}
		apis.FnReserveRoom = func() (int, error) {
			return 0, errors.New("room unavailable")
		}

		reservation, err := application.ReserveVacation()
		require.Error(t, err)
		require.NotNil(t, reservation)

		assert.Zero(t, reservation.RoomReservationID, "invalid room reservation")
		assert.NotZero(t, reservation.FlightReservationID, "invalid flight reservation")
	})

	t.Run("flight reservation failed", func(t *testing.T) {
		apis := test.NewStub()

		application := app.NewApp(apis, apis)

		apis.FnReserveFlight = func() (int, error) {
			return 0, errors.New("flight unavailable")
		}
		apis.FnReserveRoom = func() (int, error) {
			return 1, nil
		}

		reservation, err := application.ReserveVacation()
		require.Error(t, err)
		require.NotNil(t, reservation)

		assert.NotZero(t, reservation.RoomReservationID, "invalid room reservation")
		assert.Zero(t, reservation.FlightReservationID, "invalid flight reservation")
	})
}
