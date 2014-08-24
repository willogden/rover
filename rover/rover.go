package rover

import (
    //"log"
)

type Rover struct {

    // Inbound messages to rover
    receivedMessages chan Messager

    // Outbound messages from rover
    sendMessage chan Messager

}

// Create new instance of a rover.
// Passed a from and to channel for messages
func NewRover(receivedMessages chan Messager, sendMessage chan Messager) *Rover {

    rover := &Rover{
        receivedMessages: receivedMessages,
        sendMessage: sendMessage,
    }

    return rover
}

// Process / Action a message sent to Rover
func (r *Rover) processReceivedMessage(message Messager) {

    switch message.(type) {
        case *LocationMessage:
            r.processLocationMessage(message.(*LocationMessage))
    }

}

// Goroutine to listen for messages sent to Rover
func (r *Rover) listenForMessages() {
    for {
        for message := range r.receivedMessages {
            r.processReceivedMessage(message);
        }
    }
}


func (r *Rover) createMessages() {
    for {
        r.sendMessage <- &StatusMessage{Status:"message from rover"}
    }
}


func (r *Rover) processLocationMessage(lm *LocationMessage) {
    sm := StatusMessage{Status:"location message received"}
    r.sendMessage <- sm
}

// Start the Rover
func (r *Rover) Run() {
    go r.listenForMessages()
    //go r.createMessages()
}
