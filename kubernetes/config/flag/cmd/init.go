package cmd

import "github.com/bukowa/goutils/kubernetes/config/flag"

var ConfigFlag = flag.NewConfigFlags()

func init() {
	ConfigFlag.SetFlags()
}
