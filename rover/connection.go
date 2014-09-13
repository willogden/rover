package rover

import (
    //"log"
    "encoding/json"
    "errors"

    "github.com/gorilla/websocket"
)


type Connection struct {
    // The websocket connection.
    ws *websocket.Conn

    // The message broker
    broker *Broker

    // Buffered channel of outbound messages.
    toClient chan Messager
}


type InboundWebSocketMessage struct {
    Type string `json:"type"`
    Data *json.RawMessage `json:"data"`
}


type OutboundWebSocketMessage struct {
    Type string `json:"type"`
    Data Messager `json:"data"`
}


func NewConnection(ws *websocket.Conn, b *Broker) *Connection {

    c := &Connection{ws: ws, broker: b, toClient: make(chan Messager)}
    c.broker.register <- c

    go c.writer()
    go c.reader()

    return c
}


func (c *Connection) reader() {
    defer func() {
        c.broker.unregister <- c
        c.ws.Close()
    }()

    for {

        var iwm InboundWebSocketMessage
        if err := c.ws.ReadJSON(&iwm); err != nil {
            break
        }

        if m,err := c.UnmarshalWebSocketMessage(&iwm); err != nil {
            break
        } else {
            c.broker.toRover <- m
        }

    }

}


func (c * Connection) UnmarshalWebSocketMessage(iwm *InboundWebSocketMessage) (Messager,error) {

    switch {
        case iwm.Type == "location":
            var lm LocationMessage
            if err := json.Unmarshal(*iwm.Data, &lm); err != nil {
                return nil, err
            }
            return &lm,nil
        case iwm.Type == "motorspeed":
            var msm MotorSpeedMessage
            if err := json.Unmarshal(*iwm.Data, &msm); err != nil {
                return nil, err
            }
            return &msm,nil
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

        err := c.ws.WriteJSON(owm)
        if err != nil {
            break
        }
    }
}
