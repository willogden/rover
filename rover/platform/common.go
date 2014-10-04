package platform

import (
    "github.com/willogden/rover/rover/messages"
)

type RoverPlatformer interface {

    HandleReceivedMessage(message messages.Messager)
    ListenForMessages()
    Run()
    Stop()
    HandleLocationMessage(lm *messages.LocationMessage)
    HandleMotorSpeedMessage(msm *messages.MotorSpeedMessage)

}


type Common struct {

    // Inbound messages to rover
    receivedMessages chan messages.Messager

    // Outbound messages from rover
    sendMessage chan messages.Messager

    rp RoverPlatformer

}

// Process / Action a message sent to Rover
func (c *Common) HandleReceivedMessage(message messages.Messager) {

    switch message.(type) {
        case *messages.LocationMessage:
            c.rp.HandleLocationMessage(message.(*messages.LocationMessage))
        case *messages.MotorSpeedMessage:
            c.rp.HandleMotorSpeedMessage(message.(*messages.MotorSpeedMessage))
    }

}

// Listen for messages sent to Rover
func (c *Common) ListenForMessages() {

    defer func() {
        c.rp.Stop()
    }()

    for {
        for message := range c.receivedMessages {
            c.HandleReceivedMessage(message);
        }
    }
}

// Start the Rover
func (c *Common) Run() {

    go c.ListenForMessages()

}

// Stop the Rover
func (c *Common) Stop() {

}
