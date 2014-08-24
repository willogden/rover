package rover


type Messager interface {
    GetType() string
}


type Message struct {

}

func (m Message) GetType() string {
    return "message"
}

type LocationMessage struct {
    Lon float64 `json:"lon"`
    Lat float64 `json:"lat"`
}


func (lm LocationMessage) GetType() string {
    return "location"
}


type MotorSpeedMessage struct {
    Motor int `json:"motor"`
    Speed int `json:"speed"`
}

func (msm MotorSpeedMessage) GetType() string {
    return "motorspeed"
}


type StatusMessage struct {
    Status string `json:"status"`
}

func (sm StatusMessage) GetType() string {
    return "status"
}
