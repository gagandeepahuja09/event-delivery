package internal

import (
	"math/rand"
	"sync"
	"time"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

const (
	maxRetries             = 5
	errorServerUnreachable = "server_unreachable"
)

var (
	pubsub redis.PubSub
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func ProcessEvents() {
	go func() {
		defer pubsub.Close()

		var wg sync.WaitGroup
		for _, dest := range supportedDestinations {
			pubsub := Rdb.Subscribe(getChannelName(dest))

			_, err := pubsub.Receive()
			if err != nil {
				log.Fatal(err)
			}

			wg.Add(1)
			go func(dest string) {
				defer wg.Done()
				consumePayloads(pubsub.Channel(), dest)
			}(dest)
		}

		wg.Wait()
	}()
}

func consumePayloads(channel <-chan *redis.Message, dest string) {
	for msg := range channel {
		message := msg.Payload
		log.Info("MESSAGE_RECEIVED", map[string]interface{}{
			"message": message,
			"channel": msg.Channel,
		})
		logApiStatusForDestination(makeAPICallWithRetry(dest, message), dest, msg.Payload)
	}
}

func makeAPICallWithRetry(dest, payload string) error {
	retryDelay := 1 * time.Second
	var err error
	for retry := 0; retry < maxRetries; retry++ {
		err = makeAPICall(dest, payload)
		if err == nil {

			return nil
		}
		log.Error("RETRY_API_CALL_TO_DESTINATION_DUE_TO_FAILURE", map[string]interface{}{
			"destination": dest,
			"payload":     payload,
			"retry_count": retry,
		})
		retryDelay *= 2
	}
	return err
}

func makeAPICall(dest, payload string) error {
	randomNumber := rand.Intn(101)
	if randomNumber < destinationToErrorPercentageMap[dest] {
		return newCustomError(errorServerUnreachable, 400)
	}
	return nil
}

func logApiStatusForDestination(err error, dest, payload string) {
	if err == nil {
		log.Info("SUCCESSFUL_API_CALL_TO_DESTINATION", map[string]interface{}{
			"destination": dest,
			"payload":     payload,
		})
	} else {
		log.Error("API_CALL_TO_DESTINATION_FAILED_AFTER_RETRIES", map[string]interface{}{
			"destination": dest,
			"payload":     payload,
		})
	}
}
