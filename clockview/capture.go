package clockskew

import (
	"encoding/binary"
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
		hostIp := srcIP

		for _, opt := range tcp.Options {
			if opt.OptionType.String() != "Timestamps" {
				continue
			}

			hostTs := binary.BigEndian.Uint32(opt.OptionData[:4])
            localTs := uint32(time.Now().Unix())


			cs := Clock{
				HostIp  : hostIp,
				LocalTs : localTs,
				HostTs  : hostTs,
		    }

			ClockChannel <- cs
		}
	}
}
