package handler

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/bouroo/go-project-structure/api/user/usecase"
	"github.com/bouroo/go-project-structure/datasources"
	pb "github.com/bouroo/go-project-structure/pkg/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func RunGRPCServer() (err error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", datasources.AppConfig.GetInt("app.port.grpc")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(
		grpc.ConnectionTimeout(time.Second),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: time.Second * 10,
			Timeout:           time.Second * 20,
		}),
		grpc.KeepaliveEnforcementPolicy(
			keepalive.EnforcementPolicy{
				MinTime:             time.Second,
				PermitWithoutStream: true,
			}),
		grpc.MaxConcurrentStreams(5),
	)
	pb.RegisterUserAccountServiceServer(grpcServer, &usecase.UserAccountServiceServer{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return
}
