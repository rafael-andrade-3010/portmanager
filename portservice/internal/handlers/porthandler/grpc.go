package porthandler

import (
	"context"
	"portservice/internal/core/domain"
	"portservice/internal/core/ports"
)

type GrpcHandler struct {
	service ports.PortService
	UnimplementedPortDomainServer
}

func NewGrpcHandler(service ports.PortService) *GrpcHandler {
	return &GrpcHandler{
		service: service,
	}
}

func (s *GrpcHandler) SavePort(ctx context.Context, in *PortList) (*SavePortReply, error) {
	portsToSave := make([]domain.Port, 0)
	for _, port := range in.Ports {
		domainPort := domain.Port{
			Key:         port.Key,
			Name:        port.Name,
			City:        port.City,
			Country:     port.Country,
			Alias:       port.Alias,
			Coordinates: port.Coordinates,
			Province:    port.Province,
			Timezone:    port.Timezone,
			Unlocs:      port.Unlocs,
			Code:        port.Code,
		}
		portsToSave = append(portsToSave, domainPort)
	}
	err := s.service.Create(portsToSave)
	ok := err == nil
	return &SavePortReply{Ok: ok}, nil
}

func (s *GrpcHandler) GetPorts(ctx context.Context, in *GetPortsRequest) (*PortList, error) {
	//in.Start, in.Limit
	ports, err := s.service.GetAll()
	if err != nil {
		return nil, err
	}
	portArray := make([]*Port, 0)
	for _, p := range ports {
		portArray = append(portArray, &Port{
			Key:         p.Key,
			Name:        p.Name,
			City:        p.City,
			Country:     p.Country,
			Alias:       p.Alias,
			Coordinates: p.Coordinates,
			Province:    p.Province,
			Timezone:    p.Timezone,
			Unlocs:      p.Unlocs,
			Code:        p.Code,
		})
	}

	return &PortList{Ports: portArray}, nil
}
