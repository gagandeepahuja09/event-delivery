* Basic http server setup (d)
* Setup redis. (d)
* Use redis pubsub to publish all the incoming messages. (d)
* Parallely run an infinite running goroutine which receives messages from the pubsub. (d)
* Mock the behaviour of destinations. (d)
* Retry with exponential backoff (d)
    * Requirement: Delays or failures with event delivery of a single destination should not affect ingestion or other delivery on other destinations.
    * Each destination would have a separate goroutine and a separate redis channel.
* Add Dockerfile, docker-compose (d)
* Add README.md (d)
* Error handling for API request (d)
    * Throw error if destination is not supported. (d)
* Providing a better success response. (d)
* Add test cases
* Github actions to run tests.
* Moving to using worker pool.
* Using constants for logs.
* Providing the success and failure info in a better way: probably some report file. 
* Move to using configs instead of hardcoding.

/v1/proxy

{
    "payload": string
    "user_id": string
    "destinations": []string
}