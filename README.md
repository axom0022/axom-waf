a person that uses ready DDoS tools is targetting your web/application/server/etc thinking he will ruin it? 

Hit him hard with axom-waf. 
A selfhosted defense against DDoS, modifications xxs, DOM based, etc. 

Hit him with the defense that will throw back his weak DDoS that doesn't even reach 100GB 🥀🙏

For volum DDoS, 
It's defense is :

- 600Gb per minute | 36Tb per hour


For Http request rate (Http exploits) 
- 6,000,000 per minute
- 360,000,000 per hour

Packet rate :
- 30,000,000 per minute
- 1,800,000,000 per hour

Connection rate (Flood Connection) 

- 300,000 per minute
- 18,000,000 per hour


Important : as axom waf is selfhosted on your device, it's defense can defend like 600Gb per minute, 
Anything above that, don't completely rely on axom waf for better defense.

---
axom waf is a selfhosted waf (defense)!!! 

> blocks DDoS attacks and eploits,floods on your web/application/graphql endpoints/REST APIs/a server
if the attacker does not use a proxy,his ip will be exposed in the program.

defends against :
SQL injection, XSS, path traversal, HTTP flood, slowloris, slow POST, regex DoS, bad user‑agents (modified user agents)
Transport L4	SYN flood, ACK flood, RST flood, FIN flood, NULL flag attack, Christmas tree attack, SYN‑ACK flood, connection flood, socket exhaustion
Network L3	UDP flood, ICMP flood (ping flood), fragmented packet flood (via raw socket detection + iptables)
State‑exhaustion	Conntrack table flood, IP fragment reassembly attack


usage = `git clone https://github.com/axom0022/axom-waf`
`cd axomwaf`
`make build`
`sudo make run`

### important =
edit config.json :

ratepersec = HTTP requests per second allowed per IP
synfloodthreshold = SYN packets/sec to trigger block
upstreamtarget = backend server (example : http://localhost:3000)
whitelist = the user you dont want to get blocked from modifying
blacklist = blacklist someone before he even tries to do anything

made by axom.
any suggestion? join our 
# [Discord](https://discord.gg/FgR3MXqZy9)
