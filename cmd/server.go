/* Author: Kayvan Javid
   This application is a submission for certification for the Alta3 Research GoLang proficiency course.

   This is the server subcommand of the tcp subcommand, for more info run:
   swisscheese tcp server --help
*/

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	sourceFlag = "source"
)

var source string

var tcpCmd_server = createTCPServerCmd()

func createTCPServerCmd() *cobra.Command {
	return &cobra.Command {
		Use:   "server",
		Short: "Provides TCP server connectivity",
		Long: `
For those occasions when all of the decades worth of established,
reliable and convenient TCP tools are inconveniently not available.
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Visit(func(flag *pflag.Flag) {
				//fmt.Printf("Flag supplied %v: %v\n", flag.Name, flag.Value)
				switch flag.Name {
					case "source":
						fmt.Println("Creating TCP Server bound to:", flag.Value)
						tcpServer(flag.Value.String())
					default:
						fmt.Println("Nothing to do - no flags supplied")
				}
			})
			return nil
		},
	}
}



func init() {
	tcpCmd.AddCommand(tcpCmd_server)
	tcpCmd_server.Flags().StringVar(&source, sourceFlag, "", "Address to connect to serve from")
}

func tcpServer(source string) {
	fmt.Println("Not implemented")
}
