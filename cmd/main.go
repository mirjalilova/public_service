package main

import (
	"log"
	"net"
	"public/config"

	"google.golang.org/grpc"
	pb "public/genproto"
	"public/service"
	"public/storage/postgres"
)

func main() {
	cfg := config.Load()

	db, err := postgres.NewPostgresStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}

	liss, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPartyServiceServer(s, service.NewPartyService(db))
	pb.RegisterPublicServiceServer(s, service.NewPublicService(db))
	log.Printf("server listening at %v", liss.Addr())
	if err := s.Serve(liss); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
