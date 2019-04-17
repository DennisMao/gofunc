package main

import (
	"fmt"
	"log"
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

func main() {

	log.Println("Start http proxy server on :8080 ...")
	err := http.ListenAndServe(":8080", &Mux{})
	if err != nil {
		panic(err)
	}
}
