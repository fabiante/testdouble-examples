package test

type Stub struct {
	FnReserveFlight func() (int, error)
	FnReserveRoom   func() (int, error)
}

func NewStub() *Stub {
	return &Stub{}
}

// ReserveFlight implements Stub.
func (s *Stub) ReserveFlight() (int, error) {
	if s.FnReserveFlight != nil {
		return s.FnReserveFlight()
	}
	panic(ErrNotImplemented)
}

// ReserveRoom implements Stub.
func (s *Stub) ReserveRoom() (int, error) {
	if s.FnReserveRoom != nil {
		return s.FnReserveRoom()
	}
	panic(ErrNotImplemented)
}
