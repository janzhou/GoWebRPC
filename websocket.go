package main

import (
    "code.google.com/p/go.net/websocket"
    "net/rpc/jsonrpc"
    "log"
)

func jsonrpcHandler(ws *websocket.Conn) {
    jsonrpc.ServeConn(ws)
}

func pushHandler(ws *websocket.Conn) {
    args := &Args{7, 8}
    var reply int

    c := jsonrpc.NewClient(ws)

    err := c.Call("Arith.Multiply", args, &reply)
    if err != nil {
        log.Fatal("arith error:", err)
    }
    log.Printf("Arith: %d*%d=%d", args.A, args.B, reply)
}

