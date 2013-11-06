package main

import (
    "code.google.com/p/go.net/websocket"
    "net/http"
    "net/rpc"
    "flag"
)
func main() {
    port := flag.String("port", ":8080", "http service address")
    htdocs := flag.String("htdocs", "htdocs", "http dir")
    flag.Parse()

    go h.run()
    rpc.Register(new(Arith))

    http.Handle("/jsonrpc", websocket.Handler(jsonrpcHandler))
    http.Handle("/notify", websocket.Handler(notifyHandler))
    http.Handle("/push", websocket.Handler(pushHandler))
    http.Handle("/", http.FileServer(http.Dir(*htdocs)))
    err := http.ListenAndServe(*port, nil)
    if err != nil {
        panic("ListenAndServe: " + err.Error())
    }

}
