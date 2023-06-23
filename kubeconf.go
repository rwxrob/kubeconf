package kubeconf

import (
	"encoding/json"
	"fmt"
	"log"

	cc "k8s.io/client-go/tools/clientcmd"
	api "k8s.io/client-go/tools/clientcmd/api"
)

// Config wraps the client-go type of the same name to allow for
// additional marshalling methods.
type Config api.Config

// Load returns a Config struct created using the default loading rules
// for the clientcmd/api (the same used by kubectl itself). See kubectl
// config --help for more information.
func Load() (Config, error) {
	conf, err := cc.NewNonInteractiveDeferredLoadingClientConfig(
		cc.NewDefaultClientConfigLoadingRules(), nil).RawConfig()
	return Config(conf), err
}

// JSON outputs the Context as JSON (null if <nil>). Errors during
// unmashalling are logged and null returned.
func (c *Config) JSON() string {
	if c == nil {
		return `null`
	}
	byt, err := json.Marshal(c)
	if err != nil {
		log.Print(err)
		return `null`
	}
	return string(byt)
}

func (c Config) String() string { return c.JSON() }
func (c *Config) Print()        { fmt.Println(c) }
