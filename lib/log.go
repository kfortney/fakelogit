package lib

import (
	"fmt"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit"
)

const (
	//RFC3164Log Random log Source
	RFC3164Log = "<%d>%s %s %s[%d]: %s"
	//RFC5424Log Random log Source
	RFC5424Log = "<%d>%d %s %s %s %d ID%d %s %s"
	//FidelisLog Specific Log Source (CEF)
	FidelisLog = "<%d>%s fid-mgr CEF:0|Fidelis Cybersecurity|Direct|9.9|DLP Signature ID|DLP Signature Name|%d| act=alert cn1=0 cn1Label=compression cn2=%d cn2Label=vlan_id cs1=Customer DLP Policy cs1Label=policy cs2=https://fid-mgr/j/alert.html cs2Label=linkback cs3=<n/a> cs3Label=malware_name cs4=%s cs4Label=from cs5=<n/a> cs5Label=malware_type cs6=Customer Management Group cs6Label=group dpt=%d dst=%s duser=%s dvc=192.168.1.1 dvchost=fid-sensor fileHash=%s fname=random_file_%s.pdf msg=FakeLog: Alert signature for testing proto=%s reason=[{'F.PII'}] requestClientApplication=%s requestMethod=%s rt=%s sev=%d spt=32528 src=%s suser=%s target=SMTP:<%s> url=%s"
	//LancopeLog Specific Log Source (CEF)
	LancopeLog = "<%d>%s lan-mgr CEF:0|Lancope|StealthWatch|6.0|Notification:%d|High Traffic|%d| cs1=%s cs1Label= cs2=%s cs2Label= cs3=%s cs3Label= cs4=%s cs4Label= cs5=%s cs5Label= cs6=%s cs6Label= destinationTranslatedAddress=%s destinationTranslatedPort=%s deviceExternalId=%s dpt=%d dst=%s dvc=%s dvchost=%s dvcpid=%s end=time externalId=%s msg=%d proto=%d sourceTranslatedAddress=%s spt=%d src=%s start=time"
	//TrendMicroLog Specific Log Source (CEF)
	TrendMicroLog = "<%d>%s trm-mgr CEF:0|Trend Micro|SMS|9.9|%s|%s|%d| act=%s app=%s c6a1=%s c6a1Label= c6a2=%s c6a2Label= c6a3=%s c6a3Label= cat=%s cn1=%d cn1Label= cn2=%d cn2Label= cn3=%d cn3Label= cnt=%d cs1=%s cs1Label= cs2=%s cs2Label= cs3=%s cs3Label= cs4=%s cs4Label= cs5=%s cs5Label= cs6=%s cs6Label= deviceInboundInterface=%s dhost=%s dntdom=%s dpt=%d dst=%s duser=%s dvchost=%s externalId=%s proto=%d request=%s requestMethod=%d rt=time sntdom=%s sourceTranslatedAddress=%s spt=%d src=%s suser=%s"
	//PaloTrafficLog Specific Log Source (CEF)
	PaloTrafficLog = "<%d>%s pan-mgr CEF:0|Palo Alto Networks|PAN-OS|9.9|$s|TRAFFIC|1| act=%s app=%s cat=%s cn1=%s cn1Label=%s cnt=%s cs1=%s cs1Label=%s cs2=%s cs2Label=%s cs3=%s cs3Label=%s cs4=%s cs4Label=%s cs5=%s cs5Label=%s cs6=%s cs6Label=%s destinationTranslatedAddress=%s destinationTranslatedPort=%s deviceExternalId=%s deviceInboundInterface=%s deviceOutboundInterface=%s dpt=%s dst=%s duser=%s dvchost=%s externalId=%s fileId=%s flexString1=%s flexString1Label=%s flexString2=%s flexString2Label=%s proto=%s request=%s rt=%s sourceTranslatedAddress=%s sourceTranslatedPort=%s spt=%s src=%s suser=%s PanOSActionFlags=%s PanOSAssocID=%s PanOSContentVer=%s PanOSDGl1=%s PanOSDGl2=%s PanOSDGl3=%s PanOSDGl4=%s PanOSDstUUID=%s PanOSHTTP2Con=%s PanOSHTTPHeader=%s PanOSMonitorTag=%s PanOSPPID=%s PanOSParentSessionID=%s PanOSParentStartTime=%s PanOSRuleUUID=%s PanOSSrcUUID=%s PanOSThreatCategory=%s PanOSTunnelID=%s PanOSTunnelType=%s PanOSURLCatList=%s PanOSVsysName=%s"
	//PaloThreatLog Specific Log Source (CEF)
	PaloThreatLog = "<%d>%s pan-mgr CEF:0|Palo Alto Networks|PAN-OS|9.9|%d|THREAT|1| act=%s app=%s cat=%s cn1=%s cn1Label=%s cnt=%s cs1=%s cs1Label=%s cs2=%s cs2Label=%s cs3=%s cs3Label=%s cs4=%s cs4Label=%s cs5=%s cs5Label=%s cs6=%s cs6Label=%s destinationTranslatedAddress=%s destinationTranslatedPort=%s deviceExternalId=%s deviceInboundInterface=%s deviceOutboundInterface=%s dpt=%s dst=%s duser=%s dvchost=%s externalId=%s fileId=%s flexString1=%s flexString1Label=%s flexString2=%s flexString2Label=%s proto=%s request=%s rt=%s sourceTranslatedAddress=%s sourceTranslatedPort=%s spt=%s src=%s suser=%s PanOSActionFlags=%s PanOSAssocID=%s PanOSContentVer=%s PanOSDGl1=%s PanOSDGl2=%s PanOSDGl3=%s PanOSDGl4=%s PanOSDstUUID=%s PanOSHTTP2Con=%s PanOSHTTPHeader=%s PanOSMonitorTag=%s PanOSPPID=%s PanOSParentSessionID=%s PanOSParentStartTime=%s PanOSRuleUUID=%s PanOSSrcUUID=%s PanOSThreatCategory=%s PanOSTunnelID=%s PanOSTunnelType=%s PanOSURLCatList=%s PanOSVsysName=%s"
)

