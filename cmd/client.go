/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"fmt"
	"net"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	targetFlag = "target"
)

var target string

var tcpCmd_client = createTCPClientCmd()

func createTCPClientCmd() *cobra.Command {
	return &cobra.Command {
		Use:   "client",
		Short: "Provides TCP client connectivity",
		Long: `
For those occasions when all of the decades worth of established,
reliable and convenient TCP tools are inconveniently not available.
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Visit(func(flag *pflag.Flag) {
				//fmt.Printf("Flag supplied %v: %v\n", flag.Name, flag.Value)
				switch flag.Name {
					case "target":
						log.Println("Creating TCP Client connecting to:", flag.Value)
						tcpClient(flag.Value.String())
					default:
						fmt.Println("Nothing to do - no flags supplied")
				}
			})
			return nil
		},
	}
}

func init() {
	tcpCmd.AddCommand(tcpCmd_client)
	tcpCmd_client.Flags().StringVar(&target, targetFlag, "", "Address to connect to")
}

func tcpClient(target string) {
	con, err := net.Dial("tcp", target)
	if err != nil {
		log.Fatalln(err)
	}
	defer con.Close() // tidy up after ourselves
	log.Println("Connection success")
}
