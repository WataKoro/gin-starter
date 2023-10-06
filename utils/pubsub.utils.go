package utils

import (
	"context"
	"encoding/json"
	"gin-starter/config"
	"log"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

// initPubSubClient initiate pubsub client
func initPubSubClient(ctx context.Context, config config.Config) (*pubsub.Client, error) {
	projectID := config.Google.ProjectID
	c, err := pubsub.NewClient(ctx, projectID, option.WithCredentialsFile(config.Google.ServiceAccountFile))
	if err != nil {
		return nil, err
	}
	return c, nil
}

// SendTopic sending pubsub
func SendTopic(ctx context.Context, config config.Config, topicName string, payload interface{}) error {
	client, err := initPubSubClient(ctx, config)
	if err != nil {
		return err
	}

	defer func() {
		if err = client.Close(); err != nil {
			log.Println(err)
		}
	}()

	topic := client.Topic(topicName)

	itemJSON, _ := json.Marshal(payload)

	res := topic.Publish(ctx, &pubsub.Message{
		Data: itemJSON,
	})

	_, err = res.Get(ctx)
	if err != nil {
		return err
	}

	return nil
}
