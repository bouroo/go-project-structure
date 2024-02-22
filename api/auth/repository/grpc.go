package repository

import (
	"context"
	"sync"

	"github.com/bouroo/go-project-structure/datasources"
	pb "github.com/bouroo/go-project-structure/pkg/proto/user"
)

var UserAccountServiceConn = sync.Pool{
	New: func() interface{} {
		client, _ := newUserAccountServiceClient()
		return client
	},
}

func newUserAccountServiceClient() (client pb.UserAccountServiceClient, err error) {
	grpcConn, err := datasources.UserGRPCConn.Get(context.Background())
	if err != nil {
		return
	}
	defer datasources.UserGRPCConn.Put(grpcConn)

	client = pb.NewUserAccountServiceClient(grpcConn)
	return
}
