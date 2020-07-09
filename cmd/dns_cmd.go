package cmd

import (
	"Pallas/lib/dns"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	recordType string
)

var dnsCmd = &cobra.Command{
	Use:   "dnslookup [domain]",
	Short: "DNS record query tool",
	Long:  "dnslook is a dns query tool to help you query the dns records of a domain name",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("[error] no domain found, use `dnstool --help` for more infomation")
			return
		}
		fmt.Println("[Info] Start fetching data from ICANN.")
		fmt.Println("[Info] Target:" + args[0] + "\n")
		printDnsLookupSummary(args[0])
	},
}

func init() {
	dnsCmd.PersistentFlags().StringVar(&recordType, "record-type", "", "specify record type, [A|AAAA|NS|MX|PTR|SOA|CNAME]")
}

func printDnsLookupSummary(domain string) {

	s := newSpinner()
	s.Start()

	var data = make(map[string]interface{}, 0)

	switch recordType {
	case "A":
		if x, e := dns.LookUpDnsForward(domain); e != nil {
			fmt.Println("[error] ", e.Error())
		} else {
			data["A"] = x
		}
	case "AAAA":
		if x, e := dns.LookUpDnsIPV6(domain); e != nil {
			fmt.Println("[error] ", e.Error())
		} else {
			data["A"] = x
		}
	case "MX":
		if x, e := dns.LookUpDnsMX(domain); e != nil {
			fmt.Println("[error] ", e.Error())
		} else {
			data["A"] = x
		}
	case "PTR":
		if x, e := dns.LookUpDnsPTR(domain); e != nil {
			fmt.Println("[error] ", e.Error())
		} else {
			data["A"] = x
		}
	case "CNAME":
		if x, e := dns.LookUpDnsCNAME(domain); e != nil {
			fmt.Println("[error] ", e.Error())
		} else {
			data["A"] = x
		}
	case "NS":
		if x, e := dns.LookUpDnsNS(domain); e != nil {
			fmt.Println("[error] ", e.Error())
		} else {
			data["A"] = x
		}
	default:
		if x, e := dns.LookUpDNSSummary(domain); e != nil {
			fmt.Println("[error] ", e.Error())
		} else {
			data = x
		}

	}

	table := newTable([]string{"Record Type", "Values"})
	for k, v := range data {
		table.Append([]string{k, fmt.Sprintf("%v", v)})
	}

	s.Stop()
	table.Render()
}
