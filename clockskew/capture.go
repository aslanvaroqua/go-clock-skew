package clockskew

import (
	"encoding/binary"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"time"
)

func CapturePacket() {

	device := DeviceName

	handle, _ := pcap.OpenLive(
		device,
		int32(65535),
		false,
		-1*time.Second,
	)

	defer handle.Close()

	bpFilter := BpConfig

	handle.SetBPFFilter(bpFilter)

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	//startClock := monotime.Now()

	for packet := range packetSource.Packets() {

		ipLayer := packet.Layer(layers.LayerTypeIPv4)
		if ipLayer == nil {
			continue
		}
		ip, _ := ipLayer.(*layers.IPv4)
		srcIP := ip.SrcIP.String()

		tcpLayer := packet.Layer(layers.LayerTypeTCP)
		if tcpLayer == nil {
			continue
		}

		tcp, _ := tcpLayer.(*layers.TCP)
		remoteIp := srcIP

		for _, opt := range tcp.Options {
			if opt.OptionType.String() != "Timestamps" {
				continue
			}

			remoteTs := binary.BigEndian.Uint32(opt.OptionData[:4])
            delta = 6000
            localTs = time.Now().Unix()
            skew = math.Abs(i.Now().Unix() - srcTs)

			cs := ClockSkew{
				LocalTs   : localTs,
				RemoteIp  : remoteIp,
				RemoteTs: : remoteTs,
				Skew : skew
		    }

			ClockSkewChannel <- cs
		}
	}
}
