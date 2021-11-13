package main

import (
	"buildpush/lib"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func main() {
	fmt.Println("start")
	dockerBuildAndPush()
}

func dockerBuildAndPush() {
	for _, service := range lib.Services {
		imageTag := fmt.Sprintf("%s/%s/%s-%s-%s:%d", lib.Host, lib.ProjectID, lib.Env, lib.AppName, service, time.Now().Unix())
		cmd := splitCmd(fmt.Sprintf("docker build -f ./%s/Dockerfile --platform amd64 -t %s ./%s", service, imageTag, service))
		out, err := exec.Command(cmd[0], cmd[1:]...).Output()
		fmt.Println(err)
		fmt.Println(string(out))
	}
}

func splitCmd(cmd string) []string {
	return strings.Split(cmd, " ")
}
