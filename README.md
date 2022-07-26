# Bitcoin price API

Available API:

- **/rate** -- Get current price for BTCUAH from api.binance.com
- **/subscribe** -- Post request which adds sent "email" to example.csv file. If the email address is already present in the example.csv, 409 error is thrown.
- **/sendEmails** -- Post request which send the current BTCUAH price to all the subscribers from example.csv file. This method uses Mailgun for the mail delivery service.

Example usage:

 - `curl http://localhost:12321/rate`
 - `curl http://localhost:12321/subscribe --include --header "Content-Type: application/json" --request "POST" --data '{"email": "new_email_address@gmail.org"}'`
 - `curl http://localhost:12321/sendEmails --request "POST"`

TODO:
 - Save Mailgun key to environmental variable and use it more responsibly, but for tasks unrelated to production standards leave it be.
 - Swap example.csv for a relational database to increase scalability.
 - Decrease Docker container size with the staged build. 
 - Create documentation for the project
 - Add tests.
