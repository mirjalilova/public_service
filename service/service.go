package service

type Service struct {
	PBS *PublicService
	PRS *PartyService
}

func NewService(pb PublicService, pr PartyService) *Service {
	return &Service{PBS: &pb, PRS: &pr}
}
