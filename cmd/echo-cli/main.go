package main

import (
	"context"
	"fmt"
	"time"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/hpifu/tpl-go-thrift/api/echoservice"
)

func main() {
	var transport thrift.TTransport
	var err error
	transport, err = thrift.NewTSocket("localhost:3001")
	if err != nil {
		panic(err)
	}
	transport, err = thrift.NewTBufferedTransportFactory(8192).GetTransport(transport)
	if err != nil {
		panic(err)
	}
	defer transport.Close()

	if err := transport.Open(); err != nil {
		panic(err)
	}

	protocolFactory := thrift.NewTCompactProtocolFactory()
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	client := echoservice.NewEchoServiceClient(thrift.NewTStandardClient(iprot, oprot))

	req := &echoservice.EchoReq{Rid: "1234567", Message: "hello world"}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(50*time.Millisecond))
	defer cancel()
	res, err := client.Echo(ctx, req)
	fmt.Println(res, err)
}
