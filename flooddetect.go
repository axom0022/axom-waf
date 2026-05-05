package main

import (
	"log"
	"sync"
	"time"
)

type floodrecord struct {
	count int
	start time.Time
}

var synfloodmap sync.Map
var udpfloodmap sync.Map
var icmpfloodmap sync.Map
var httpfloodmap sync.Map
var slowlorismap sync.Map

func startflooddetector() {
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for range ticker.C {
			checksynflood()
			checkudpflood()
			checkicmpflood()
			checkhttpflood()
			checkslowloris()
		}
	}()
}

func checksynflood() {
	synfloodmap.Range(func(key, value interface{}) bool {
		ip := key.(string)
		rec := value.(*floodrecord)
		if rec.count > config.synfloodthreshold {
			log.Printf("syn flood attacker ip: %s with %d syn packets per second", ip, rec.count)
			addiptablesblock(ip)
			synfloodmap.Delete(ip)
		} else {
			rec.count = 0
			rec.start = time.Now()
		}
		return true
	})
}

func checkudpflood() {
	udpfloodmap.Range(func(key, value interface{}) bool {
		ip := key.(string)
		rec := value.(*floodrecord)
		if rec.count > config.udpfloodthreshold {
			log.Printf("udp flood attacker ip: %s with %d udp packets per second", ip, rec.count)
			addiptablesblock(ip)
			udpfloodmap.Delete(ip)
		} else {
			rec.count = 0
			rec.start = time.Now()
		}
		return true
	})
}

func checkicmpflood() {
	icmpfloodmap.Range(func(key, value interface{}) bool {
		ip := key.(string)
		rec := value.(*floodrecord)
		if rec.count > config.icmpfloodthreshold {
			log.Printf("icmp flood attacker ip: %s with %d icmp packets per second", ip, rec.count)
			addiptablesblock(ip)
			icmpfloodmap.Delete(ip)
		} else {
			rec.count = 0
			rec.start = time.Now()
		}
		return true
	})
}

func checkhttpflood() {
	httpfloodmap.Range(func(key, value interface{}) bool {
		ip := key.(string)
		rec := value.(*floodrecord)
		if rec.count > config.httpfloodthreshold {
			log.Printf("http flood attacker ip: %s with %d requests per second", ip, rec.count)
			addiptablesblock(ip)
			httpfloodmap.Delete(ip)
		} else {
			rec.count = 0
			rec.start = time.Now()
		}
		return true
	})
}

func checkslowloris() {
	slowlorismap.Range(func(key, value interface{}) bool {
		ip := key.(string)
		rec := value.(*floodrecord)
		if rec.count > config.slowloristhreshold {
			log.Printf("slowloris attacker ip: %s with %d partial connections", ip, rec.count)
			addiptablesblock(ip)
			slowlorismap.Delete(ip)
		} else {
			rec.count = 0
			rec.start = time.Now()
		}
		return true
	})
}

func recordsyn(ip string) {
	recordflood(ip, &synfloodmap)
}

func recorducp(ip string) {
	recordflood(ip, &udpfloodmap)
}

func recordicmp(ip string) {
	recordflood(ip, &icmpfloodmap)
}

func recordhttp(ip string) {
	recordflood(ip, &httpfloodmap)
}

func recordslowloris(ip string) {
	recordflood(ip, &slowlorismap)
}

func recordflood(ip string, m *sync.Map) {
	val, _ := m.LoadOrStore(ip, &floodrecord{count: 0, start: time.Now()})
	rec := val.(*floodrecord)
	if time.Since(rec.start) < 1*time.Second {
		rec.count++
	} else {
		rec.count = 1
		rec.start = time.Now()
	}
}
