# email
A simple Go (golang) library for sending email across a variety of existing services.

[![wercker status](https://app.wercker.com/status/dda3c4501ded6999dd0a784fddfd3041/m "wercker status")](https://app.wercker.com/project/bykey/dda3c4501ded6999dd0a784fddfd3041)

Currently, there are three supported implementations:

1. noopSender - simply logs received messages
1. amazonSESSender - sends email via Amazon SES
1. smtpSender - sends email over SMTP

You to choose an implementation, you will set the value of the factory method for the implementation you want to use.

For example, to use the AmazonSESSender implementation:

`NewSenderFactory = NewAmazonSESSender("endpoint-url", "access-key-id", "secret-access-key")`

Then, you can initialize an email.Sender anywhere that requires it using the factory method:

`sender := NewSenderFactory()`


## Email Template Support

Currently, there are two supported MessageBody implementations. simpleMessageBody is simply a wrapper around a string and can be used to send very basic email messages.

The second implementation has been designed for HTML templates saved in files. This implementation is a light wrapper around Go's [html/template](https://golang.org/pkg/html/template/) package. An example of how to use this implementation is found in the [integration tests](_integration_tests/amazonsesemailsender_test.go)
