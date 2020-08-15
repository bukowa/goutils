package flag

import (
	"flag"
	"fmt"
	"reflect"
	"testing"
)

// create flags with new flagset
func newConfigFlags() *ConfigFlags {
	fs := flag.NewFlagSet("test", flag.ExitOnError)
	cf := NewConfigFlags()
	cf.SetFlagSet(fs)
	return cf
}

func TestConfigFlags_SetFlagSet(t *testing.T) {
	fs := flag.NewFlagSet("test", flag.ExitOnError)
	cf := NewConfigFlags()
	cf.SetFlagSet(fs)

	if !reflect.DeepEqual(fs, cf.Path.FlagSet) {
		t.Error()
	}
}

func TestNewConfigFlags(t *testing.T) {
	NewConfigFlags()
}

func ExampleConfigFlags_NewClient() {
	cf := NewConfigFlags()

	// by default its home directory config file
	cf.Path.Default = "./kube/config"

	// default flagset == `os.CommandLine`
	fs := flag.NewFlagSet("test", flag.ExitOnError)
	cf.SetFlagSet(fs)
	cf.SetFlags()

	// parse flags or call directly `flags.Parse`
	if err := fs.Parse([]string{""}); err != nil {
		panic(err)
	}

	client, err := cf.NewClient()
	fmt.Println(client, err)
	// Output: <nil> CreateFile ./kube/config: The system cannot find the path specified.
}
