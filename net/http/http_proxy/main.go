package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	TargetHost = "https://www.baidu.com"
)

type Mux struct{}

// 反向代理
// 127.0.0.1:8080/*-> https://www.baidu.com/*
// origin: http://127.0.0.1:8080
// target: https://www.baidu.com
func (this *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tStart := time.Now()

	transport := http.DefaultTransport

	tarReq := new(http.Request)
	*tarReq = *r

	tarUrlStr := fmt.Sprintf("%s%s", TargetHost, r.RequestURI)
	defer log.Printf("HTTP | %s | %d(ms) | Ori: %s | Tar: %s | RemoteIP:%s \n", tStart.Format("2006-01-02 15:04:05"), time.Now().Sub(tStart).Nanoseconds()*1000, r.URL.String(), tarUrlStr, r.RemoteAddr)

	tarUrl, err := url.Parse(tarUrlStr)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	// 替换为新的Url
	tarReq.URL = tarUrl
	tarReq.Host = tarUrl.Host

	if clientIP, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
		if prior, ok := tarReq.Header["X-Forwarded-For"]; ok {
			clientIP = strings.Join(prior, ", ") + ", " + clientIP
		}
		tarReq.Header.Set("X-Forwarded-For", clientIP)
	}

	res, err := transport.RoundTrip(tarReq)
	if err != nil {
		w.WriteHeader(502)
		return
	}
	defer res.Body.Close()

	for key, value := range res.Header {
		for _, v := range value {
			w.Header().Set(key, v)
		}
	}

	w.WriteHeader(res.StatusCode)
	io.Copy(w, res.Body)
	return

}

func main() {

	log.Println("Start http proxy server on :8080 ...")
	err := http.ListenAndServe(":8080", &Mux{})
	if err != nil {
		panic(err)
	}
}
