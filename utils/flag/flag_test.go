package flag

import (
	"flag"
	"os"
	"reflect"
	"testing"
)

func setArgs(args []string) {
	os.Args = args
}

func TestNewOpts(t *testing.T) {
	var flagSet = flag.NewFlagSet("test", flag.ExitOnError)
	var opts = NewOpts("valid", true, "", flagSet)
	if !reflect.DeepEqual(flagSet, opts.FlagSet){
		t.Error(flagSet, opts.FlagSet)
	}
}

func TestString_Set(t *testing.T) {
	var baseArgs = os.Args
	defer func() {
		setArgs(baseArgs)
	}()
	os.Args = []string{baseArgs[0], "--paraasdddm=test"}
	s := String{
		Opts: NewOpts("paraasdddm", "", "c"),
	}
	s.Set()
	flag.Parse()
	if s.Value != "test" {
		t.Error(s.Value)
	}
}

func TestBool_Set(t *testing.T) {
	var flagSet = flag.NewFlagSet("test", flag.ExitOnError)

	b := Bool{
		Opts: NewOpts("valid", false, "", flagSet),
	}
	if !reflect.DeepEqual(flagSet, b.FlagSet){

	}
	b.Set()
	if err := flagSet.Parse([]string{"--valid=true"}); err != nil {
		panic(err)
	}
	if b.Value != true {
		t.Error(b.Value)
	}
}
