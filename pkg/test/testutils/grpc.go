package testutils

import (
	"context"
	"net"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

type GrpcServer struct {
	*grpc.Server
	lis *bufconn.Listener
}

func (s *GrpcServer) Start() {
	go func() {
		if err := s.Serve(s.lis); err != nil {
			panic(err)
		}
	}()
}

func NewGrpcServer(ctx context.Context) (*GrpcServer, *grpc.ClientConn, error) {
	listener := bufconn.Listen(1024 * 1024)
	s := grpc.NewServer()
	conn, err := grpc.DialContext(
		ctx,
		"",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return listener.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	return &GrpcServer{
		Server: s,
		lis:    listener,
	}, conn, nil
}
