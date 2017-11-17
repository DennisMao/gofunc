package main

import (
	"net"
	"runtime"

	"cnt"
	"dev_gPPC_Counter_Server/models/counter"

	"github.com/astaxie/beego/logs"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	RPC_COUNTER_PORT = "41005"
)

// Cnt RPC服务
var server CounterServer

type CounterServer struct{}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	logs.Debug("RPC COUNTER 服务端")
	// 开启TCP监听
	lis, err := net.Listen("tcp", ":"+RPC_COUNTER_PORT)
	if err != nil {
		logs.Error("failed to listen: %v", err)
	}

	// 新建gRPC服务器实例
	s := grpc.NewServer()

	// 将Cat的RPC服务绑定到 Server实例中
	cnt.RegisterCounterServer(s, &CounterServer{})
	s.Serve(lis)

	logs.Debug("grpc server in: %s", RPC_COUNTER_PORT)
}

// 增加
func (t *CounterServer) Add(ctx context.Context, request *cnt.CntRq) (*cnt.CntRp, error) {
	err := counter.Add(request.Uid, request.Num)

	//	测试用
	_, lists := counter.ShowList()
	logs.Debug(lists)

	if err != nil {
		return &cnt.CntRp{
			Errcode: "00001",
			Errmsg:  err.Error(),
		}, nil
	}

	return &cnt.CntRp{
		Errcode: "00000",
	}, nil
}

//	获取
func (t *CounterServer) Get(ctx context.Context, request *cnt.CntRq) (*cnt.CntRp, error) {
	err, num := counter.Get(request.Uid)
	if err != nil {
		return &cnt.CntRp{
			Errcode: "00001",
			Errmsg:  err.Error(),
		}, nil
	}

	return &cnt.CntRp{
		Errcode: "00000",
		Num:     num,
	}, nil
}

// 删除
func (t *CounterServer) Del(ctx context.Context, request *cnt.CntRq) (*cnt.CntRp, error) {

	err := counter.Del(request.Uid)
	if err != nil {
		return &cnt.CntRp{
			Errcode: "00001",
			Errmsg:  err.Error(),
		}, nil
	}

	return &cnt.CntRp{
		Errcode: "00000",
	}, nil
}
