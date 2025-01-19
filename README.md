# Booking-server

This is a simple API for booking train tickets between different stations. It uses gRPC for communication and allows users to book, manage, and view their ticket details.

## How to Run the Project
### 1. Build the Project 

To build the project, use the following command:

```
make build
```

### 2. Run the Project
To start the project, use the following command:

```
make run
```
### 3. Stop the Project
To stop the project, use the following command:

```
make stop
```
##### Note: The make stop command is assumed to stop the running project. If your Makefile has a different target for stopping, please modify accordingly.


## How to Test the API
You can use grpcurl to interact with the gRPC server and test the available APIs.

### Install grpcurl
To install grpcurl, run the following command:

```
brew install grpcurl
```

Once grpcurl is installed, you can test the various APIs listed below.

## API Testing Commands
### 1. Book a Ticket
To book a ticket, run the following command:

```

grpcurl -plaintext -d '{"first_name":"bala","last_name":"murugan","from":"XX","to":"yyy","email":"bala@gmail.com","date":"2025-01-19"}' localhost:8089 proto.TicketBookingService/PurchaseTicket

```


### 2. Get Booking Details
To fetch booking details for a user, run the following command:

```

grpcurl -plaintext -d '{"email":"cccc@bala"}' localhost:8089 proto.TicketBookingService/GetReceipt

```
This will retrieve the booking details (such as from, to, seat, and price) for the user with the provided email.

### 3. Get Seat Allocation
To view the seat allocation for a user, run the following command:

```

grpcurl -plaintext -d '{"email":"cccc@bala"}' localhost:8089 proto.TicketBookingService/GetAllocationSeats

```
This command will return the seat allocation details for the specified user based on their email.


### 4. Cancel a Booking
To cancel a booking, run the following command:

```

grpcurl -plaintext -d '{"email":"cccc@bala","seatnumber":"B7","date":"2025-01-19"}' localhost:8089 proto.TicketBookingService/CancelBookingTicket

```

This command will cancel the ticket for the user with the specified email and seat number on the given date.

### 5. Get Available Seats
To get a list of available seats for a specific date, run the following command:

```

grpcurl -plaintext -d '{"date":"2025-01-19"}' localhost:8089 proto.TicketBookingService/GetAvailableSeats

```
This will return the list of available seats for the specified date.


### 6. Update User Seat
To update a user's seat, run the following command:

```

grpcurl -plaintext -d '{"email":"cccc@bala","date":"2025-01-19","seatnumber":"A6"}' localhost:8089 proto.TicketBookingService/UpdateUserSeat

```
This command allows the user to update their seat allocation to a new seat number.


### Additional Notes
1. Make sure your gRPC server is running and accessible on localhost:8089 before testing the APIs.
2. Ensure the TicketBookingService.proto file is correctly compiled and the generated gRPC server code is in place.



