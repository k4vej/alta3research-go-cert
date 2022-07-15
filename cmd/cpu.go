/* Author: Kayvan Javid
   This application is a submission for certification for the Alta3 Research GoLang proficiency course.

   This is the cpu subcommand of the main CLI, it wraps the existing gopsutil library to expose CPU hardware
   info and utilization via CLI, for more info run:
   swisscheese cpu --info|--top
*/

package cmd

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/shirou/gopsutil/cpu"
)

const (
	infoFlag = "info"
	topFlag = "top"
)

var info, top bool

var cpuCmd = createCPUCmd()

func createCPUCmd() *cobra.Command {
	return &cobra.Command {
		Use:   "cpu",
		Short: "Provides CPU information",
		Long: `
For those occasions when all of the decades worth of established,
reliable and convenient CPU info and utilization tools are inconveniently not available.
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Visit(func(flag *pflag.Flag) {
				//fmt.Printf("Flag supplied %v: %v\n", flag.Name, flag.Value)
				switch flag.Name {
					case "info":
						fmt.Println("Retrieving CPU hardware info")
						cpuInfo()
					case "top":
						fmt.Println("Retrieving CPU utilization info")
						topInfo()
					default:
						fmt.Println("Nothing to do - neither info nor top supplied")
				}
			})
			return nil
		},
	}
}

func init() {
	rootCmd.AddCommand(cpuCmd)

	// flags and their configuration
	cpuCmd.Flags().Bool(infoFlag, false, "Obtains CPU hardware info")
	cpuCmd.Flags().Bool(topFlag, false, "Obtains CPU utilization info")
}

func cpuInfo() {
	output, _ := cpu.Info()
	spew.Dump(output)
}

func topInfo() {
	output, _ := cpu.Times(true)
	spew.Dump(output)
}

