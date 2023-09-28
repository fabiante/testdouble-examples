package test

type Stub struct {
	FnReserveFlight func(seats int) (int, error)
	FnReserveRoom   func(adults, children int) (int, error)
}

func NewStub() *Stub {
	return &Stub{}
}

// ReserveFlight implements Stub.
func (s *Stub) ReserveFlight(seats int) (int, error) {
	if s.FnReserveFlight != nil {
		return s.FnReserveFlight(seats)
	}
	panic(ErrNotImplemented)
}

// ReserveRoom implements Stub.
func (s *Stub) ReserveRoom(adults, children int) (int, error) {
	if s.FnReserveRoom != nil {
		return s.FnReserveRoom(adults, children)
	}
	panic(ErrNotImplemented)
}
