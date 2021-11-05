package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func main() {
	b, err := ioutil.ReadFile("./kustomization.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}
	var m map[string]interface{}

	err = yaml.Unmarshal(b, &m)
	if err != nil {
		fmt.Println(err)
		return
	}

	for k, v := range m {
		fmt.Printf("key: %s, value: %d\n", k, v)
	}
}
