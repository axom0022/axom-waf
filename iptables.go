package main

import (
	"os/exec"
	"strconv"
	"sync"
)

var blockedips sync.Map

func startiptables() {
	exec.Command("iptables", "-A", "INPUT", "-p", "tcp", "--dport", strconv.Itoa(config.listenport), "-m", "connlimit", "--connlimit-above", strconv.Itoa(config.maxconnsperip), "--connlimit-mask", "32", "-j", "DROP").Run()
	exec.Command("iptables", "-A", "INPUT", "-p", "tcp", "--dport", strconv.Itoa(config.listenport), "-m", "limit", "--limit", "1000/s", "--limit-burst", "2000", "-j", "ACCEPT").Run()
	exec.Command("iptables", "-A", "INPUT", "-p", "tcp", "--tcp-flags", "SYN", "SYN", "-m", "limit", "--limit", "200/s", "--limit-burst", "400", "-j", "ACCEPT").Run()
	exec.Command("iptables", "-A", "INPUT", "-p", "tcp", "--tcp-flags", "SYN", "SYN", "-j", "DROP").Run()
	exec.Command("iptables", "-A", "INPUT", "-p", "udp", "-m", "limit", "--limit", "500/s", "--limit-burst", "1000", "-j", "ACCEPT").Run()
	exec.Command("iptables", "-A", "INPUT", "-p", "udp", "-j", "DROP").Run()
	exec.Command("iptables", "-A", "INPUT", "-p", "icmp", "-m", "limit", "--limit", "100/s", "--limit-burst", "200", "-j", "ACCEPT").Run()
	exec.Command("iptables", "-A", "INPUT", "-p", "icmp", "-j", "DROP").Run()
	exec.Command("iptables", "-A", "INPUT", "-p", "tcp", "--tcp-flags", "ALL", "NONE", "-j", "DROP").Run()
	exec.Command("iptables", "-A", "INPUT", "-p", "tcp", "--tcp-flags", "SYN,ACK", "SYN,ACK", "-m", "limit", "--limit", "100/s", "-j", "ACCEPT").Run()
	exec.Command("iptables", "-A", "INPUT", "-p", "tcp", "--tcp-flags", "SYN,ACK", "SYN,ACK", "-j", "DROP").Run()
	exec.Command("iptables", "-A", "INPUT", "-p", "tcp", "--tcp-flags", "FIN,ACK", "FIN,ACK", "-m", "limit", "--limit", "200/s", "-j", "ACCEPT").Run()
	exec.Command("iptables", "-A", "INPUT", "-p", "tcp", "--tcp-flags", "RST", "RST", "-m", "limit", "--limit", "200/s", "-j", "ACCEPT").Run()
	exec.Command("iptables", "-A", "INPUT", "-m", "state", "--state", "INVALID", "-j", "DROP").Run()
}

func addiptablesblock(ip string) {
	if _, loaded := blockedips.LoadOrStore(ip, true); loaded {
		return
	}
	exec.Command("iptables", "-A", "INPUT", "-s", ip, "-j", "DROP").Run()
	exec.Command("iptables", "-A", "FORWARD", "-s", ip, "-j", "DROP").Run()
}
