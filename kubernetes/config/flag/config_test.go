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
func newConfigFlags() *ConfigFlag {
	fs := flag.NewFlagSet("test", flag.ExitOnError)
	cf := NewConfigFlag()
	cf.FlagSet = fs
	return cf
}

func TestConfigFlags_SetFlagSet(t *testing.T) {
	fs := flag.NewFlagSet("test", flag.ExitOnError)
	cf := NewConfigFlag()
	cf.FlagSet = fs

	if !reflect.DeepEqual(fs, cf.FlagSet) {
		t.Error()
	}
}

func TestNewConfigFlags(t *testing.T) {
	NewConfigFlag()
}

// use provided config path
func TestConfigFlags_NewClient(t *testing.T) {
	fs := flag.NewFlagSet("test", flag.ExitOnError)
	cf := NewConfigFlag()
	cf.FlagSet = fs
	cf.Set()
	args := []string{"--kubeconfig=" + path.Join("./", "_config.test.txt")}
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
	cf := NewConfigFlag()
	cf.FlagSet = fs
	cf.Set()
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
	cf := NewConfigFlag()

	// by default its home directory config file
	cf.Default = "./kube/config"

	// default flagset == `os.CommandLine`
	fs := flag.NewFlagSet("test", flag.ExitOnError)
	cf.FlagSet = (fs)
	cf.Set()

	// parse flags or call directly `flags.Parse`
	if err := fs.Parse([]string{""}); err != nil {
		panic(err)
	}

	client, err := cf.NewClient()
	fmt.Println(client, err)
}