//NewLog chooses new log source
func NewLog(source string) string {
	t := LogTimeFormat()
	switch source {
	case "rfc3164":
		return NewRFC3164Log(t)
	case "rfc5424":
		return NewRFC5424Log(t)
	case "fidelis":
		return NewFidelisLog(t)
	case "lancope":
		return NewLancopeLog(t)
	case "trendmicro":
		return NewTrendMicroLog(t)
	case "palottraffic":
		return NewPaloTrafficLog(t)
	case "palothreat":
		return NewPaloThreatLog(t)
	default:
		return ""
	}
}

//NewRFC3164Log Specific Log Source (CEF)
func NewRFC3164Log(t time.Time) string {
	return fmt.Sprintf(
		RFC3164Log,
		gofakeit.Number(0, 191),
		t.Format("Jan 02 15:04:05"),
		strings.ToLower(gofakeit.Username()),
		gofakeit.Word(),
		gofakeit.Number(1, 10000),
		gofakeit.HackerPhrase(),
	)
}

//NewRFC5424Log Specific Log Source (CEF)
func NewRFC5424Log(t time.Time) string {
	return fmt.Sprintf(
		RFC5424Log,
		gofakeit.Number(0, 191),
		gofakeit.Number(1, 3),
		t.Format("2006-01-02T15:04:05.000Z"),
		gofakeit.DomainName(),
		gofakeit.Word(),
		gofakeit.Number(1, 10000),
		gofakeit.Number(1, 1000),
		"-",
		gofakeit.HackerPhrase(),
	)
}

//NewFidelisLog Specific Log Source (CEF)
func NewFidelisLog(t time.Time) string {
	rt := t.Add(-1 * time.Minute)
	return fmt.Sprintf(
		FidelisLog,
		gofakeit.Number(0, 191),
		t.Format("Jan 02 15:04:05"),
		gofakeit.Number(1, 4),
		gofakeit.Number(1000, 4094),
		gofakeit.Email(),
		RandomPort(),
		gofakeit.IPv4Address(),
		gofakeit.Email(),
		RandomString(32),
		gofakeit.Numerify("##########"),
		RandomProtocol(),
		gofakeit.ChromeUserAgent(),
		RandomRequest(),
		rt.Format("Jan 02 15:04:05"),
		gofakeit.Number(1, 4),
		gofakeit.IPv4Address(),
		gofakeit.Email(),
		gofakeit.Email(),
		RandomURL(),
	)
}

//NewLancopeLog Specific Log Source (CEF) todo| map fake lancope data to the correct fields
func NewLancopeLog(t time.Time) string {
	return fmt.Sprintf(
		LancopeLog,
		gofakeit.Number(0, 191),
		t.Format("Jan 02 15:04:05"),
		gofakeit.Number(1, 4),
		gofakeit.Number(1, 4),
		gofakeit.Numerify("########"),
		gofakeit.Email(),
		gofakeit.Numerify("########"),
		gofakeit.IPv4Address(),
		gofakeit.Email(),
		RandomString(32),
		gofakeit.Numerify("##########"),
		RandomProtocol(),
		gofakeit.ChromeUserAgent(),
		gofakeit.Number(1, 4),
		gofakeit.IPv4Address(),
		gofakeit.Email(),
		gofakeit.Email(),
		RandomURL(),
		RandomString(32),
		gofakeit.Number(1, 4),
		gofakeit.Number(1, 4),
		RandomString(32),
		gofakeit.Number(1, 4),
		RandomString(32),
	)
}

