package internal

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
)

var ErrUnsupportedDestination = fmt.Sprintf("destinations array should contain values only from the following supported destinations list: %v", supportedDestinations)
var ErrPublishFailedDestinationTpl = "publish failed at the following destinations: %v"

func handleProxyRequest(ctx context.Context, pr proxyRequest) error {
	for _, dest := range pr.Destinations {
		if !slices.Contains(supportedDestinations, dest) {
			log.Error("SKIPPING_UNSUPPORTED_DESTINATION", dest)
			return newCustomError(ErrUnsupportedDestination, 400)
		}
	}

	publishFailedDestinations := []string{}
	for _, dest := range pr.Destinations {
		if err := Rdb.Publish(getChannelName(dest), pr.Payload).Err(); err != nil {
			publishFailedDestinations = append(publishFailedDestinations, dest)
			log.Error("EVENT_PUBLISH_ERROR", err)
		}
		log.Info("EVENT_PUBLISHED", map[string]interface{}{
			"destination": dest,
			"payload":     pr.Payload,
		})
	}
	if len(publishFailedDestinations) > 0 {
		return newCustomError(fmt.Sprintf(ErrPublishFailedDestinationTpl, publishFailedDestinations), 500)
	}
	return nil
}
