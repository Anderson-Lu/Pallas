package cmd

import (
	"Pallas/lib/dns"
	"fmt"

	"github.com/spf13/cobra"
)

var beautify bool

func init() {
	whoisCmd.PersistentFlags().BoolVar(&beautify, "beautify", false, "beautify output [true]")
}

var whoisCmd = &cobra.Command{
	Version: "v1.0",
	Use:     "whois",
	Short:   "whois query domain name or IP attribution information",
	Long:    "whois query domain name or IP attribution information",
	Example: "whois baidu.com --beautify true",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("[error] no domain or IP specified")
			return
		}

		fmt.Println("whois: start fetching information for ", args[0])

		s := newSpinner()
		switch beautify {
		case true:
			s.Start()
			if x, err := dns.WhoisBeautifyCN(args[0]); err != nil {
				s.Stop()
				fmt.Println("[error] " + err.Error())
				return
			} else {
				s.Stop()
				table := newTable([]string{"Attribute", "Value"})
				for k, v := range x {
					table.Append([]string{k, fmt.Sprintf("%v", v)})
				}
				table.Render()
			}
		default:
			s.Start()
			result, err := dns.Whois(args[0])
			s.Stop()
			if err != nil {
				fmt.Println("[error] " + err.Error())
				return
			}
			fmt.Println(result)
		}
	},
}
