package room

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"log"
	"pms.com/project-api/config"
	"pms.com/project-common/discovery"
	"pms.com/project-common/logs"
	"pms.com/project-grpc/room"
)

var RoomServiceClient room.RoomServiceClient

func InitRpcRoomClient() {
	etcdRegister := discovery.NewResolver(config.C.EtcdConfig.Addrs, logs.LG)
	resolver.Register(etcdRegister)
	conn, err := grpc.NewClient("etcd:///room", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	RoomServiceClient = room.NewRoomServiceClient(conn)
}
