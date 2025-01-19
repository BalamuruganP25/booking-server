package handler

import "booking-server/proto"

var seats = []string{"A1", "A2", "A3", "A4", "A5", "A6", "A7", "A8", "A9", "A10", "B1", "B2", "B3", "B4", "B5", "B6", "B7", "B8", "B9", "B10"}

const (
	bookingActiveStatus   = "ACTIVE"
	bookingInActiveStatus = "INACTIVE"
)

type BookingSeat struct {
	SeatNumber string
	Date       string
	Name       string
	Staus      string
}
type TicketBookingService struct {
	BookingUsers  map[string]*proto.TicketResponse
	SeatSections  map[string][]BookingSeat
	AvailableSeat map[string][]string
	proto.UnimplementedTicketBookingServiceServer
}
