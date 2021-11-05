package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func main() {
	kFilePath := "./kustomization.yaml"
	b, err := ioutil.ReadFile(kFilePath)
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

	secretFileName := "secret.yaml"
	if val, ok := m["resources"]; ok {
		s, ok := val.([]interface{})
		if !ok {
			fmt.Println("cast できなかった")
		}

		s = append(s, secretFileName)
		m["resources"] = s
	} else {
		m["resources"] = []string{secretFileName}
	}

	f, err := os.OpenFile(kFilePath, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	data, err := yaml.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
	f.Write(data)
}
