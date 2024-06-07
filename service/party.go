package service

import (
	"context"
	pb "public/genproto"
	"public/storage"
)

type PartyService struct {
	pb.UnimplementedPartyServiceServer
	stg storage.StorageI
}

func NewPartyService(stg storage.StorageI) *PartyService {
	return &PartyService{stg: stg}
}

func (p *PartyService) Create(ctx context.Context, req *pb.PartyCreate) (*pb.Void, error) {
	res, err := p.stg.Party().Create(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *PartyService) GetById(ctx context.Context, req *pb.GetByIdReq) (*pb.PartyRes, error) {
	res, err := p.stg.Party().GetById(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *PartyService) GetAlls(ctx context.Context, req *pb.GetAllPartysRequest) (*pb.GetAllPartysResponse, error) {
	res, err := p.stg.Party().GetAll(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *PartyService) Update(ctx context.Context, req *pb.PartyUpdate) (*pb.Void, error) {
	res, err := p.stg.Party().Update(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *PartyService) Delete(ctx context.Context, req *pb.GetByIdReq) (*pb.Void, error) {
	res, err := p.stg.Party().Delete(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
