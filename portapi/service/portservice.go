package service

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"log"
	portgrpc "portapi/proto"
)

func SavePorts(ports []*portgrpc.Port) error {
	serviceHost := getEnv("SERVICE_HOST", "localhost")
	log.Printf("Service Host %v", serviceHost)
	conn, err := grpc.Dial(serviceHost+":50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return err
	}
	defer conn.Close()
	c := portgrpc.NewPortDomainClient(conn)

	r, err := c.SavePort(context.Background(), &portgrpc.PortList{Ports: ports})
	if err != nil {
		return err
	}
	if !r.GetOk() {
		return errors.Errorf("Error saving into downstream service %v", r.Ok)
	}
	return nil
}

func GetPorts(start, limit int32) ([]*portgrpc.Port, error) {
	conn, err := grpc.Dial(getEnv("SERVICE_HOST", "localhost")+":50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := portgrpc.NewPortDomainClient(conn)
	res, err := c.GetPorts(context.Background(), &portgrpc.GetPortsRequest{})
	if err != nil {
		return nil, err
	}
	return res.Ports, nil
}
