package rover

import (
    "github.com/willogden/rover/rover/messages"
)

type Broker struct {

    // Registered connections.
    connections map[*Connection]bool

    // Inbound messages from the connections.
    toRover chan messages.Messager

    // Outbound messages to clients
    fromRover chan messages.Messager

    // Register requests from the connections.
    register chan *Connection

    // Unregister requests from connections.
    unregister chan *Connection

}

// Create a new Broker that acts as the middle man between websocket connections and Rover
func NewBroker() *Broker {

    broker := &Broker{
        toRover:   make(chan messages.Messager),
        fromRover:   make(chan messages.Messager),
        register:    make(chan *Connection),
        unregister:  make(chan *Connection),
        connections: make(map[*Connection]bool),
    }

    return broker
}

// Return the channel for sending messages to Rover
func (b *Broker) GetToRoverChannel() chan messages.Messager {
    return b.toRover
}

// Return the channel for receiving messages from Rover
func (b *Broker) GetFromRoverChannel() chan messages.Messager {
    return b.fromRover
}


// Start the Broker
func (b *Broker) Run() {

    go func() {
        for {
            select {
            case c := <-b.register:
                b.connections[c] = true
            case c := <-b.unregister:
                if _, ok := b.connections[c]; ok {
                    delete(b.connections, c)
                    close(c.toClient)
                }
            case m := <-b.fromRover:
                for c := range b.connections {
                    c.toClient <- m
                }
            }
        }
    }()
}
