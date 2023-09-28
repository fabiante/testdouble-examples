package app

// This file contains go generate commands used to generate mock implementations.
//
// Generator used: https://github.com/matryer/moq
// Install via: go install github.com/matryer/moq@latest

//go:generate moq -pkg test -out test/mocks_api_hotel.go . HotelAPI

//go:generate moq -pkg test -out test/mocks_api_airline.go . AirlineAPI
