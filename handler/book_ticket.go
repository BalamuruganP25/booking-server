package handler

import (
	"booking-server/proto"
	"context"
	"fmt"
)

func NewBookingService() *TicketBookingService {
	return &TicketBookingService{
		BookingUsers:  make(map[string]*proto.TicketResponse),
		SeatSections:  make(map[string][]BookingSeat),
		AvailableSeat: make(map[string][]string),
	}
}

func (s *TicketBookingService) PurchaseTicket(ctx context.Context, req *proto.TicketRequest) (*proto.TicketResponse, error) {
	var (
		selectedSeat   string
		remainingSeats []string
	)

	err := validatePurchaseTicketRequest(req)
	if err != nil {
		return nil, err
	}

	if val, ok := s.AvailableSeat[req.Date]; ok {
		selectedSeat, remainingSeats = getRandomSeat(val)
		s.AvailableSeat[req.Date] = remainingSeats
		s.SeatSections = StoreUserBookingDeatils(selectedSeat, req.Date, req.FirstName, req.LastName, req.Email, s.SeatSections)
	} else {
		selectedSeat, remainingSeats = getRandomSeat(seats)
		s.AvailableSeat[req.Date] = remainingSeats
		s.SeatSections = StoreUserBookingDeatils(selectedSeat, req.Date, req.FirstName, req.LastName, req.Email, s.SeatSections)
	}

	res := &proto.TicketResponse{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		From:      req.From,
		To:        req.To,
		Seat:      selectedSeat,
		Price:     20.0,
		Email:     req.Email,
	}
	s.BookingUsers[req.Email] = res

	return res, nil
}

func (s *TicketBookingService) GetReceipt(ctx context.Context, req *proto.GetUserTicketRequest) (*proto.TicketResponse, error) {
	if val, ok := s.BookingUsers[req.Email]; ok {

		response := &proto.TicketResponse{
			From:      val.From,
			To:        val.To,
			Seat:      val.Seat,
			Price:     20.0,
			Email:     val.Email,
			FirstName: val.FirstName,
			LastName:  val.LastName,
		}
		return response, nil

	}

	return nil, fmt.Errorf("no booking found for user: %s", req.Email)
}

func (s *TicketBookingService) GetAllocationSeats(ctx context.Context, req *proto.GetSeatAllocationRequest) (*proto.GetSeatAllocationResponse, error) {
	var bookingSeats []*proto.Bookingseat
	if val, ok := s.SeatSections[req.Email]; ok {

		for i := 0; i < len(val); i++ {
			if val[i].Staus == bookingActiveStatus {
				bookingSeats = append(bookingSeats, &proto.Bookingseat{
					Name:       val[i].Name,
					Date:       val[i].Date,
					Seatnumber: val[i].SeatNumber,
				})

			}

		}
		return &proto.GetSeatAllocationResponse{Bookingseats: bookingSeats}, nil

	}
	return nil, fmt.Errorf("the user doesn't book the ticket ")
}

func (s *TicketBookingService) CancelBookingTicket(ctx context.Context, req *proto.CancelBookingTicketRequest) (*proto.CancelBookingTicketResponse, error) {

	if val, ok := s.BookingUsers[req.Email]; ok {

		res := &proto.CancelBookingTicketResponse{
			From:  val.From,
			To:    val.To,
			Price: val.Price,
			Seat:  val.Seat,
			Email: val.Email,
		}

		if val, ok := s.SeatSections[req.Email]; ok {

			for i := 0; i < len(val); i++ {
				if val[i].Date == req.Date && val[i].SeatNumber == req.Seatnumber {
					val[i].Staus = bookingInActiveStatus

				}

			}
		}
		delete(s.BookingUsers, req.Email)
		return res, nil

	}

	return nil, fmt.Errorf("no ticket found")
}

func (s *TicketBookingService) GetAvailableSeats(ctx context.Context, req *proto.GetAvailableSeatsRequest) (*proto.GetAvailableSeatsResponse, error) {

	if val, ok := s.AvailableSeat[req.Date]; ok {

		res := &proto.GetAvailableSeatsResponse{
			SeatNumbers: val,
		}
		return res, nil

	}

	res := &proto.GetAvailableSeatsResponse{
		SeatNumbers: seats,
	}
	return res, nil
}

func (s *TicketBookingService) UpdateUserSeat(ctx context.Context, req *proto.UpdateUserSeatRequest) (*proto.UpdateUserSeatResponse, error) {

	if bookingVal, ok := s.BookingUsers[req.Email]; ok {

		if val, ok := s.AvailableSeat[req.Date]; ok {
			isUserSeatStatus := isSeatIsAvailable(req.Seatnumber, val)
			if isUserSeatStatus {
				remainingSeats := removeUserPickSeat(val, req.Seatnumber)
				bookingVal.Seat = req.Seatnumber
				s.BookingUsers[req.Email] = bookingVal
				s.SeatSections = StoreUserBookingDeatils(req.Seatnumber, req.Date, bookingVal.FirstName, bookingVal.LastName, bookingVal.Email, s.SeatSections)
				s.AvailableSeat[req.Date] = append(remainingSeats, bookingVal.Seat)
				return &proto.UpdateUserSeatResponse{
					From:  bookingVal.From,
					To:    bookingVal.To,
					Email: bookingVal.Email,
					Price: bookingVal.Price,
					Seat:  bookingVal.Seat,
				}, nil
			} else {
				return nil, fmt.Errorf("you are booking seats not available")
			}

		} else {

			s.AvailableSeat[req.Date] = seats
			remainingSeats := removeUserPickSeat(seats, req.Seatnumber)
			bookingVal.Seat = req.Seatnumber
			s.BookingUsers[req.Email] = bookingVal
			s.AvailableSeat[req.Date] = remainingSeats
			s.SeatSections = StoreUserBookingDeatils(req.Seatnumber, req.Date, bookingVal.FirstName, bookingVal.LastName, bookingVal.Email, s.SeatSections)
			return &proto.UpdateUserSeatResponse{
				From:  bookingVal.From,
				To:    bookingVal.To,
				Email: bookingVal.Email,
				Price: bookingVal.Price,
				Seat:  bookingVal.Seat,
			}, nil

		}

	}
	return nil, fmt.Errorf("the user doesn't book any ticket")
}

func validatePurchaseTicketRequest(req *proto.TicketRequest) error {
	if req.Email == "" {
		return fmt.Errorf("email-id shouldn't be empty")
	}
	if req.To == "" {
		return fmt.Errorf("place(to) shouldn't be empty")
	}
	if req.From == "" {
		return fmt.Errorf("place(from) shouldn't be empty")
	}
	if req.FirstName == "" {
		return fmt.Errorf("first name shouldn't be empty")
	}
	if req.LastName == "" {
		return fmt.Errorf("last name shouldn't be empty")
	}
	if req.Date == "" {
		return fmt.Errorf("date shouldn't be empty")
	}
	return nil
}
