![Lint, Test, Build, and Deploy](https://github.com/kfortney/fakelogit/workflows/Lint,%20Test,%20Build,%20and%20Deploy/badge.svg?branch=master)
![Docker](https://github.com/kfortney/fakelogit/workflows/Docker/badge.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/kfortney/fakelogit)](https://goreportcard.com/report/github.com/kfortney/fakelogit)

# Fakelogit

Generate fake logs to test logging applications.

> Thanks to [gofakeit](https://github.com/brianvoe/gofakeit)

## Installation

### GO
```bash
go get -u github.com/kfortney/fakelogit
```

### Docker
```
docker run -it --rm docker.pkg.github.com/kfortney/fakelogit/fakelogit:0.0.5
```

## Options

```
Usage:
  fakelogit [command]

Available Commands:
  file        Write fakelogs to a file
  help        Help about any command
  print       Print fakelogs to screen
  syslog      Write fakelogs to a syslog server
  version     Print version

Flags:
  -h, --help            help for fakelogit
  -s, --source string   Choose log source below: ("rfc3164"|"rfc5424"|"fidelis"|"landcope"|"trendmicro"|"palottraffic"|"palothreat") (default "rfc3164")

Use "fakelogit [command] --help" for more information about a command.
```

CLI usage examples
```sh
fakelogit print --source fidelis
```
```sh
fakelogit syslog --source rfc3164 --server 0.0.0.0:5514 --count 100
```
### Event Sources
RFC3164 
```sh
<%d>%s %s %s[%d]: %s     
```
RFC5424
```sh         
<%d>%d %s %s %s %d ID%d %s %s                  
```
Fidelis Cybersecurity (CEF)
```sh    
<%d>%s fid-mgr CEF:0|Fidelis Cybersecurity|Direct|9.9|DLP Signature ID|DLP Signature Name|%d| act=alert cn1=0 cn1Label=compression cn2=%d cn2Label=vlan_id cs1=Customer DLP Policy cs1Label=policy cs2=https://fid-mgr/j/alert.html cs2Label=linkback cs3=<n/a> cs3Label=malware_name cs4=%s cs4Label=from cs5=<n/a> cs5Label=malware_type cs6=Customer Management Group cs6Label=group dpt=%d dst=%s duser=%s dvc=192.168.1.1 dvchost=fid-sensor fileHash=%s fname=random_file_%s.pdf msg=fakelogit: Alert signature for testing proto=%s reason=[{'F.PII'}] requestClientApplication=%s requestMethod=%s rt=%s sev=%d spt=32528 src=%s suser=%s target=SMTP:<%s> url=%s
```
Lancope (CEF)     
```sh    
<%d>%s lan-mgr StealthWatch[%s]: CEF:0|Lancope|StealthWatch|6.0|Notification:%d|High Traffic|%d| cs1=%s cs1Label= cs2=%s cs2Label= cs3=%s cs3Label= cs4=%s cs4Label= cs5=%s cs5Label= cs6=%s cs6Label= destinationTranslatedAddress=%s destinationTranslatedPort=%s deviceExternalId=%s dpt=%d dst=%s dvc=%s dvchost=%s dvcpid=%s end=time externalId=%s msg=%d proto=%d sourceTranslatedAddress=%s spt=%d src=%s start=time"
```
Trend Micro (CEF)
```sh    
<%d>%s trm-mgr CEF:0|Trend Micro|SMS|9.9|%s|%s|%d| act=%s app=%s c6a1=%s c6a1Label= c6a2=%s c6a2Label= c6a3=%s c6a3Label= cat=%s cn1=%d cn1Label= cn2=%d cn2Label= cn3=%d cn3Label= cnt=%d cs1=%s cs1Label= cs2=%s cs2Label= cs3=%s cs3Label= cs4=%s cs4Label= cs5=%s cs5Label= cs6=%s cs6Label= deviceInboundInterface=%s dhost=%s dntdom=%s dpt=%d dst=%s duser=%s dvchost=%s externalId=%s proto=%d request=%s requestMethod=%d rt=time sntdom=%s sourceTranslatedAddress=%s spt=%d src=%s suser=%s"
```
Palo Alto Networks (CEF)
- TRAFFIC
```sh    
<%d>%s pan-mgr CEF:0|Palo Alto Networks|PAN-OS|9.9|$s|TRAFFIC|1| act=%s app=%s cat=%s cn1=%s cn1Label=%s cnt=%s cs1=%s cs1Label=%s cs2=%s cs2Label=%s cs3=%s cs3Label=%s cs4=%s cs4Label=%s cs5=%s cs5Label=%s cs6=%s cs6Label=%s destinationTranslatedAddress=%s destinationTranslatedPort=%s deviceExternalId=%s deviceInboundInterface=%s deviceOutboundInterface=%s dpt=%s dst=%s duser=%s dvchost=%s externalId=%s fileId=%s flexString1=%s flexString1Label=%s flexString2=%s flexString2Label=%s proto=%s request=%s rt=%s sourceTranslatedAddress=%s sourceTranslatedPort=%s spt=%s src=%s suser=%s PanOSActionFlags=%s PanOSAssocID=%s PanOSContentVer=%s PanOSDGl1=%s PanOSDGl2=%s PanOSDGl3=%s PanOSDGl4=%s PanOSDstUUID=%s PanOSHTTP2Con=%s PanOSHTTPHeader=%s PanOSMonitorTag=%s PanOSPPID=%s PanOSParentSessionID=%s PanOSParentStartTime=%s PanOSRuleUUID=%s PanOSSrcUUID=%s PanOSThreatCategory=%s PanOSTunnelID=%s PanOSTunnelType=%s PanOSURLCatList=%s PanOSVsysName=%s
```
- THREAT
```sh     
<%d>%s pan-mgr CEF:0|Palo Alto Networks|PAN-OS|9.9|%d|THREAT|1| act=%s app=%s cat=%s cn1=%s cn1Label=%s cnt=%s cs1=%s cs1Label=%s cs2=%s cs2Label=%s cs3=%s cs3Label=%s cs4=%s cs4Label=%s cs5=%s cs5Label=%s cs6=%s cs6Label=%s destinationTranslatedAddress=%s destinationTranslatedPort=%s deviceExternalId=%s deviceInboundInterface=%s deviceOutboundInterface=%s dpt=%s dst=%s duser=%s dvchost=%s externalId=%s fileId=%s flexString1=%s flexString1Label=%s flexString2=%s flexString2Label=%s proto=%s request=%s rt=%s sourceTranslatedAddress=%s sourceTranslatedPort=%s spt=%s src=%s suser=%s PanOSActionFlags=%s PanOSAssocID=%s PanOSContentVer=%s PanOSDGl1=%s PanOSDGl2=%s PanOSDGl3=%s PanOSDGl4=%s PanOSDstUUID=%s PanOSHTTP2Con=%s PanOSHTTPHeader=%s PanOSMonitorTag=%s PanOSPPID=%s PanOSParentSessionID=%s PanOSParentStartTime=%s PanOSRuleUUID=%s PanOSSrcUUID=%s PanOSThreatCategory=%s PanOSTunnelID=%s PanOSTunnelType=%s PanOSURLCatList=%s PanOSVsysName=%s
```
