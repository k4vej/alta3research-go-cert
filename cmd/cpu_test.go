package cmd

import (
  "testing"
  "strings"
  "github.com/k4vej/alta3research-go-cert/utils"
)

func Test_createCPUCmd(t *testing.T) {
  cmd := createCPUCmd()
  output, err := utils.ExecuteCobraCmd(cmd, []string{"--help"})
  utils.Ok(t, err)
  utils.Assert(t, strings.Contains(output, "help for cpu"), "Default output doesn't contain expected string 'help for cpu'")
}
