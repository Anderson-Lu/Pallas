package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var longDesc = `
____   __    __    __      __    ___ 
(  _ \ /__\  (  )  (  )    /__\  / __)
 )___//(__)\  )(__  )(__  /(__)\ \__ \
(__) (__)(__)(____)(____)(__)(__)(___/ v1.0

-------------------------------------------

Pallas is a collection of network security tools.
Use Pallas to help you solve most network security problems.
For suggestions or comments, please visit https://github.com/Anderson-Lu/Pallas
`

var rootCmd = &cobra.Command{
	Use:   "pallas",
	Short: "Pallas is a collection of network security tools.",
	Long:  longDesc,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}

func buildLongDesc(desc string) string {
	return fmt.Sprintf("%s\n\n%s", longDesc, desc)
}

func Execute() {

	rootCmd.AddCommand(dnsCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(whoisCmd)
	rootCmd.AddCommand(cmdCmd)
	rootCmd.AddCommand(scanCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func newSpinner() *spinner.Spinner {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	return s
}

func newTable(header []string) *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetRowLine(true)
	table.SetAutoWrapText(false)
	return table
}
