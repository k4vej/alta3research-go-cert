/* Author: Kayvan Javid
   This application is a submission for certification for the Alta3 Research GoLang proficiency course.

   This is the real main entry point as per orthodox CobraCLI projects.
   Execute is wrapped by main.main() - providing the default CLI command if no additional
   parameters or subcommands are specified, see: Execute().
*/

package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = createRootCmd()

// represents the base command when called without any subcommands
func createRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "swisscheese",
		Short: "Like a Swiss army knife, just less useful.... but more tasty",
		Long: `
This CLI to provides numerous utilities which can emulate the
functionality of other readily available tools - but for one reason or another
those "readily available tools" are not so..... readily available.
For example if you are operating in a restricted environment which doesn't have
open access to obtain or install aforementioned "readily available tools",
thus providing a poor mans (or womans, or non-binary or both) alternative.
	
In all likelyhood the functionality you really want is unlikely to be emulated,
hence the name swisscheese - a tool which is full of holes.... like the cheese.

Please enjoy this minimum set of uselessness:
- DNS resolution
- TCP socket client
- TCP socket server
- CPU info output
`,
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}


