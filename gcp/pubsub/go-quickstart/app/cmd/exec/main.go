package main

import (
	"fmt"
	"os"
	"pubsubgo/pkg/publish"
	"pubsubgo/pkg/pull"
)

func main() {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "./infrastructure/terraform/stg/output/secrets/projects/playground-318023/serviceAccounts/stg-pubsub-go-main-topic@playground-318023.iam.gserviceaccount.com.json")
	projectID := "playground-318023"
	topicID := "stg-pubsub-go-my-topic"

	for i := 0; i < 11; i++ {
		err := publish.Publish(os.Stdout, projectID, topicID, fmt.Sprintf("test message %d", i))
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	subID := "stg-pubsub-go-my-topic-sub1"
	err := pull.PullMsgs(os.Stdout, projectID, subID)
	if err != nil {
		fmt.Println(err)
		return
	}
}
