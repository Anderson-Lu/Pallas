package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionTxt = `
Version    :  v1.0
Author     :  Anderson Lu
UpdateTime :  2020-07
Github     :  https://github.com/Anderson-Lu/Pallas
`

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print Pallas version",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(longDesc)
		fmt.Println(versionTxt)
	},
}
