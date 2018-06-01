package ethrpc

import (
	"bytes"

	"log"
	"net/http"

	"../gorilla/rpc/json"
	//"./rpcexample"
)

func JRpcCallNoParams() {

	log.Output(1, "JRpcCallNoParams")
}

func JRpcCall(valfunc string) (string, error) {
	url := "http://127.0.0.1:8545/" //"http://localhost:1234/rpc"
	/*
		args := &Args{
			A: 2,
			B: 3,
		}
	*/
	//log.Output(1, "web3_clientVersion1")
	var FuncName string
	FuncName = valfunc //"eth_getBalance"
	var args []string
	if FuncName == "web3_clientVersion" || FuncName == "eth_coinbase" || FuncName == "eth_protocolVersion" {
		args = []string{}
	}
	if FuncName == "eth_getBalance" {
		args = []string{"0x407d73d8a49eeb85d32cf465507dd71d507100c1", "latest"}
	}

	message, err := json.EncodeClientRequest(FuncName, args)
	if err != nil {
		log.Fatalf("%s", err)
	}
	//log.Output(1, "web3_clientVersion2")

	//message2 := "{\"jsonrpc\":\"2.0\",\"method\":\"web3_clientVersion\",\"params\":[],\"id\":67}"
	//message2 := "{\"jsonrpc\":\"2.0\",\"method\":\"eth_coinbase\",\"params\":[],\"id\":64}"
	//message2 := "{\"jsonrpc\":\"2.0\",\"method\":\"eth_protocolVersion\",\"params\":[],\"id\":67}"
	//message2 := "{\"jsonrpc\":\"2.0\",\"method\":\"eth_getBalance\",\"params\":[\"0x407d73d8a49eeb85d32cf465507dd71d507100c1\", \"latest\"],\"id\":1}"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(message)))
	if err != nil {
		log.Fatalf("%s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error in sending request to %s. %s", url, err)
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	s := buf.String()
	log.Output(1, s)
	//log.Output(1, string("123"))
	//var result Result
	//err = json.DecodeClientResponse(resp.Body, &result)
	//if err != nil {
	//	log.Fatalf("Couldn't decode response. %s", err)
	//}
	//log.Printf("%d*%d=%d\n", args.A, args.B, result)

	return s, err
}
