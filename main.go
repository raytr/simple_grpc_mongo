package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/kaysush/protobuf-demo/proto/echo" //<-- Take a note that I've created my code folder in $GOPATH/src
)

func main() {
	req := &echo.EchoRequest{Name: "Sushil"}
	data, err := proto.Marshal(req)
	if err != nil {
		log.Fatalf("Error while marshalling the object : %v", err)
	}

	res := &echo.EchoRequest{}
	err = proto.Unmarshal(data, res)
	if err != nil {
		log.Fatalf("Error while un-marshalling the object : %v", err)
	}
	fmt.Printf("Value from un-marshalled data is %v", res.GetName())

}
