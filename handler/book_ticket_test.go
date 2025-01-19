package handler

import (
	"booking-server/proto"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPurchaseTicket_ValidRequest(t *testing.T) {
	// Arrange
	service := NewBookingService()
	req := &proto.TicketRequest{
		Email:     "test@example.com",
		FirstName: "John",
		LastName:  "Doe",
		From:      "London",
		To:        "Paris",
		Date:      "2025-01-20",
	}

	// Act
	resp, err := service.PurchaseTicket(context.Background(), req)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "John", resp.FirstName)
	assert.Equal(t, "Doe", resp.LastName)
	assert.Equal(t, "London", resp.From)
	assert.Equal(t, "Paris", resp.To)
	assert.Equal(t, 20.0, resp.Price)
	assert.Equal(t, req.Email, resp.Email)
	assert.NotEmpty(t, resp.Seat)
}

func TestPurchaseTicket_InvalidEmail(t *testing.T) {
	// Arrange
	service := NewBookingService()
	req := &proto.TicketRequest{
		Email:     "",
		FirstName: "John",
		LastName:  "Doe",
		From:      "London",
		To:        "Paris",
	}

	// Act
	resp, err := service.PurchaseTicket(context.Background(), req)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, "email-id shouldn't be empty", err.Error())
}

func TestPurchaseTicket_InvalidFrom(t *testing.T) {
	// Arrange
	service := NewBookingService()
	req := &proto.TicketRequest{
		Email:     "test@example.com",
		FirstName: "John",
		LastName:  "Doe",
		From:      "",
		To:        "Paris",
	}

	// Act
	resp, err := service.PurchaseTicket(context.Background(), req)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, "place(from) shouldn't be empty", err.Error())
}

func TestGetReceipt_ValidRequest(t *testing.T) {
	// Arrange
	service := NewBookingService()
	req := &proto.TicketRequest{
		Email:     "test@example.com",
		FirstName: "John",
		LastName:  "Doe",
		From:      "London",
		To:        "Paris",
	}
	service.PurchaseTicket(context.Background(), req) // Ensure the ticket is purchased first

	getReq := &proto.GetUserTicketRequest{Email: "test@example.com"}

	// Act
	resp, err := service.GetReceipt(context.Background(), getReq)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "John", resp.FirstName)
	assert.Equal(t, "Doe", resp.LastName)
	assert.Equal(t, "London", resp.From)
	assert.Equal(t, "Paris", resp.To)
	assert.Equal(t, 20.0, resp.Price)
}

func TestGetReceipt_NoBookingFound(t *testing.T) {
	// Arrange
	service := NewBookingService()
	getReq := &proto.GetUserTicketRequest{Email: "nonexistent@example.com"}

	// Act
	resp, err := service.GetReceipt(context.Background(), getReq)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, "no booking found for user: nonexistent@example.com", err.Error())
}

func TestGetAvailableSeats_ValidRequest(t *testing.T) {
	// Arrange
	service := NewBookingService()
	service.AvailableSeat["2025-01-20"] = []string{"A1", "A2", "A3", "A4"}

	req := &proto.GetAvailableSeatsRequest{
		Date: "2025-01-20",
	}

	// Act
	resp, err := service.GetAvailableSeats(context.Background(), req)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, []string{"A1", "A2", "A3", "A4"}, resp.SeatNumbers)
}

func TestGetAvailableSeats_NoSeatsAvailable(t *testing.T) {
	// Arrange
	service := NewBookingService()
	req := &proto.GetAvailableSeatsRequest{
		Date: "2025-01-20",
	}

	// Act
	resp, err := service.GetAvailableSeats(context.Background(), req)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, seats, resp.SeatNumbers) // Fallback to default `seats`
}

func TestUpdateUserSeat_ValidRequest(t *testing.T) {
	// Arrange
	service := NewBookingService()
	req := &proto.TicketRequest{
		Email:     "test@example.com",
		FirstName: "John",
		LastName:  "Doe",
		From:      "London",
		To:        "Paris",
		Date:      "2025-01-20",
	}
	service.PurchaseTicket(context.Background(), req) // Ensure the ticket is purchased first

	updateReq := &proto.UpdateUserSeatRequest{
		Email:      "test@example.com",
		Date:       "2025-01-20",
		Seatnumber: "A2",
	}

	// Act
	resp, err := service.UpdateUserSeat(context.Background(), updateReq)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "A2", resp.Seat)
}
