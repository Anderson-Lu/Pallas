package cmd

import (
	"Pallas/lib/arp"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var scanType string
var iface string
var netSegment string

var scanDesc = `
Scan tool including intranet host scanning, ICMP detection, ARP detection etc. Need to run with administrator rights.

ARP scanner: 
  Detect the IP and MAC addresses of active hosts by broadcasting ARP packets to the LAN.

ICMP/Ping scanner:
  ICMP detection sends ICMP messages to the host IP in the LAN to detect whether the host is online, that is, Ping.
`

func init() {
	scanCmd.PersistentFlags().StringVar(&scanType, "type", "", "Scan type, choose one [ARP|ICMP|PING]")
	scanCmd.PersistentFlags().StringVar(&iface, "iface", "", "Host NIC name, eg eth0")
	scanCmd.PersistentFlags().StringVar(&netSegment, "segment", "", "LAN segment, for example 11.11.11.1/23")
}

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan is a collection of network scanning tools.",
	Long:  buildLongDesc(scanDesc),
	Run: func(cmd *cobra.Command, args []string) {
		spinner := newSpinner()
		spinner.Start()
		switch strings.ToUpper(scanType) {
		case "ARP":
			data, err := arp.ScanARP(iface, netSegment)
			if err != nil {
				fmt.Println("[error]", err.Error())
			}
			table := newTable([]string{"IP", "Mac"})
			for _, v := range data {
				table.Append([]string{v.IP, v.Mac})
			}
			spinner.Stop()
			table.Render()
		}
	},
}
