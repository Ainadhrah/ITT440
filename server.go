package main

import (
    "fmt"
    "net"
    "os"
    "time"
)

func main() {
    service := ":5678"
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)
    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)
    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        go handleClient(conn)
    }
}

func handleClient(conn net.Conn) {
    fmt.Println("Welcome to AINA'S profile! ")
    defer conn.Close()
    daytime := time.Now().String()
    conn.Write([]byte(daytime)
    // we're finished with this client
}
func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1);
    }
}
