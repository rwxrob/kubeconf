package kubeconf

import (
	cc "k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

// Load returns an api.Config struct created using the default loading rules
// for the clientcmd/api (the same as kubectl itself). See kubectl
// config --help for more information. This is useful for viewing how
// kubectl actually resolved loading multiple configurations from
// different files.
func Load() (api.Config, error) {
	return cc.NewNonInteractiveDeferredLoadingClientConfig(
		cc.NewDefaultClientConfigLoadingRules(), nil).RawConfig()
}
