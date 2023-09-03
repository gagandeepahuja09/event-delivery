package internal

import (
	"context"

	log "github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
)

func handleProxyRequest(ctx context.Context, pr proxyRequest) {
	for _, dest := range pr.Destinations {
		if !slices.Contains(supportedDestinations, dest) {
			log.Errorf("SKIPPING_UNSUPPORTED_DESTINATION", dest)
			continue
		}
		if err := Rdb.Publish(getChannelName(dest), pr.Payload).Err(); err != nil {
			log.Errorf("EVENT_PUBLISH_ERROR", err)
		}
		log.Infof("EVENT_PUBLISHED", map[string]interface{}{
			"destination": dest,
			"payload":     pr.Payload,
		})
	}
}
