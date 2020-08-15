package flag

import (
	"flag"
	"fmt"
	"path"
	"reflect"
	"strings"
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

// use provided config path
func TestConfigFlags_NewClient(t *testing.T) {
	fs := flag.NewFlagSet("test", flag.ExitOnError)
	cf := NewConfigFlags()
	cf.SetFlagSet(fs)
	cf.SetFlags()
	args := []string{"--kubeconfig="+path.Join("./", "config.test.txt")}
	if err := fs.Parse(args); err != nil {
		panic(err)
	}
	client, _ := cf.NewClient()
	_, err := client.ServerVersion()
	if !strings.Contains(err.Error(), "dial tcp: lookup kajsdaoidjasiodasduiaosdusa89das7d8as97d8sad78sa97d89sadasuduasiodasiodjsaijdas98dusa98dz") {
		t.Error(err)
	}
}

// use cluster config when path is empty
func TestConfigFlags_NewClient2(t *testing.T) {
	fs := flag.NewFlagSet("test", flag.ExitOnError)
	cf := NewConfigFlags()
	cf.SetFlagSet(fs)
	cf.SetFlags()
	args := []string{"--kubeconfig="}
	if err := fs.Parse(args); err != nil {
		panic(err)
	}
	_, err := cf.NewClient()
	if !strings.Contains(err.Error(), "invalid configuration: no configuration has been provided, try setting KUBERNETES_MASTER environment variable") {
		t.Error(err)
	}
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
}
