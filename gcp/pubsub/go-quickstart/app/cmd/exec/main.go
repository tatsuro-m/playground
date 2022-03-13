package main

import (
	"fmt"
	"os"
	"pubsubgo/pkg/publish"
)

func main() {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "./infrastructure/terraform/stg/output/secrets/projects/playground-318023/serviceAccounts/stg-pubsub-go-main-topic@playground-318023.iam.gserviceaccount.com.json")
	projectID := "playground-318023"
	topicID := "stg-pubsub-go-my-topic"

	err := publish.Publish(os.Stdout, projectID, topicID, "test message!!")
	if err != nil {
		fmt.Println(err)
	}
}
