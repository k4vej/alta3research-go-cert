/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
