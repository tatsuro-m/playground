package main

import (
	"buildpush/lib"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func main() {
	fmt.Println("build and push start")

	// コミットSHA とかを動的に取っても良いけど今回は簡単に unix 時でタグ付けする。
	tag := time.Now().Unix()
	fmt.Println(tag)
	for _, service := range lib.Services {
		dockerBuildAndPush(service, tag)
	}
}

func dockerBuildAndPush(service string, tag int64) {
	imageTag := fmt.Sprintf("%s/%s/%s-%s-%s/%s:%d", lib.Host, lib.ProjectID, lib.Env, lib.AppName, service, service, tag)

	cmd := splitCmd(fmt.Sprintf("docker build -f ./%s/Dockerfile --platform amd64 -t %s ./%s", service, imageTag, service))
	out, _ := exec.Command(cmd[0], cmd[1:]...).Output()
	fmt.Println(string(out))

	cmd = splitCmd(fmt.Sprintf("docker push %s", imageTag))
	out, _ = exec.Command(cmd[0], cmd[1:]...).Output()
	fmt.Println(string(out))
}

func splitCmd(cmd string) []string {
	return strings.Split(cmd, " ")
}
