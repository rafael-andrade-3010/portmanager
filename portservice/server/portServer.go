package server

import (
	"context"
	"portservice/domain"
	pb "portservice/proto"
	"portservice/service"
)

type PortServer struct {
	pb.UnimplementedPortDomainServer
}

func (s *PortServer) SavePort(ctx context.Context, in *pb.PortList) (*pb.SavePortReply, error) {
	portsToSave := make([]*domain.Port, 0)
	for _, port := range in.Ports {
		domainPort := &domain.Port{
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
	err := service.Save(portsToSave)
	ok := err == nil
	return &pb.SavePortReply{Ok: ok}, nil
}

func (s *PortServer) GetPorts(ctx context.Context, in *pb.GetPortsRequest) (*pb.PortList, error) {
	ports, err := service.Get(in.Start, in.Limit)
	if err != nil {
		return nil, err
	}
	portArray := make([]*pb.Port, 0)
	for _, p := range ports {
		portArray = append(portArray, &pb.Port{
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

	return &pb.PortList{Ports: portArray}, nil
}
