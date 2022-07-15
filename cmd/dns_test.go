package cmd

import (
  "testing"
  "strings"
  "github.com/k4vej/alta3research-go-cert/utils"
)

func Test_createDNSCmd(t *testing.T) {
  cmd := createDNSCmd()
  output, err := utils.ExecuteCobraCmd(cmd, []string{"--help"})
  utils.Ok(t, err)
  utils.Assert(t, strings.Contains(output, "help for dns"), "Default output doesn't contain expected string 'help for dns'")
}
