package client

import (
	"fmt"
    "github.com/gorilla/websocket"
    "log"
    "net/url"
    "os"
    "os/signal"
    "time"
)

func Printer() {
    fmt.Println("client")
}

func Run(addr string) error {
    interrupt := make(chan os.Signal, 1)
    signal.Notify(interrupt, os.Interrupt)

    u := url.URL{Scheme: "ws", Host: addr, Path: "/echo"}
    log.Printf("connecting to %s", u.String())

    c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
    if err != nil {
        log.Fatal("dial:", err)
    }
    defer c.Close()

    done := make(chan struct{})

    go func() {
        defer close(done)
        for {
            _, message, err := c.ReadMessage()
            if err != nil {
                log.Println("read:", err)
                return
            }
            log.Printf("recv: %s", message)
        }
    }()

    ticker := time.NewTicker(time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-done:
            return nil
        case t := <-ticker.C:
            err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
            if err != nil {
                log.Println("write:", err)
                return nil
            }
        case <-interrupt:
            log.Println("interrupt")

            // Cleanly close the connection by sending a close message and then
            // waiting (with timeout) for the server to close the connection.
            err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
            if err != nil {
                log.Println("write close:", err)
                return err
            }
            select {
            case <-done:
            case <-time.After(time.Second):
            }
            return nil
        }
    }
}