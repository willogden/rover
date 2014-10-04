package platform

import (
    "github.com/willogden/rover/rover/messages"
)

type Rover struct {
    *Common
}

// Create new instance of a rover.
// Passed a from and to channel for messages
func NewRover(receivedMessages chan messages.Messager, sendMessage chan messages.Messager) *Rover {

    rover := &Rover{}

    common := &Common{
        receivedMessages: receivedMessages,
        sendMessage: sendMessage,
        rp: rover,
    }

    rover.Common = common


    return rover
}

func (r *Rover) HandleLocationMessage(lm *messages.LocationMessage) {

    sm := messages.NewStatusMessage()
    sm.Status = "location message received"

    r.sendMessage <- sm

}

func (r *Rover) HandleMotorSpeedMessage(msm *messages.MotorSpeedMessage) {

    sm := messages.NewStatusMessage()
    sm.Status = "motorspeed message received"

    r.sendMessage <- sm

}
