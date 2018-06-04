package clockskew

type ClockSkew struct {
	LocalTs   int64  `json:localTs`
	HostIp    string `json:hostIp`
	HostTs    int64  `json:hostTs`
	Skew      int64  `json:skew`
}

var DeviceName   string
var BpConfig     string
var ClockSkewChannel chan ClockSkew

