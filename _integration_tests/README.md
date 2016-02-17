# Functional Testing

Functional testing requires real SMTP and AmazonSES accounts. The following environment variables are required to run the integration tests:

Email settings:
* EMAIL_FROM (email address)
* EMAIL_TO (email address)

Amazon SES:
* AWS_ENDPOINT
* AWS_ACCESS_KEY_ID
* AWS_SECRET_ACCESS_KEY

SMTP:
* SMTP_HOST
* SMTP_PORT
* SMTP_USERNAME
* SMTP_PASSWORD


Once you have all of the necessary environment variables set, you may run the integration test suite by executing the following:

`go test ./_integration_tests/`
