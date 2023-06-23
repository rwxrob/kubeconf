package main

import (
	"log"

	"github.com/rwxrob/kubeconf"
)

func main() {
	// TODO observe --kubeconfig by overriding KUBECONFIG env variable
	c, err := kubeconf.Load()
	if err != nil {
		log.Print(err)
	}
	c.Print()
}
