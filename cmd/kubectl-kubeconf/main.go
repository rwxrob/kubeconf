package main

import (
	"log"

	"github.com/rwxrob/kubeconf"
)

func main() {
	c, err := kubeconf.Load()
	if err != nil {
		log.Print(err)
	}
	c.Print()
}
