package main

import (
    "fmt"
    "log"
    "net"
    "bufio"
)

func main() {
    listen, err := net.Listen("tcp", ":8080")
    if err != nil {
        panic(err)
    }

    for {
        conn, err := listen.Accept()
        if err != nil {
            log.Printf("error: %s\n", err)
        }
        go handle(conn)
    }
}

func handle(conn net.Conn) {
    msg := make(chan string, 1)
    defer conn.Close()
    defer close(msg)
    go send(conn, msg)
    receive(conn, msg)
}

func send(conn net.Conn, message <-chan string) {
    for msg := range message {
        fmt.Println(conn, msg)
    }
}

func receive(conn net.Conn, message chan string) {
    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        fmt.Println("message: %s\n", scanner.Text())
        message <- scanner.Text()
    }
}

