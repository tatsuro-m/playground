package publish

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/pubsub"
)

func Publish(w io.Writer, projectID, topicID, msg, orderingKey string) error {
	// projectID := "my-project-id"
	// topicID := "my-topic"
	// msg := "Hello World"
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	t := client.Topic(topicID)
	t.EnableMessageOrdering = true
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
		Attributes: map[string]string{
			"origin": "golang",
			"user":   "gcp",
			"test":   "true",
		},
		OrderingKey: orderingKey,
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("Get: %v", err)
	}
	fmt.Fprintf(w, "Published a message; msg ID: %v\n", id)
	return nil
}
