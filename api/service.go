package api

type Service struct {
	Service		string `json:"service"`
	Purposes	[]*Purpose `json:"purposes"`
}

func (s *Service) AddPurpose(purpose *Purpose) {
	s.Purposes = append(s.Purposes, purpose)
}

func NewServiceMultiPurpose(serviceName string, purposes []*Purpose) *Service {
	return &Service{
		Service:	serviceName,
		Purposes:	purposes,
	}
}

func NewServiceSinglePurpose(serviceName string, purpose *Purpose) *Service {
	s := &Service{
		Service: serviceName,
	}

	s.AddPurpose(purpose)

	return s
}