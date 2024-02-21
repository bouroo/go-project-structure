package repository

import (
	"fmt"
	"log"
	"sync"

	"github.com/bouroo/go-project-structure/datasources"
	pb "github.com/bouroo/go-project-structure/pkg/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var UserAccountServiceConn = sync.Pool{
	New: func() interface{} {
		client, _ := newUserAccountServiceClient()
		return client
	},
}

func newUserAccountServiceClient() (client pb.UserAccountServiceClient, err error) {
	// if UserAccountServiceClient != nil {
	// 	return UserAccountServiceClient, nil
	// }
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	usergRPCAddr := fmt.Sprintf("%s:%d", datasources.AppConfig.GetString("service.user.grpc.host"), datasources.AppConfig.GetInt("service.user.grpc.port"))
	conn, err := grpc.Dial(usergRPCAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	// defer conn.Close()
	client = pb.NewUserAccountServiceClient(conn)
	return
}
