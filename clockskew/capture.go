package clockskew

import (
	"encoding/binary"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"time"
	"github.com/gavv/monotime"
	"math"
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
		srcPort := tcp.SrcPort
		taddr := srcIP + ":" + fmt.Sprintf("%d", srcPort)

		for _, opt := range tcp.Options {
			if opt.OptionType.String() != "Timestamps" {
				continue
			}

			srcTS := binary.BigEndian.Uint32(opt.OptionData[:4])

			cs := ClockSkew{
				Clock : int64(monotime.Now()),
				Taddr: taddr,
				SrcTS: int64(srcTS),
				Skew: float64(int64(monotime.Now()) - int64(srcTS)),

		}

			ClockSkewChannel <- cs
		}
	}
}
