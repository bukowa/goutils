package cmd

import "github.com/bukowa/goutils/kubernetes/config/flag"

var Config = flag.NewConfigFlag()

// by importing this package we setup a default flag set
// to use as a global ConfigFlag shortcut to get a client
func init() {
	Config.Set()
}
