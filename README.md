# booking-server

This is a simple API for booking train tickets between different stations. It uses gRPC for communication and allows users to book, manage, and view their ticket details.

How to Run the Project
Build the Project
To build the project, use the following command:

bash
Copy
make build
Run the Project
To start the project, use the following command:

bash
Copy
make run
Stop the Project
To stop the project, use the following command:

bash
Copy
make stop
Note: The make stop command is assumed to stop the running project. If your Makefile has a different target for stopping, please modify accordingly.

How to Test the API
You can use grpcurl to interact with the gRPC server and test the available APIs.

Install grpcurl
To install grpcurl, run:

bash
Copy
brew install grpcurl
Once grpcurl is installed, you can test the various APIs listed below.

Book a Ticket
To book a ticket, run the following command:

bash
Copy
grpcurl -plaintext -d '{"first_name":"bala","last_name":"murugan","from":"XX","to":"yyy","email":"cccc@bala","date":"2025-01-19"}' localhost:8089 proto.TicketBookingService/PurchaseTicket
Get Booking Details
To fetch booking details for a user, run:

bash
Copy
grpcurl -plaintext -d '{"email":"cccc@bala"}' localhost:8089 proto.TicketBookingService/GetReceipt
Get Seat Allocation
To view the seat allocation for a user, run:

bash
Copy
grpcurl -plaintext -d '{"email":"cccc@bala"}' localhost:8089 proto.TicketBookingService/GetAllocationSeats
Cancel a Booking
To cancel a booking, run:

bash
Copy
grpcurl -plaintext -d '{"email":"cccc@bala","seatnumber":"B7","date":"2025-01-19"}' localhost:8089 proto.TicketBookingService/CancelBookingTicket
Get Available Seats
To get a list of available seats for a specific date, run:

bash
Copy
grpcurl -plaintext -d '{"date":"2025-01-19"}' localhost:8089 proto.TicketBookingService/GetAvailableSeats
Update User Seat
To update a user's seat, run:

bash
Copy
grpcurl -plaintext -d '{"email":"cccc@bala","date":"2025-01-19","seatnumber":"A6"}' localhost:8089 proto.TicketBookingService/UpdateUserSeat