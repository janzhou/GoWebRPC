package main

import (
    "code.google.com/p/go.net/websocket"
    "net/http"
    "net/rpc"
    "net/rpc/jsonrpc"
    "io"
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

func echoHandler(ws *websocket.Conn) {
        io.Copy(ws, ws)
}

func main() {
    rpc.Register(new(Arith))

    http.Handle("/rpc", websocket.Handler(rpcHandler))
    http.Handle("/echo", websocket.Handler(echoHandler))
    http.Handle("/", http.FileServer(http.Dir(".")))
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        panic("ListenAndServe: " + err.Error())
    }

}
