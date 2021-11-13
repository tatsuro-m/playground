package main

import (
	"applymanifest/lib"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("Start!")
	getCredential()
}

func getCredential() {
	cmd := splitCmd(fmt.Sprintf("gcloud container clusters get-credentials %s --region %s", lib.ClusterName, lib.ClusterRegion))
	out, err := exec.Command(cmd[0], cmd[1:]...).Output()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}

func splitCmd(cmd string) []string {
	return strings.Split(cmd, " ")
}
