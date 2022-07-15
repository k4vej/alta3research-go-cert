/* Author: Kayvan Javid
   This application is a submission for certification for the Alta3 Research GoLang proficiency course.

   This is the tcp subcommand of the main CLI, for more info run:
   swisscheese tcp --help

   By itself it performs no action, so has not real implementation other than to wrap further subcommands.
*/

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var tcpCmd = createTCPCmd()

func createTCPCmd() *cobra.Command {
	return &cobra.Command {
		Use:   "tcp",
		Short: "Provides TCP connectivity both client and server",
		Long: `
For those occasions when all of the decades worth of established,
reliable and convenient TCP tools are inconveniently not available.
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Visit(func(flag *pflag.Flag) {
				fmt.Printf("Flag supplied %v: %v\n", flag.Name, flag.Value)
			})
			return nil
		},
	}
}

func init() {
	rootCmd.AddCommand(tcpCmd)
}
