package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/rwxrob/kubeconf"
)

func asJSON(a any) string {
	if a == nil {
		return `null`
	}
	byt, err := json.Marshal(a)
	if err != nil {
		log.Print(err)
		return `null`
	}
	return string(byt)
}

func printAsJSON(a any) { fmt.Println(asJSON(a)) }

func main() {
	c, err := kubeconf.Load()
	if err != nil {
		log.Print(err)
	}
	printAsJSON(c)
}
