/* Author: Kayvan Javid
   This application is a submission for certification for the Alta3 Research GoLang proficiency course.

   Some testing convenience wrappers to avoid the long and drawn out GoLang if err != nil nonsense,
   making the tests shorter, easier to read, and more akin to other testing frameworks used in mainstream.

   Borrowed from: github.com/benbjohnson/testing
*/

package utils

import (
	"testing"
	"fmt"
	"bytes"
	"runtime"
	"reflect"
	"io/ioutil"
	"path/filepath"
	"github.com/spf13/cobra"
)

// assert fails the test if the condition is false.
func Assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func Ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func Equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

// wraps execution of a Cobra command so that we can capture STDOUT + STDERR for inspection
func ExecuteCobraCmd(cmd *cobra.Command, args []string) (string, error) {
  captured_output := bytes.NewBufferString("")
  cmd.SetOut(captured_output)
  cmd.SetErr(captured_output)
  cmd.SetArgs(args)

  err := cmd.Execute()
  if err != nil {
    return "", err
  }

  parsed_output, err := ioutil.ReadAll(captured_output)
  if err != nil {
    return "", err
  }

  return string(parsed_output), err
}
