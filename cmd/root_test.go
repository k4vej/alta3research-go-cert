package cmd

import (
  "testing"
  "fmt"
  "bytes"
  "strings"
  "io/ioutil"
  "github.com/spf13/cobra"
)

func Test_rootCommand(t *testing.T) {
  command := createRootCmd()
  output, err := executeTestCommand(command, []string{""})
  if err != nil {
    t.FailNow()
  }
  if !strings.Contains(output, "swisscheese") {
    fmt.Println("Output doesnt contain string: ", output)
    t.FailNow()
  }
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
