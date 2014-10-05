package main

import (
    "github.com/codegangsta/negroni"
    "github.com/gorilla/websocket"
    "net/http"
    "fmt"
    "log"
    "os"
    "os/signal"
    "syscall"

    "github.com/willogden/rover/rover"
    "github.com/willogden/rover/rover/platform"

)

func main() {

    upgrader := websocket.Upgrader{
        ReadBufferSize:  1024,
        WriteBufferSize: 1024,
    }

    broker := rover.NewBroker()
    broker.Run()

    r := platform.NewRover(broker.GetToRoverChannel(),broker.GetFromRoverChannel())
    r.Run()

    // Handle SIGINT and SIGTERM.
    ch := make(chan os.Signal)
    signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        <-ch
        // Do cleanup here
        r.Stop()
        os.Exit(1)
    }()

    mux := http.NewServeMux()
    mux.HandleFunc("/s", func(w http.ResponseWriter, req *http.Request) {
        fmt.Fprintf(w, "Welcome to the home page!")
    })

    mux.HandleFunc("/api/ws",func (w http.ResponseWriter, r *http.Request) {
        conn, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
            log.Println(err)
            return
        }

        rover.NewConnection(conn,broker)
    })

    n := negroni.Classic()
    n.UseHandler(mux)
    n.Run(":3000")

}
