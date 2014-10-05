package platform

import (
    "os/exec"
    //"fmt"

    //"github.com/kidoman/embd"
    _ "github.com/kidoman/embd/host/rpi"
    "github.com/kidoman/embd/controller/servoblaster"

    "github.com/willogden/rover/rover/messages"
)

type Rover struct {
    *Common

    // Servoblaster (software PWM)
    sb *servoblaster.ServoBlaster
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

    rover.startServoBlaster()

    return rover
}

// Stop the Rover
func (r *Rover) Stop() {
    _ = r.sb.Close()

    r.stopServoBlaster()
}

// Start servoblaster service
func (r *Rover) startServoBlaster() {

    cmd := exec.Command("bash", "-c", "sudo ~/PiBits/ServoBlaster/user/servod --min=0 --max=2000")
    err := cmd.Run()
    if err != nil {
        panic(err)
    }

    r.sb = servoblaster.New()
}

func (r *Rover) stopServoBlaster() {

    cmd := exec.Command("bash", "-c", "sudo killall servod")
    err := cmd.Run()
    if err != nil {
        panic(err)
    }

}


func (r *Rover) HandleLocationMessage(lm *messages.LocationMessage) {

    sm := messages.NewStatusMessage()
    sm.Status = "location message received"

    r.sendMessage <- sm

}

func (r *Rover) HandleMotorSpeedMessage(msm *messages.MotorSpeedMessage) {

    r.sb.Channel(msm.Motor).SetMicroseconds((20000/100)*msm.Speed)

    sm := messages.NewStatusMessage()
    sm.Status = "motorspeed message received"

    r.sendMessage <- sm

}
