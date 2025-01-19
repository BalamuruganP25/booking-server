package main

import (
	"booking-server/handler"
	"booking-server/proto"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("GRPC server started")
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	serverRegister := grpc.NewServer()
	ticketBookingService := handler.NewBookingService()
	proto.RegisterTicketBookingServiceServer(serverRegister, ticketBookingService)
	reflection.Register(serverRegister)
	err = serverRegister.Serve(lis)
	if err != nil {
		log.Fatalf("failed to listen server: %v", err)
	}
}