//NewTrendMicroLog Specific Log Source (CEF) todo| map fake trendmicro data to the correct fields
func NewTrendMicroLog(t time.Time) string {
	return fmt.Sprintf(
		TrendMicroLog,
		gofakeit.Number(0, 191),
		t.Format("Jan 02 15:04:05"),
		RandomString(10),
		RandomString(10),
		gofakeit.Number(1, 4),
		RandomString(7),
		RandomString(8),
		RandomString(9),
		RandomString(10),
		RandomString(11),
		RandomString(10),
		gofakeit.Number(1, 6),
		gofakeit.Number(1, 7),
		gofakeit.Number(1, 8),
		gofakeit.Number(1, 9),
		RandomString(11),
		RandomString(1),
		RandomString(2),
		RandomString(3),
		RandomString(4),
		RandomString(5),
		RandomString(6),
		RandomString(7),
		RandomString(8),
		gofakeit.Number(1, 4),
		RandomString(10),
		RandomString(10),
		RandomString(10),
		RandomString(8),
		gofakeit.Number(1, 4),
		RandomString(8),
		gofakeit.Number(1, 4),
		RandomString(10),
		RandomString(10),
		gofakeit.Number(1, 4),
		RandomString(10),
		RandomString(10),
	)
}

//NewPaloTrafficLog Specific Log Source (CEF) todo| map fake palo data to the correct fields
func NewPaloTrafficLog(t time.Time) string {
	return fmt.Sprintf(
		PaloTrafficLog,
		gofakeit.Number(0, 191),
		t.Format("Jan 02 15:04:05"),
		RandomString(10),
		RandomString(10),
		RandomString(10),
		RandomString(7),
		RandomString(8),
		RandomString(9),
		RandomString(10),
		RandomString(11),
		RandomString(10),
		RandomString(10),
		RandomString(10),
		RandomString(10),
		RandomString(10),
		RandomString(11),
		RandomString(1),
		RandomString(2),
		RandomString(3),
		RandomString(4),
		RandomString(5),
		RandomString(6),
		RandomString(7),
		RandomString(8),
		RandomString(10),
		RandomString(10),
		RandomString(10),
		RandomString(10),
		RandomString(8),
		RandomString(8),
		RandomString(8),
		RandomString(10),
		RandomString(10),
		RandomString(10),
		RandomString(10),
		RandomString(10),
		RandomString(10),
		RandomString(11),
		RandomString(1),
		RandomString(2),
		RandomString(3),
		RandomString(4),
		RandomString(5),
		RandomString(6),
		RandomString(7),
		RandomString(8),
		RandomString(10),
		RandomString(10),
		RandomString(11),
		RandomString(1),
		RandomString(2),
		RandomString(3),
		RandomString(4),
		RandomString(5),
		RandomString(6),
		RandomString(7),
		RandomString(8),
		RandomString(1),
		RandomString(2),
		RandomString(3),
		RandomString(4),
		RandomString(5),
		RandomString(6),
		RandomString(7),
	)
}

//NewPaloThreatLog Specific Log Source (CEF) todo| map fake palo data to the correct fields
func NewPaloThreatLog(t time.Time) string {
	return fmt.Sprintf(
		PaloThreatLog,
		gofakeit.Number(0, 191),
		t.Format("Jan 02 15:04:05"),
		gofakeit.Number(0, 191),
		RandomString(2),
		RandomString(3),
		RandomString(7),
		RandomString(8),
		RandomString(9),
		RandomString(1),
		RandomString(2),
		RandomString(3),
		RandomString(4),
		RandomString(5),
		RandomString(6),
		RandomString(7),
		RandomString(11),
		RandomString(1),
		RandomString(2),
		RandomString(3),
		RandomString(4),
		RandomString(5),
		RandomString(6),
		RandomString(7),
		RandomString(8),
		RandomString(2),
		RandomString(3),
		RandomString(4),
		RandomString(5),
		RandomString(6),
		RandomString(7),
		RandomString(8),
		RandomString(11),
		RandomString(12),
		RandomString(13),
		RandomString(14),
		RandomString(15),
		RandomString(16),
		RandomString(17),
		RandomString(1),
		RandomString(2),
		RandomString(3),
		RandomString(4),
		RandomString(5),
		RandomString(6),
		RandomString(7),
		RandomString(8),
		RandomString(21),
		RandomString(22),
		RandomString(23),
		RandomString(1),
		RandomString(2),
		RandomString(3),
		RandomString(4),
		RandomString(5),
		RandomString(6),
		RandomString(7),
		RandomString(8),
		RandomString(1),
		RandomString(2),
		RandomString(3),
		RandomString(4),
		RandomString(5),
		RandomString(6),
		RandomString(6),
		RandomString(6),
	)
}
