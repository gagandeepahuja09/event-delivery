## Event-Delivery App

To get started with our server, follow these simple steps:

1. Run the following command to set up the server:
   ```bash
   docker-compose up --build
   ```
   This will start the server and make it available on port 9410.

2. Now that the server is up and running, you can send event payloads along with their destinations using the following route. We've provided a sample `curl` request for your convenience:

   ```bash
   curl --location --request POST 'http://localhost:9410/v1/proxy' \
   --header 'Content-Type: application/json' \
   --data-raw '{
       "payload": "Hello123",
       "user_id": "123456",
       "destinations": [ "salesforce", "marketo", "c", "google_analytics"]
   }'
   ```

   Replace the sample payload, user ID, and destinations with your own data as needed.

3. **Supported Destinations:**

   - Salesforce: `"salesforce"`
   - Snowflake: `"snowflake"`
   - Marketo: `"marketo"`
   - Google Analytics: `"google_analytics"`

4. **Logs and Their Meanings:**
    The logs will be available as the output of the docker-compose command, providing you with insights into the server's operations.

   - `SUCCESSFUL_API_CALL_TO_DESTINATION`: This log indicates that the API call to the specified destination was successful. It means your data has been successfully delivered to the destination without any issues.

   - `API_CALL_TO_DESTINATION_FAILED_AFTER_RETRIES`: This log signifies that the API call to the destination initially failed but was retried multiple times. Unfortunately, even after retries, the call still failed. This generally happens due to the downstream destinations facing any sort of outage. You can contact the downstream support team. For any help, our support team is also available.

That's it! You're all set to interact with the server, understand the logs, and send event data to your specified destinations. Happy coding! 🚀