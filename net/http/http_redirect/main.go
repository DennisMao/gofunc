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

	// 重定向
	// 标准状态码 201 301-307
	// 常用状态码 301 302 303
	//201   Created 创建成功,手动跳转
	//301	Moved Permanently	永久移动。请求的资源已被永久的移动到新URI，返回信息会包括新的URI，浏览器会自动定向到新URI。今后任何新的请求都应使用新的URI代替
	//302	Found	临时移动。与301类似。但资源只是临时被移动。客户端应继续使用原有URI
	//303	See Other	查看其它地址。与301类似。使用GET和POST请求查看
	//304	Not Modified	未修改。所请求的资源未修改，服务器返回此状态码时，不会返回任何资源。客户端通常会缓存访问过的资源，通过提供一个头信息指出客户端希望只返回在指定日期之后修改的资源
	//305	Use Proxy	使用代理。所请求的资源必须通过代理访问
	//307	Temporary Redirect	临时重定向
	http.Redirect(w, r, tarUrlStr, 301)

	return

}

func main() {

	log.Println("Start http proxy server on :8080 ...")
	err := http.ListenAndServe(":8080", &Mux{})
	if err != nil {
		panic(err)
	}
}
