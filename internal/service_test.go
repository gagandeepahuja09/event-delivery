package internal

import (
	"context"
	"testing"

	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

var testPubsub *redis.PubSub

func TestMain(m *testing.M) {
	initRedis()
}

func TestHandleProxyRequestSalesforce(t *testing.T) {
	testPubsub = Rdb.Subscribe("salesforce_channel")
	err := handleProxyRequest(context.TODO(), proxyRequest{
		Payload:      "test_payload",
		UserId:       "user_001",
		Destinations: []string{"salesforce"},
	})
	assert.NoError(t, err)
	assertRedisMessagePublished(t)
}

func TestHandleProxyRequestUnsupportedDestination(t *testing.T) {
	err := handleProxyRequest(context.TODO(), proxyRequest{
		Payload:      "test_payload",
		UserId:       "user_001",
		Destinations: []string{"salesforce2"},
	})
	assert.EqualError(t, err, "destinations array should contain values only from the following supported destinations list: [salesforce snowflake marketo google_analytics]")
}

func assertRedisMessagePublished(t *testing.T) {
	defer testPubsub.Close()

	msg, err := testPubsub.ReceiveMessage()
	assert.Equal(t, "test_payload", msg.Payload)
	assert.NoError(t, err)
}
