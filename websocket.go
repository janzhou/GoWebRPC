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
    var id int

    c := jsonrpc.NewClient(ws)

    err := c.Call("User.Getid", nil, &id)
    if err != nil {
        log.Print("User.Getid error:", err)
        return
    } else {
        user.client[id] = c;
        user.mutex[id].Lock()
    }

}

