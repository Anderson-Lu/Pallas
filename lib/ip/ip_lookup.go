package ip

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

const DNSSERVER = "8.8.8.8:80"
const MYEXTERNALIP = "http://myexternalip.com/raw"
const TAOBAOAPI = "http://ip.taobao.com/service/getIpInfo.php"

type IPInfo struct {
	Code int `json:"code"`
	Data IP  `json:"data"`
}

type IP struct {
	Country   string `json:"country"`
	CountryId string `json:"country_id"`
	Area      string `json:"area"`
	AreaId    string `json:"area_id"`
	Region    string `json:"region"`
	RegionId  string `json:"region_id"`
	City      string `json:"city"`
	CityId    string `json:"city_id"`
	Isp       string `json:"isp"`
}

func GetPulicIPFromDnsServer() (string, error) {
	conn, e := net.Dial("udp", DNSSERVER)
	if e != nil {
		return "", e
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	return localAddr[0:idx], nil
}

func GetPublicIPFromMyExternalIP() (string, error) {
	resp, err := http.Get(MYEXTERNALIP)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	return string(content), nil
}

func GetLocalIP() ([]string, error) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return []string{}, err
	}

	var ret = make([]string, 0)

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ret = append(ret, ipnet.IP.String())
			}
		}
	}
	return ret, nil
}

func GetIPDetailFromTabaoAPI(ip string) (*IPInfo, error) {

	resp, err := http.Get(TAOBAOAPI + "?ip=" + ip)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result IPInfo
	if err := json.Unmarshal(out, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
