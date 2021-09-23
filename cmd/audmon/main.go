package main

import (
	"github.com/lowellmower/audmon/pkg/client"
    "github.com/lowellmower/audmon/pkg/daemon"
    "github.com/lowellmower/audmon/pkg/server"
)

func main() {
    client.Printer()
    daemon.Printer()
    server.Printer()
}
