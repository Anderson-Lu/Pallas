package dns

import (
	"fmt"
	"net"
)

func LookUpDNSSummary(domain string) (map[string]interface{},error) {
	
	var ret = make(map[string]interface{},0)
	
	ns,_ := LookUpDnsNS(domain)
	a, _ := LookUpDnsForward(domain)
	mx, _ := LookUpDnsMX(domain)
	cname, _ := LookUpDnsCNAME(domain)
	ptr, _ := LookUpDnsPTR(domain)
	aaaa, _ := LookUpDnsIPV6(domain)

	ret["NS"] = ns
	ret["A"] = a
	ret["MX"] = mx
	ret["CNAME"] = cname
	ret["PTR"] = ptr
	ret["AAAA"] = aaaa

	return ret, nil
}

func LookUpDnsIPV6(domain string) (string, error) {
	x, err := net.ResolveIPAddr("ip6", domain)
	if err !=nil {
		return "", err
	}
	return x.IP.To16().String(),nil
}

// LookUpDnsForward return A records
func LookUpDnsForward(domain string) ([]string, error) {
	ips, err := net.LookupIP(domain)
	if err != nil {
		return []string{}, err
	}

	var ret = make([]string,0)
	for _, ip := range ips {	
		if _, v := ParseIP(ip.String());v == 4 {
			ret =append(ret,ip.String())
		}		
	} 

	return ret,nil
}

func ParseIP(s string) (net.IP, int) {
	ip := net.ParseIP(s)
	if ip == nil {
		return nil, 0
	}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '.':
			return ip, 4
		case ':':
			return ip, 6
		}
	}
	return nil, 0
}

// LookUpDnsForward return CNAME records
func LookUpDnsCNAME(domain string) (string, error) {
	cname, err := net.LookupCNAME(domain)
	if err != nil {
		return "", err
	}

	return cname,nil
}

// LookUpDnsPTR return PTR records
func LookUpDnsPTR(server string) ([]string, error) {
	ptr, err := net.LookupAddr(server)
	if err != nil {
		return []string{}, err
	}
	return ptr,nil
}

// LookUpDnsNS return NS records
func LookUpDnsNS(domain string) ([]string, error) {
	ns, err := net.LookupNS(domain)
	if err != nil {
		return []string{}, err
	}
	var ret = make([]string,0)
	for _,v := range ns {
		ret = append(ret,v.Host)
	}
	return ret,nil
}

// LookUpDnsMX return NS records
func LookUpDnsMX(domain string) ([]string, error) {
	ns, err := net.LookupMX(domain)
	if err != nil {
		return []string{}, err
	}
	var ret = make([]string,0)
	for _,v := range ns {
		ret = append(ret,v.Host + " " + fmt.Sprintf("%d",v.Pref))
	}
	return ret,nil
}