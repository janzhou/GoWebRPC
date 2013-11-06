package main

import (
    "net/rpc"
    "sync"
)

type User struct {
    id int
    client [32]*rpc.Client
    mutex  [32]sync.Mutex
}

func (u *User) Signup (args *string, id *int) error {
    *id = u.id

    if *id == 32 {
        *id = 0
        u.id = 0
    } else {
        u.id++
    }

    u.mutex[*id].Lock()
    return nil
}

var user = new(User)
