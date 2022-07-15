package cmd

import (
  "testing"
  "strings"
  "github.com/k4vej/alta3research-go-cert/utils"
)

func Test_createRootCmd(t *testing.T) {
  command := createRootCmd()
  output, err := utils.ExecuteCobraCmd(command, []string{""})
  utils.Ok(t, err)
  utils.Assert(t, strings.Contains(output, "swisscheese"), "Default output doesn't contain expected string 'swisscheese'")
}
