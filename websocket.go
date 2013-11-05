package main

import (
    "code.google.com/p/go.net/websocket"
    "net/rpc/jsonrpc"
    "io"
    "log"
    "errors"
)

type Args struct {
    A int
    B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
    *reply = args.A * args.B
    return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func rpcHandler(ws *websocket.Conn) {
    jsonrpc.ServeConn(ws)
}

func rpcclientHandler(ws *websocket.Conn) {
    args := &Args{7, 8}
    var reply int

    c := jsonrpc.NewClient(ws)

    err := c.Call("Arith.Multiply", args, &reply)
    if err != nil {
        log.Fatal("arith error:", err)
    }
    log.Printf("Arith: %d*%d=%d", args.A, args.B, reply)
}

func echoHandler(ws *websocket.Conn) {
    io.Copy(ws, ws)
}
