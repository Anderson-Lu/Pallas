package dns

import (
	"fmt"
	"io/ioutil"
	"net"
	"strings"
	"time"

	whoisparser "github.com/likexian/whois-parser-go"
)

const (
	IANA_WHOIS_SERVER  = "whois.iana.org"
	DEFAULT_WHOIS_PORT = "43"
)

func WhoisBeautifyCN(domain string, servers ...string) (map[string]interface{}, error) {
	raw, err := WhoisBeautify(domain, servers...)
	if err != nil {
		return nil, err
	}
	var ret = make(map[string]interface{}, 0)

	if raw.Registrar != nil {
		ret["注册商"] = raw.Registrar.Name
	}

	if raw.Domain != nil {
		ret["注册日期"] = raw.Domain.CreatedDate
		ret["到期日期"] = raw.Domain.ExpirationDate
		ret["更新时间"] = raw.Domain.UpdatedDate
		ret["域名状态"] = raw.Domain.Status
		ret["DNS服务器"] = raw.Domain.NameServers
	}

	if raw.Registrant != nil {
		ret["所有者"] = raw.Registrant.Name
		ret["所有者联系邮箱"] = raw.Registrant.Email
		ret["注册机构"] = raw.Registrant.Organization
		ret["所在国家"] = raw.Registrant.Country
		ret["所在省"] = raw.Registrant.Province
		ret["所在城市"] = raw.Registrant.City
	}

	return ret, nil
}

func WhoisBeautify(domain string, servers ...string) (*whoisparser.WhoisInfo, error) {
	raw, err := Whois(domain, servers...)
	if err != nil {
		return nil, err
	}

	result, err := whoisparser.Parse(raw)
	return &result, nil
}

// Whois do the whois query and returns whois info
func Whois(domain string, servers ...string) (result string, err error) {
	domain = strings.Trim(strings.TrimSpace(domain), ".")
	if domain == "" {
		return "", fmt.Errorf("whois: domain is empty")
	}

	if !strings.Contains(domain, ".") && !strings.Contains(domain, ":") {
		return query(domain, IANA_WHOIS_SERVER)
	}

	var server string
	if len(servers) == 0 || servers[0] == "" {
		ext := getExtension(domain)
		result, err := query(ext, IANA_WHOIS_SERVER)
		if err != nil {
			return "", fmt.Errorf("whois: query for whois server failed: %v", err)
		}
		server = getServer(result)
		if server == "" {
			return "", fmt.Errorf("whois: no whois server found for domain: %s", domain)
		}
	} else {
		server = strings.ToLower(servers[0])
	}

	result, err = query(domain, server)
	if err != nil {
		return
	}

	refServer := getServer(result)
	if refServer == "" || refServer == server {
		return
	}

	data, err := query(domain, refServer)
	if err == nil {
		result += data
	}

	return
}

// query send query to server
func query(domain, server string) (string, error) {
	if server == "whois.arin.net" {
		domain = "n + " + domain
	}

	conn, err := net.DialTimeout("tcp", net.JoinHostPort(server, DEFAULT_WHOIS_PORT), time.Second*30)
	if err != nil {
		return "", fmt.Errorf("whois: connect to whois server failed: %v", err)
	}

	defer conn.Close()
	_ = conn.SetWriteDeadline(time.Now().Add(time.Second * 30))
	_, err = conn.Write([]byte(domain + "\r\n"))
	if err != nil {
		return "", fmt.Errorf("whois: send to whois server failed: %v", err)
	}

	_ = conn.SetReadDeadline(time.Now().Add(time.Second * 30))
	buffer, err := ioutil.ReadAll(conn)
	if err != nil {
		return "", fmt.Errorf("whois: read from whois server failed: %v", err)
	}

	return string(buffer), nil
}

// getExtension returns extension of domain
func getExtension(domain string) string {
	ext := domain

	if net.ParseIP(domain) == nil {
		domains := strings.Split(domain, ".")
		ext = domains[len(domains)-1]
	}

	if strings.Contains(ext, "/") {
		ext = strings.Split(ext, "/")[0]
	}

	return ext
}

// getServer returns server from whois data
func getServer(data string) string {
	tokens := []string{
		"Registrar WHOIS Server: ",
		"whois: ",
	}

	for _, token := range tokens {
		start := strings.Index(data, token)
		if start != -1 {
			start += len(token)
			end := strings.Index(data[start:], "\n")
			return strings.TrimSpace(data[start : start+end])
		}
	}

	return ""
}
