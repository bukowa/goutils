package flag

import (
	"github.com/bukowa/goutils/pkg"
	"github.com/bukowa/goutils/pkg/flag"
	"path/filepath"
)

const (
	Name        = "kubeconfig"
	Help        = "path to kubernetes config file"
	ClusterPath = "/var/run/secrets/kubernetes.io/serviceaccount/token"
)

var FlagPathHome = &flag.String{
	Opts: flag.Opts{
		Name:    Name,
		Default: filepath.Join(pkg.HomeOrWd(), ".kube", "config"),
		Help:    Help,
	},
}

var PathCluster = &flag.String{
	Opts: flag.Opts{
		Name:    Name,
		Default: ClusterPath,
		Help:    Help,
	},
}
