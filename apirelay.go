package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/sirupsen/logrus"

	msgraph "github.com/R2D2Env/go-msgraph"
	"github.com/chrj/smtpd"
)

var graphClient msgraph.GraphClient

func InitMSGraph() {
	// initialize the GraphClient via JSON-Load.
	// Specify JSON-Fields TenantID, ApplicationID and ClientSecret
	fileContents, err := ioutil.ReadFile(*apiCredentialFile)
	if err != nil {
		log.WithFields(logrus.Fields{
			"apiCredentialFile": *apiCredentialFile,
		}).WithError(err).Warn("could not read api credentials file")
	}
	json.Unmarshal(fileContents, &graphClient)
}

func GraphRelay(env smtpd.Envelope) {
	if err := graphClient.SendMailMIME(env.Sender, env.Data); err != nil {
		log.WithFields(
			logrus.Fields{
				"Sender":     env.Sender,
				"Recipients": env.Recipients,
				"Body":       string(env.Data),
			}).WithError(err).Warn("send mail error")
	}
}
