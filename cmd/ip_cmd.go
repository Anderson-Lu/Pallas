package cmd

import (
	"Pallas/lib/ip"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var ipTarget string

func init() {
	cmdCmd.PersistentFlags().StringVar(&ipTarget, "ip", "", "Specify an IP address")
}

var cmdCmd = &cobra.Command{
	Use:   "iplookup",
	Short: "IP query tool",
	Long:  buildLongDesc("IP query tool, including gathering local ip, public ip and detail for target ip."),
	Run: func(cmd *cobra.Command, args []string) {
		if ipTarget != "" {
			getTargetIPInfo(ipTarget)
		} else {
			getLocalIPInfo()
		}
	},
}

func getTargetIPInfo(ipAddr string) {

	fmt.Println("Start fetching information, target:", ipAddr)
	spinner := newSpinner()
	spinner.Start()

	ipInfo, err := ip.GetIPDetailFromTabaoAPI(ipAddr)
	spinner.Stop()

	if err != nil {
		fmt.Println("[error] ", err.Error())
		return
	}

	data := newTable([]string{"Attribute", "Value"})
	data.Append([]string{"ISP", ipInfo.Data.Isp})
	data.Append([]string{"Country", ipInfo.Data.Country})
	data.Append([]string{"Area", ipInfo.Data.Area})
	data.Append([]string{"Region", ipInfo.Data.Region})
	data.Append([]string{"City", ipInfo.Data.City})

	data.Render()
}

func getLocalIPInfo() {

	spinner := newSpinner()
	spinner.Start()

	data := newTable([]string{"Attribute", "Value"})

	if localIP, err := ip.GetLocalIP(); err != nil {
		fmt.Println("[error]", err.Error())
	} else {
		data.Append([]string{"Local IP", strings.Join(localIP, "\n")})
	}

	if publicIP, err := ip.GetPublicIPFromMyExternalIP(); err != nil {
		fmt.Println("[error]", err.Error())
	} else {
		data.Append([]string{"Public IP(From myexternalip.com)", publicIP})
	}

	if publicIP, err := ip.GetPulicIPFromDnsServer(); err != nil {
		fmt.Println("[error]", err.Error())
	} else {
		data.Append([]string{"Public IP(From Dns Server)", publicIP})
	}

	if mask, mac, err := ip.GetLocalIPMaskAndMac(); err != nil {
		fmt.Println("[error]", err.Error())
	} else {
		data.Append([]string{"Mac", mac})
		data.Append([]string{"Mask", mask})
	}

	min, max := ip.GetLoalNetRange()
	data.Append([]string{"IP Range", fmt.Sprintf("%s ~ %s", min, max)})

	spinner.Stop()
	data.Render()
}
