package main

import (
	"inf"
	"log"
	"net"
	"runtime"
	"strconv"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = "41005"
)

// Inf RPC服务
type InfServer struct{}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	log.Println("Begin gprc")
	// 开启TCP监听
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 新建gRPC服务器实例
	s := grpc.NewServer()

	// 将INF的RPC服务绑定到 Server实例中
	inf.RegisterDataServer(s, &Data{})
	s.Serve(lis)

	log.Println("grpc server in: %s", port)
}

// 定义方法
func (t *Data) GetUser(ctx context.Context, request *inf.UserRq) (response *inf.UserRp, err error) {
	response = &inf.UserRp{
		Name: strconv.Itoa(int(request.Id)) + ":test",
	}
	return response, err
}
