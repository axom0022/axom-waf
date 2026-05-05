package main

import (
	"os/exec"
)

func applysysctl() {
	exec.Command("sysctl", "-w", "net.core.somaxconn=65535").Run()
	exec.Command("sysctl", "-w", "net.ipv4.tcp_syncookies=1").Run()
	exec.Command("sysctl", "-w", "net.ipv4.tcp_fin_timeout=10").Run()
	exec.Command("sysctl", "-w", "net.ipv4.tcp_tw_reuse=1").Run()
	exec.Command("sysctl", "-w", "net.core.rmem_max=134217728").Run()
	exec.Command("sysctl", "-w", "net.core.wmem_max=134217728").Run()
	exec.Command("sysctl", "-w", "net.ipv4.tcp_rmem=4096 87380 134217728").Run()
	exec.Command("sysctl", "-w", "net.ipv4.tcp_wmem=4096 65536 134217728").Run()
	exec.Command("sysctl", "-w", "net.core.netdev_max_backlog=5000").Run()
}
