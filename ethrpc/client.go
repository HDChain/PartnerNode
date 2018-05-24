package ethrpc

import (
	"log"
	"net/rpc"
	//"../rpcexample"
)

func main10() {
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatalf("Error in dialing. %s", err)
	}
	args := &Args{
		A: 2,
		B: 3,
	}
	var result Result
	err = client.Call("Arith.Multiply", args, &result)
	if err != nil {
		log.Fatalf("error in Arith", err)
	}
	log.Printf("%d*%d=%d\n", args.A, args.B, result)
}
