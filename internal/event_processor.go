package internal

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

var pubsub redis.PubSub

func ProcessEvents() {
	pubsub := Rdb.Subscribe(eventsChannel)

	go func() {
		defer pubsub.Close()
		for {
			msg, err := pubsub.ReceiveMessage()
			if err != nil {
				log.Fatal(err)
			}

			log.Infof("MESSAGE_RECEIVED", msg)
		}
	}()
}
