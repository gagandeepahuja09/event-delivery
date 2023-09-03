* Basic http server setup
* Setup redis.
* Use redis pubsub to publish all the incoming messages.
* Parallely run an infinite running goroutine which receives messages from the pubsub.
* Mock the behaviour of destinations.
* Retry with exponential backoff

/v1/proxy

{
    "payload": string
    "user_id": string
    "destinations": []string
}