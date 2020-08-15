package flag

import (
	"flag"
	"github.com/bukowa/goutils/utils"
	uflag "github.com/bukowa/goutils/utils/flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"path/filepath"
)

const (
	ConfigPathFlagName = "kubeconfig"
	ConfigPathFlagHelp = "path to kubernetes config file, can be empty - API client will load in-cluster config then"
)

type ConfigFlags struct {
	Path uflag.String
}

func (cf *ConfigFlags) NewClient() (c *kubernetes.Clientset, err error) {
	cfg, err := clientcmd.BuildConfigFromFlags("", cf.Path.Value)
	if err != nil {
		return
	}
	return kubernetes.NewForConfig(cfg)
}

var ConfigPathFlag = uflag.String{
	Opts: uflag.NewOpts(
		ConfigPathFlagName,
		filepath.Join(utils.HomeOrWd(), ".kube", "config"),
		ConfigPathFlagHelp),
}

// set flags
func (cf *ConfigFlags) SetFlags() {
	cf.Path.Set()
}

// should be called after flags are parsed
// kubernetes `client-go` function `clientcmd.BuildConfigFromFlags`
// will handle empty kubeconfig path as `InClusterConfig` configuration

// NewConfigFlags returns flags used to configure
// clientcmd.BuildConfigFromFlags
func NewConfigFlags() *ConfigFlags {
	return &ConfigFlags{
		Path: ConfigPathFlag,
	}
}

func (cf *ConfigFlags) SetFlagSet(set *flag.FlagSet) {
	cf.Path.FlagSet = set
}
