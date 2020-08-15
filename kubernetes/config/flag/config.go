/*
	Package flag implements shortcuts for parsing flags
	and creating kubernetes client from them.

	Usage

	Create ConfigFlag that by default looks for config in your home dir and use a shortcut to create kubernetes.Clientset
		import kflag "github.com/bukowa/goutils/kubernetes/config/flag"
		import "flag"

		cf := kflag.NewConfigFlag()
		// set flags
		cf.Set()
		// parse flags
		flag.Parse()
		// create kubernetes client
		c, err := kubeConfigFlag.NewClient()

	By default config file from home directory will be used `$HOME/.kube/config`.
	If you set --kubeconfig to empty value then client will use in-cluster config files.

		cf := kflag.NewConfigFlag()
		// change default to use in-cluster config
		cf.Default = ""
		// parse flags
		flag.Parse()
		// create kubernetes client
		c, err := kubeConfigFlag.NewClient()
*/
package flag

import (
	"flag"
	"github.com/bukowa/goutils/utils"
	uflag "github.com/bukowa/goutils/utils/flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"path"
)

const (
	ConfigPathFlagName = "kubeconfig"
	ConfigPathFlagHelp = "path to kubernetes config file, can be empty - API client will load in-cluster config then"
)

type ConfigFlag struct {
	*uflag.String
}

func (cf *ConfigFlag) NewClient() (c *kubernetes.Clientset, err error) {
	cfg, err := clientcmd.BuildConfigFromFlags("", cf.Value)
	if err != nil {
		return
	}
	return kubernetes.NewForConfig(cfg)
}

func NewConfigFlag() *ConfigFlag {
	return &ConfigFlag{&uflag.String{
		Opts: uflag.NewOpts(
			ConfigPathFlagName,
			path.Join(utils.HomeDir(), ".kube", "config"),
			ConfigPathFlagHelp),
	}}
}

func (cf *ConfigFlag) SetFlagSet(fs *flag.FlagSet) {
	cf.FlagSet = fs
}
