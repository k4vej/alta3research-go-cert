package cmd

import (
  "testing"
  "bytes"
  "strings"
  "io/ioutil"
  "github.com/spf13/cobra"
  "github.com/k4vej/alta3research-go-cert/utils"
)

func Test_createRootCmd(t *testing.T) {
  command := createRootCmd()
  output, err := executeTestCommand(command, []string{""})
  utils.Ok(t, err)
  utils.Assert(t, strings.Contains(output, "swisscheese"), "Default output doesn't contain expected string 'swisscheese'")
}

func executeTestCommand(cmd *cobra.Command, args []string) (string, error) {
  b := bytes.NewBufferString("")
  cmd.SetOut(b)
  cmd.SetErr(b)
  cmd.SetArgs(args)

  err := cmd.Execute()
  if err != nil {
    return "", err
  }

  out, err := ioutil.ReadAll(b)
  return string(out), err
}
