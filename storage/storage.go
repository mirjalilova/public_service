package storage

import pb "public/genproto"

type StorageI interface {
	Public() PublicI
	Party() PartyI
}

type PublicI interface {
	Create(public *pb.PublicCreate) (*pb.Void, error)
	GetById(id *pb.GetByIdReq) (*pb.PublicRes, error)
	GetAll(filter *pb.GetAllPublicsRequest) (*pb.GetAllPublicsResponse, error)
	Update(public *pb.PublicUpdate) (*pb.Void, error)
	Delete(id *pb.GetByIdReq) (*pb.Void, error)
}

type PartyI interface {
	Create(party *pb.PartyCreate) (*pb.Void, error)
	GetById(id *pb.GetByIdReq) (*pb.PartyRes, error)
	GetAll(filter *pb.GetAllPartysRequest) (*pb.GetAllPartysResponse, error)
	Update(public *pb.PartyUpdate) (*pb.Void, error)
	Delete(id *pb.GetByIdReq) (*pb.Void, error)
}
