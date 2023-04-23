package main

import (
    "fmt"
    "net/url"
)

func main() {
    // uri := "https://grootpiano.net"
    uri := "grootpiano.net"

    u, err :=  url.Parse(uri)
    if err != nil {
        panic(err)
    }

    // do not contain port
    hostname := u.Hostname()

    // can not parse url without schema
    fmt.Println(u.IsAbs())
    fmt.Println(hostname)
}
