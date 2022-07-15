/* Author: Kayvan Javid
   This application is a submission for certification for the Alta3 Research GoLang proficiency course.

   This is the dns subcommand of the main CLI, it uses the GoLang DNS resolver to perform forward and
   reverse DNS lookups of the specified addresses, for more info:
   swisscheese dns --ip|--hostname
*/

package cmd

import (
	"fmt"
	"net"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	ipFlag = "ip"
	hostnameFlag = "hostname"
)

var ip, hostname string

var dnsCmd = createDNSCmd()

func createDNSCmd() *cobra.Command {
	return &cobra.Command {
		Use:   "dns",
		Short: "Provides DNS resolution both forward and reverse",
		Long: `
For those occasions when all of the decades worth of established,
reliable and convenient DNS resolution tools are inconveniently not available.
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Visit(func(flag *pflag.Flag) {
				//fmt.Printf("Flag supplied %v: %v\n", flag.Name, flag.Value)
				switch flag.Name {
					case "ip":
						fmt.Println("Performing reverse DNS resolution for supplied ip:", flag.Value.String())
						reverseLookupIPv4(flag.Value.String())
					case "hostname":
						fmt.Println("Performing forward DNS resolution for supplied hostname:", flag.Value)
						forwardLookupHostname(flag.Value.String())
					default:
						fmt.Println("Nothing to do - neither ip nor hostname supplied")
				}
			})
			return nil
		},
	}
}

func init() {
	rootCmd.AddCommand(dnsCmd)

	// flags and their configuration
	dnsCmd.Flags().StringVar(&ip, ipFlag, "", "IPv4 address to perform reverse DNS resolution to hostname")
	dnsCmd.Flags().StringVar(&hostname, hostnameFlag, "", "Hostname to perform forward DNS resolution to IP address")
	dnsCmd.MarkFlagsMutuallyExclusive(ipFlag, hostnameFlag)
}

// Input: ipv4 address (e.g. 8.8.8.8)
// Output: List of hostnames
func reverseLookupIPv4(ipv4 string) {
	fmt.Println("Resolving address for:", ipv4)
	names, err := net.LookupAddr(ipv4)
	if err != nil {
		fmt.Println("Error resolving DNS:", err)
	}
	fmt.Println(names)
}

// Input: hostname address (e.g. google.com)
// Output: List of ip addresses
func forwardLookupHostname(hostname string) {
	fmt.Println("Resolving address for:", hostname)
	names, err := net.LookupHost(hostname)
	if err != nil {
		fmt.Println("Error resolving DNS:", err)
	}
	fmt.Println(names)
}

