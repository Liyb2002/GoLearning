package main

import (
        "fmt"
        "log"
        "net/rpc"
)

type Args struct {
        A, B int
}

type Quotient struct {
        Quo, Rem int
}

type Arith int

func main() {
        client, err := rpc.DialHTTP("tcp", ":1234")
        if err != nil {
                log.Fatal("dialing:", err)
        }
        // Synchronous call
        args := &Args{3, 5}
        var reply int
        err = client.Call("Arith.Multiply", args, &reply)
        if err != nil {
                log.Fatal("arith error:", err)
        }
        fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

        args = &Args{100, 23}
        var quot Quotient
        err = client.Call("Arith.Divide", args, &quot)
        if err != nil {
                log.Fatal("divide error:", err)
        }
        fmt.Printf("Arith: %d/%d = q%d r%d\n", args.A, args.B, quot.Quo, quot.Rem)

}