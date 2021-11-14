package main

import (
	"buildpush/lib"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("build and push start")
	getCredential()

	// コミットSHA とかを動的に取っても良いけど今回は簡単に latest でやる
	tag := "latest"
	fmt.Println(tag)
	for _, service := range lib.Services {
		dockerBuildAndPush(service, tag)
	}
}

func dockerBuildAndPush(service string, tag string) {
	imageTag := fmt.Sprintf("%s/%s/%s-%s-%s/%s:%s", lib.Host, lib.ProjectID, lib.Env, lib.AppName, service, service, tag)

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

func getCredential() {
	cmd := splitCmd(fmt.Sprintf("gcloud container clusters get-credentials %s --region %s", lib.ClusterName, lib.ClusterRegion))
	out, _ := exec.Command(cmd[0], cmd[1:]...).Output()
	fmt.Println(string(out))
}
