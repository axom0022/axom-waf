package main

import (
	"net"
	"sync"
	"syscall"
)

func startrawdetector() {
	go func() {
		fd, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, int(htons(syscall.ETH_P_IP)))
		if err != nil {
			return
		}
		buf := make([]byte, 2048)
		for {
			n, _, err := syscall.Recvfrom(fd, buf, 0)
			if err != nil || n < 20 {
				continue
			}
			parsepacket(buf[:n])
		}
	}()
}

func htons(i uint16) uint16 {
	return (i<<8)&0xff00 | i>>8
}

func parsepacket(pkt []byte) {
	if pkt[0]>>4 != 4 {
		return
	}
	protocol := pkt[9]
	srcip := net.IP(pkt[12:16]).String()
	if protocol == 6 {
		if len(pkt) > 40 {
			tcpflags := pkt[47]
			if tcpflags&0x02 != 0 {
				recordsyn(srcip)
			}
		}
	} else if protocol == 17 {
		recorducp(srcip)
	} else if protocol == 1 {
		recordicmp(srcip)
	}
}
