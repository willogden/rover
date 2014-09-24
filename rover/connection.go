package rover

import (
    "log"
    "encoding/json"
    "errors"
    "sync"

    "github.com/gorilla/websocket"
)


type Connection struct {
    // The websocket connection.
    ws *websocket.Conn

    // The message broker
    broker *Broker

    // Buffered channel of outbound messages.
    toClient chan Messager

    mut *sync.Mutex
}


type InboundWebSocketMessage struct {
    Type string `json:"type"`
    Data *json.RawMessage `json:"data"`
}


type OutboundWebSocketMessage struct {
    Type string `json:"type"`
    Data Messager `json:"data"`
}


func NewConnection(ws *websocket.Conn, b *Broker) {

    c := &Connection{ws: ws, broker: b, toClient: make(chan Messager), mut: &sync.Mutex{}}
    c.broker.register <- c

    go c.writer()
    c.reader()
}


func (c *Connection) reader() {
    defer func() {
        c.broker.unregister <- c
        c.ws.Close()
    }()

    for {

        var iwm InboundWebSocketMessage

        if err := c.ws.ReadJSON(&iwm); err != nil {
            log.Println(err.Error())
            return
        }

        if m,err := c.UnmarshalWebSocketMessage(&iwm); err != nil {
            log.Println(err.Error())
            return
        } else {
            c.broker.toRover <- m
        }

    }

}


func (c * Connection) UnmarshalWebSocketMessage(iwm *InboundWebSocketMessage) (Messager,error) {

    if m := NewMessageByType(iwm.Type); m != nil {
        if err := json.Unmarshal(*iwm.Data, &m); err != nil {
            return nil, err
        }
        return m,nil
    }

    return nil,errors.New("InboundWebSocketMessage type not recognised")
}


func (c *Connection) writer() {
    defer func() {
        c.broker.unregister <- c
        c.ws.Close()
    }()

    for message := range c.toClient {

        owm := &OutboundWebSocketMessage{Type: message.GetType(), Data: message}

        if err := c.ws.WriteJSON(owm); err != nil {
            log.Println(err.Error())
            return
        }
    }
}
