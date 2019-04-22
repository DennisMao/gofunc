package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

const (
	TargetHost = "https://www.baidu.com"
)

type Mux struct{}

// 重定向
// 127.0.0.1:8080/*-> https://www.baidu.com/*
// origin: http://127.0.0.1:8080
// target: https://www.baidu.com
func (this *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tStart := time.Now()
	tarUrlStr := fmt.Sprintf("%s%s", TargetHost, r.RequestURI)

	defer log.Printf("HTTP | %s | %d(ms) | Ori: %s | Tar: %s | RemoteIP:%s \n", tStart.Format("2006-01-02 15:04:05"), time.Now().Sub(tStart).Nanoseconds()*1000, r.URL.String(), tarUrlStr, r.RemoteAddr)

	w.WriteHeader(200)
	return
}

type tcpKeepAliveListener struct {
	*net.TCPListener
	timeout time.Duration
}

func (ln tcpKeepAliveListener) Accept() (net.Conn, error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return nil, err
	}

	if ln.timeout != 0 {
		tc.SetKeepAlive(true)
		tc.SetKeepAlivePeriod(ln.timeout)
	}

	return tc, nil
}

func newTcpKeepAlveListener(addr string, timeout time.Duration) (*tcpKeepAliveListener, error) {
	if addr == "" {
		addr = ":http"
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &tcpKeepAliveListener{ln.(*net.TCPListener), timeout}, nil
}

func main() {

	log.Println("Start http proxy server on :8080 ...")

	// http.ListenAndServe 底层使用的是http.DefaultTransport
	// 其默认就支持了TCP和Http的keepAlive
	// 所以我们直接使用  err := http.ListenAndServe(":8080", &Mux{}) 就可以开启支持KeepAlive的http服务
	// 当然我们也可以按下面的设置手动修改相关参数以实现更高性能的服务
	s := http.Server{
		Addr:        ":8080",
		Handler:     &Mux{},
		IdleTimeout: 30 * time.Second, // 在处理完一个http请求后,下一次请求来的最大时间间隔,http层面的超时
	}

	t, err := newTcpKeepAlveListener(
		s.Addr,
		6*time.Minute, // TCP层的KeepAlive超时,在处理完一个tcp请求后,下一次请求来的最大时间间隔
	)
	if err != nil {
		panic(err)
	}

	err = s.Serve(t)
	if err != nil {
		panic(err)
	}
}
