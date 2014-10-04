package messages

import (

)

type Messager interface {
    GetType() string
}


type Message struct {
    messageType string
}

func (m Message) GetType() string {
    return m.messageType
}

const (
    LocationMessageType = "location"
    MotorSpeedMessageType = "motorspeed"
    StatusMessageType = "status"
)

func NewMessageByType(messageType string) Messager {
    switch messageType {
        case LocationMessageType:
            return NewLocationMessage()
        case MotorSpeedMessageType:
            return NewMotorSpeedMessage()
        case StatusMessageType:
            return NewStatusMessage()
        default:
            return nil
    }
}


type LocationMessage struct {
    Message
    Lon float64 `json:"lon"`
    Lat float64 `json:"lat"`
}

func NewLocationMessage() *LocationMessage {
    return &LocationMessage{Message: Message{messageType: LocationMessageType}}
}


type MotorSpeedMessage struct {
    Message
    Motor int `json:"motor"`
    Speed int `json:"speed"`
}

func NewMotorSpeedMessage() *MotorSpeedMessage {
    return &MotorSpeedMessage{Message: Message{messageType: MotorSpeedMessageType}}
}


type StatusMessage struct {
    Message
    Status string `json:"status"`
}

func NewStatusMessage() *StatusMessage {
    return &StatusMessage{Message: Message{messageType: StatusMessageType}}
}
