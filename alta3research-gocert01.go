/* Author: Kayvan Javid
   This application is a submission for certification for the Alta3 Research GoLang proficiency course.

   This is the main entry point to the application which as in orthodox CobraCLI fashion
   is simply a thin wrapper around the root command Execute() function, see: cmd/root.go
*/

package main

import "github.com/k4vej/alta3research-go-cert/cmd"

func main() {
	cmd.Execute()
}
