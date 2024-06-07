package client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "public/genproto"
)

type GrpcClients struct {
	PublicService pb.PublicServiceClient
	PartyService  pb.PartyServiceClient
}

func NewGrpcClients() (*GrpcClients, error) {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &GrpcClients{
		PublicService: pb.NewPublicServiceClient(conn),
		PartyService:  pb.NewPartyServiceClient(conn),
	}, nil
}
