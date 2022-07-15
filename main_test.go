package main

import (
  "testing"
  "os/exec"
  "strings"
  "github.com/k4vej/alta3research-go-cert/utils"
)

func TestMain(t *testing.T) {
  cmd := exec.Command("go", "run", "alta3research-gocert01.go")
  boutput, err := cmd.CombinedOutput()
  output := string(boutput) // because out is []byte
  utils.Ok(t, err)
  utils.Assert(t, strings.Contains(output, "swisscheese"), "Default output doesn't contain expected string 'swisscheese'")
}
