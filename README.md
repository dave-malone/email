# email
A simple Go (golang) library for sending email across a variety of existing services.

Currently, there are three supported implementations:

1. noopSender - simply logs received messages
1. amazonSESSender - sends email via Amazon SES
1. smtpSender - sends email over SMTP

You to choose an implementation, you will set the value of the factory method for the implementation you want to use.

For example, to use the AmazonSESSender implementation:

`NewSenderFactory = NewAmazonSESSender("endpoint-url", "access-key-id", "secret-access-key")`

Then, you can initialize an email.Sender anywhere that requires it using the factory method:

`sender := NewSenderFactory()`
