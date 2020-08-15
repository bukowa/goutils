package flag

import (
	"flag"
	"log"
	"os"
	"testing"
)

func setArgs(args []string) {
	os.Args = args
}

func TestString_Set(t *testing.T) {
	var baseArgs = os.Args
	log.Println(baseArgs)
	defer func() {
		setArgs(baseArgs)
	}()
	os.Args = []string{baseArgs[0], "--paraasdddm=test"}
	s := String{
		Opts: Opts{
			Name:    "paraasdddm",
			Default: "",
			Help:    "c",
		},
	}
	s.Set()
	flag.Parse()
	if s.Value != "test" {
		t.Error(s.Value)
	}
}
