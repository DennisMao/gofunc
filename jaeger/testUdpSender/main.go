package main

import (
	"io"
	"time"

	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
)

func NewJaegerTracer(serviceName string) (opentracing.Tracer, io.Closer) {
	// jaeger服务端接收compact thrift protocol地址 默认6381
	sender, _ := jaeger.NewUDPTransport("192.168.33.10:6831", 0)

	// create Jaeger tracer
	tracer, closer := jaeger.NewTracer(
		serviceName,
		jaeger.NewConstSampler(true), // sample all traces
		jaeger.NewRemoteReporter(sender),
	)

	return tracer, closer
}

func main() {
	trace, closer := NewJaegerTracer("test_jaeger")
	defer closer.Close()

	//起始span
        //通常会从header里面然后解析出来
	yfSpan := trace.StartSpan("dennis")
	defer yfSpan.Finish()
	time.Sleep(100 * time.Millisecond)

	webServiceSpan := trace.StartSpan("webplatform", opentracing.ChildOf(yfSpan.Context()))
	defer webServiceSpan.Finish()
	time.Sleep(200 * time.Millisecond)

	accountSpan := trace.StartSpan("account", opentracing.ChildOf(webServiceSpan.Context()))
	defer accountSpan.Finish()
	time.Sleep(300 * time.Millisecond)

	authSpan := trace.StartSpan("auth", opentracing.ChildOf(accountSpan.Context()))
	defer authSpan.Finish()
	time.Sleep(400 * time.Millisecond)

}

