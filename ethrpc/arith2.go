package ethrpc

import (
	"net/http"
	//"ethrpc"
)

//Represents service Arith on JSON-RPC
type Arith2 int

//Invoked by JSON-RPC client and calls rpcexample.Multiply which stores product of args.A and args.B in result
func (t *Arith2) Multiply(r *http.Request, args *Args, result *Result) error {
	return Multiply(*args, result)
}
