package main

import (
	"encoding/json"
	"os"
)

type configstruct struct {
	listenport          int      `json:"listenport"`
	upstreamtarget      string   `json:"upstreamtarget"`
	ratepersec          int      `json:"ratepersec"`
	rateburst           int      `json:"rateburst"`
	maxconnsperip       int      `json:"maxconnsperip"`
	iptablesblocksec    int      `json:"iptablesblocksec"`
	synfloodthreshold   int      `json:"synfloodthreshold"`
	udpfloodthreshold   int      `json:"udpfloodthreshold"`
	icmpfloodthreshold  int      `json:"icmpfloodthreshold"`
	slowloristhreshold  int      `json:"slowloristhreshold"`
	httpfloodthreshold  int      `json:"httpfloodthreshold"`
	whitelist           []string `json:"whitelist"`
	blacklist           []string `json:"blacklist"`
	sqlpatterns         []string `json:"sqlpatterns"`
	xsspatterns         []string `json:"xsspatterns"`
}

var config configstruct

func loadconfig() {
	data, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
}
