package main

import (
    "code.google.com/p/go.net/websocket"
    "net/http"
    "net/rpc"
    "net/rpc/jsonrpc"
    "io"
    "log"
    "errors"
    "flag"
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

func main() {
    port := flag.String("port", ":8080", "http service address")
    flag.Parse()

    go h.run()
    rpc.Register(new(Arith))

    http.Handle("/rpc", websocket.Handler(rpcHandler))
    http.Handle("/rpcclient", websocket.Handler(rpcclientHandler))
    http.Handle("/notify", websocket.Handler(notifyHandler))
    http.Handle("/echo", websocket.Handler(echoHandler))
    http.Handle("/", http.FileServer(http.Dir(".")))
    err := http.ListenAndServe(*port, nil)
    if err != nil {
        panic("ListenAndServe: " + err.Error())
    }

}