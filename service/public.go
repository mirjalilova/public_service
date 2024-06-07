package service

import (
	"context"
	pb "public/genproto"
	"public/storage"
)

type PublicService struct {
	stg storage.StorageI
	pb.UnimplementedPublicServiceServer
}

func NewPublicService(stg storage.StorageI) *PublicService {
	return &PublicService{stg: stg}
}

func (p *PublicService) Create(ctx context.Context, req *pb.PublicCreate) (*pb.Void, error) {
	res, err := p.stg.Public().Create(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *PublicService) GetById(ctx context.Context, req *pb.GetByIdReq) (*pb.PublicRes, error) {
	res, err := p.stg.Public().GetById(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *PublicService) GetAlls(ctx context.Context, req *pb.GetAllPublicsRequest) (*pb.GetAllPublicsResponse, error) {
	res, err := p.stg.Public().GetAll(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *PublicService) Update(ctx context.Context, req *pb.PublicUpdate) (*pb.Void, error) {
	res, err := p.stg.Public().Update(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *PublicService) Delete(ctx context.Context, req *pb.GetByIdReq) (*pb.Void, error) {
	res, err := p.stg.Public().Delete(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
