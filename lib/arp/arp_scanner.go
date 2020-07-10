package arp

import (
	"Pallas/lib/ip"
	"net"
	"time"

	"github.com/mdlayher/arp"
)

type ARPDetectInfo struct {
	IP  string
	Mac string
}

func ScanARP(ifaceName string, netSegment string) ([]ARPDetectInfo, error) {

	var result = make([]ARPDetectInfo, 0)

	ifIndex, err := net.InterfaceByName(ifaceName)
	if err != nil {
		return result, err
	}

	ips, e := ip.GetIPBySegment(netSegment)
	if e != nil {
		return result, e
	}

	conn, conErr := arp.Dial(ifIndex)
	if conErr != nil {
		return result, conErr
	}

	defer conn.Close()

	for _, ipString := range ips {

		if err := conn.SetDeadline(time.Now().Add(time.Second * 5)); err != nil {
			continue
		}

		targetIp := net.ParseIP(ipString).To4()
		hwAddr, _ := conn.Resolve(targetIp)

		result = append(result, ARPDetectInfo{
			IP:  targetIp.String(),
			Mac: hwAddr.String(),
		})

	}
	return result, nil
}
