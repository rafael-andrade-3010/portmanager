package porthandler

import (
	"context"
	"log"
	"net"
	"portservice/internal/core/service/portservice"
	"portservice/internal/repositories/portrepo/kv"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()
	repo := kv.New()
	service := portservice.New(repo)
	handler := NewGrpcHandler(service)

	RegisterPortDomainServer(server, handler)

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func TestSave(t *testing.T) {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := NewPortDomainClient(conn)

	request := &PortList{Ports: []*Port{{
		Key:         "1",
		Name:        "1",
		City:        "1",
		Country:     "1",
		Alias:       nil,
		Coordinates: nil,
		Province:    "1",
		Timezone:    "1",
		Unlocs:      nil,
		Code:        "5001",
	}}}

	response, err := client.SavePort(ctx, request)

	if response == nil || !response.Ok {
		t.Errorf("Expected %v and got %v", true, response.Ok)
	}

}

func TestSaveAndGet(t *testing.T) {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := NewPortDomainClient(conn)

	request := &PortList{Ports: []*Port{{
		Key:         "1",
		Name:        "1",
		City:        "1",
		Country:     "1",
		Alias:       nil,
		Coordinates: nil,
		Province:    "1",
		Timezone:    "1",
		Unlocs:      nil,
		Code:        "5001",
	}}}

	response, err := client.SavePort(ctx, request)

	if response == nil || !response.Ok {
		t.Errorf("Expected %v and got %v", true, response.Ok)
	}

	ports, err := client.GetPorts(ctx, &GetPortsRequest{})
	if err != nil {
		t.Errorf("Got error %v", err)
	}
	if ports == nil || len(ports.Ports) <= 0 {
		t.Errorf("Unexpected response %v", ports)
	}
}
