# smtprelay

[![Go Report Card](https://goreportcard.com/badge/github.com/decke/smtprelay)](https://goreportcard.com/report/github.com/decke/smtprelay)

Simple Golang based SMTP relay/proxy server that accepts mail via SMTP
and forwards it directly to another SMTP server, and/or forwards it to Microsoft Graph.


## Why another SMTP server?

Outgoing mails are usually send via SMTP to an MTA (Mail Transfer Agent)
which is one of Postfix, Exim, Sendmail or OpenSMTPD on UNIX/Linux in most
cases. You really don't want to setup and maintain any of those full blown
kitchensinks yourself because they are complex, fragile and hard to
configure.

Many applications require an SMTP endpoint to configure email notifications, but R2D2E enterprise does not provide a well-supported SMTP endpoint to connect to, instead providing only a Microsoft Graph HTTP endpoint. This application relays incoming SMTP messages to the Microsoft Graph API via the Microsoft Graph /sendMail endpoint and a Microsoft Graph API SDK.


## Main features

* Supports SMTPS/TLS (465), STARTTLS (587) and unencrypted SMTP (25)
* Checks for sender, receiver, client IP
* Authentication support with file (LOGIN, PLAIN)
* Enforce encryption for authentication
* Forwards all mail to:
    * a smarthost (GMail, MailGun or any other SMTP server)
    * Microsoft Graph sendMail API
* Small codebase
* IPv6 support

## Use

1. Configure your application settings in smtprelay.ini
2. Run with a `--config` flag. For example, in your source folder:

```sh
go build .
./smtprelay --config ./smtpconfig.ini
```