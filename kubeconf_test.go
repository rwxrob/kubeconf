package kubeconf_test

import (
	"fmt"
	"os"

	"github.com/rwxrob/kubeconf"
)

func ExampleLoad() {

	prev := os.Getenv(`KUBECONFIG`)
	defer os.Setenv(`KUBECONFIG`, prev)
	os.Setenv(`KUBECONFIG`, `/tmp/conf1:/tmp/conf2`)

	kconf, err := kubeconf.Load()
	if err != nil {
		fmt.Println(err)
	}
	kconf.Print()

	// Output:
	// something localized
}
