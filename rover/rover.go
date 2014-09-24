package rover

import (
    //"log"

    //"github.com/kidoman/embd"
    _ "github.com/kidoman/embd/host/rpi"
    "github.com/kidoman/embd/controller/servoblaster"
)

type Rover struct {

    // Inbound messages to rover
    receivedMessages chan Messager

    // Outbound messages from rover
    sendMessage chan Messager

    // Servoblaster (software PWM)
    sb *servoblaster.ServoBlaster

}

// Create new instance of a rover.
// Passed a from and to channel for messages
func NewRover(receivedMessages chan Messager, sendMessage chan Messager) *Rover {

    rover := &Rover{
        receivedMessages: receivedMessages,
        sendMessage: sendMessage,
        sb: servoblaster.New(),
    }

    return rover
}

// Process / Action a message sent to Rover
func (r *Rover) processReceivedMessage(message Messager) {

    switch message.(type) {
        case *LocationMessage:
            r.processLocationMessage(message.(*LocationMessage))
        case *MotorSpeedMessage:
            r.processMotorSpeedMessage(message.(*MotorSpeedMessage))
    }

}

// Goroutine to listen for messages sent to Rover
func (r *Rover) listenForMessages() {

    defer func() {
        _ = r.sb.Close()
    }()

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
    sm := NewStatusMessage()
    sm.Status = "location message received"

    r.sendMessage <- sm
}


func (r *Rover) processMotorSpeedMessage(msm *MotorSpeedMessage) {

    r.sb.Channel(msm.Motor).SetMicroseconds((20000/100)*msm.Speed)

    sm := NewStatusMessage()
    sm.Status = "motorspeed message received"

    r.sendMessage <- sm
}

// Start the Rover
func (r *Rover) Run() {

    go r.listenForMessages()

    //go r.createMessages()
}
