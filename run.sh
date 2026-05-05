#!/bin/bash
ulimit -n 1048576
sysctl -w net.core.somaxconn=65535
sysctl -w net.ipv4.tcp_syncookies=1
sysctl -w net.ipv4.tcp_fin_timeout=10
sysctl -w net.ipv4.tcp_tw_reuse=1
iptables -F
go build -o axomwaf .
sudo ./axomwaf
