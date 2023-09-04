* Basic http server setup
* Setup redis.
* Use redis pubsub to publish all the incoming messages.
* Parallely run an infinite running goroutine which receives messages from the pubsub.
* Mock the behaviour of destinations.
* Retry with exponential backoff
    * Requirement: Delays or failures with event delivery of a single destination should not affect ingestion or other delivery on other destinations.
    * Each destination would have a separate goroutine and a separate redis channel.
* Add Dockerfile, docker-compose
* Add README.md
* Error handling
* Add test cases
* Github actions to run tests.

/v1/proxy

{
    "payload": string
    "user_id": string
    "destinations": []string
}