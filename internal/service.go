package internal

import (
	"context"

	log "github.com/sirupsen/logrus"
)

const (
	eventsChannel = "events_channel"
)

func handleProxyRequest(ctx context.Context, prJsonStr string) {
	log.Infof("EVENT_TO_PUBLISH", prJsonStr)
	if err := Rdb.Publish(eventsChannel, prJsonStr).Err(); err != nil {
		log.Errorf("EVENT_PUBLISH_ERROR", err)
	}
}
