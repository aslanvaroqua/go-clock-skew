package clockskew

type Clock struct {
	LocalTs   uint32  `json:localTs`
	HostIp    string `json:hostIp`
	HostTs    uint32  `json:hostTs`
}

var DeviceName   string
var BpConfig     string
var ClockChannel chan Clock

