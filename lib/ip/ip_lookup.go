package ip

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"math"
	"net"
	"net/http"
	"strconv"
	"strings"
)

const DNSSERVER = "8.8.8.8:80"
const MYEXTERNALIP = "http://myexternalip.com/raw"
const TAOBAOAPI = "http://ip.taobao.com/service/getIpInfo.php"

type IPInfo struct {
	Code int    `json:"code"`
	Data IP     `json:"data"`
	Msg  string `json:"msg"`
}

type IPINT uint32

func (ip IPINT) String() string {
	var bf bytes.Buffer
	for i := 1; i <= 4; i++ {
		bf.WriteString(strconv.Itoa(int((ip >> ((4 - uint(i)) * 8)) & 0xff)))
		if i != 4 {
			bf.WriteByte('.')
		}
	}
	return bf.String()
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

	if result.Code != 0 {
		return nil, errors.New(result.Msg)
	}

	return &result, nil
}

func GetLocalIPMaskAndMac() (mask, mac string, e error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		e = err
		return
	}
	for i, a := range addrs {
		if ip, ok := a.(*net.IPNet); ok && !ip.IP.IsLoopback() {
			if ip.IP.To4() != nil {
				mask = ip.Mask.String()
				it, _ := net.InterfaceByIndex(i)
				if it != nil {
					mac = it.HardwareAddr.String()
				}
				continue
			}
		}
	}
	return
}

func GetIntraNetIPs() []string {

	var data = make([]string, 0)
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return data
	}

	for _, a := range addrs {
		if ip, ok := a.(*net.IPNet); ok && !ip.IP.IsLoopback() {
			x := getIPInt(ip)
			if len(x) > 0 {
				for _, v := range x {
					data = append(data, v.String())
				}
			}
		}
	}

	return data
}

func GetLoalNetRange() (min, max string) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return
	}

	for _, a := range addrs {
		if ip, ok := a.(*net.IPNet); ok && !ip.IP.IsLoopback() {
			x := getIPInt(ip)
			if len(x) > 0 {
				min = x[0].String()
				max = x[len(x)-1].String()
				return
			}
		}
	}
	return

}

// 根据IP和mask换算内网IP范围
func getIPInt(ipNet *net.IPNet) []IPINT {
	ip := ipNet.IP.To4()
	var min, max IPINT
	var data []IPINT
	for i := 0; i < 4; i++ {
		b := IPINT(ip[i] & ipNet.Mask[i])
		min += b << ((3 - uint(i)) * 8)
	}
	one, _ := ipNet.Mask.Size()
	max = min | IPINT(math.Pow(2, float64(32-one))-1)

	for i := min; i < max; i++ {
		if i&0x000000ff == 0 {
			continue
		}
		data = append(data, i)
	}
	return data
}

func GetIPBySegment(cider string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cider)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	return ips[1 : len(ips)-1], nil
}

func inc(ip net.IP) {
	for i := len(ip) - 1; i >= 0; i-- {
		ip[i]++
		if ip[i] > 0 {
			break
		}
	}
}
