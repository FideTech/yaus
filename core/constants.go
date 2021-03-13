package core

import (
	"fmt"
)

//Commit holds the commit hash
var Commit = "local"

//Version holds version of this server
var Version = "0.1.0"

//GetVersion returns a formatted string of the version and commit information
func GetVersion() string {
	return fmt.Sprintf("%s-%s", Version, Commit)
}
