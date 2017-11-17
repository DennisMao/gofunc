package main

import (
	"errors"
	"log"
	"runtime"

	"cnt"

	"github.com/astaxie/beego/logs"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	RPC_COUNTER_IP   = "127.0.0.1"
	RPC_COUNTER_PORT = "41005"
	RPC_COUNTER_UID  = "u2701" //用户编号
)

// Inf RPC服务
type CarServer struct{}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	logs.Debug("RPC COUNTER 客户端")

	//测试RPC COUNTER服务
	err, num := RPC_Counter_Get()
	if err != nil {
		logs.Debug("获取计数值失败，原因:" + err.Error())
	} else {
		logs.Debug("获取到计数值=", num)
	}

	err = RPC_Counter_Add("1")
	if err != nil {
		logs.Debug("获取计数值自增，原因:" + err.Error())
	} else {
		logs.Debug("计数值自增成功")
	}

	//测试RPC COUNTER服务
	err, num = RPC_Counter_Get()
	if err != nil {
		logs.Debug("获取计数值失败，原因:" + err.Error())
	} else {
		logs.Debug("获取到计数值=", num)
	}
}

// RPC Counter Add
func RPC_Counter_Add(num string) error {
	//建立连接
	conn, err := grpc.Dial(RPC_COUNTER_IP+":"+RPC_COUNTER_PORT,
		grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Error connect", err.Error())
		return err
	}
	defer conn.Close()
	client := cnt.NewCounterClient(conn)

	resp, err := client.Add(context.Background(),
		&cnt.CntRq{
			Uid: RPC_COUNTER_UID,
			Num: num,
		})
	if err != nil {
		return err
	}

	if resp.Errcode != "00000" {
		return errors.New(resp.Errmsg)
	}

	return nil
}

// RPC Counter Get
func RPC_Counter_Get() (error, string) {
	//建立连接
	conn, err := grpc.Dial(RPC_COUNTER_IP+":"+RPC_COUNTER_PORT,
		grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Error connect", err.Error())
		return err, ""
	}
	defer conn.Close()
	client := cnt.NewCounterClient(conn)

	resp, err := client.Get(context.Background(),
		&cnt.CntRq{
			Uid: RPC_COUNTER_UID,
		})
	if err != nil {
		return err, ""
	}

	if resp.Errcode != "00000" {
		return errors.New(resp.Errmsg), ""
	}

	return nil, resp.Num
}
