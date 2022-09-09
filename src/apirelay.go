package main

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/sirupsen/logrus"

	msgraph "github.com/R2D2Env/go-msgraph"
	"github.com/chrj/smtpd"
)

var graphClients map[string]msgraph.GraphClient

type SMTPUser struct {
	address    string
	hashedpass string
	username   string
}

type GraphClientSMTPAuth struct {
	users []SMTPUser
}

func InitMSGraph() {
	// initialize the GraphClient via JSON-Load.
	// Specify JSON-Fields TenantID, ApplicationID and ClientSecret
	graphClients = make(map[string]msgraph.GraphClient)
	files := strings.Split(*apiCredentialFiles, ",")
	for _, file := range files {
		fileContents, err := ioutil.ReadFile(file)
		if err != nil {
			log.WithFields(logrus.Fields{
				"apiCredentialFiles": *apiCredentialFiles,
			}).WithError(err).Warn("could not read api credentials file")
		}

		// Unmarshel the file contents into objects
		// use the 0'th sender address in the auth credentials file as the key
		// for the map
		smtpmap := GraphClientSMTPAuth{}
		gc := msgraph.GraphClient{}
		json.Unmarshal(fileContents, &smtpmap)
		json.Unmarshal(fileContents, &gc)
		graphClients[smtpmap.users[0].address] = gc
	}
}

func GraphRelay(env smtpd.Envelope) {
	if gc, ok := graphClients[env.Sender]; ok {
		if err := gc.SendMailMIME(env.Sender, env.Data); err != nil {
			log.WithFields(
				logrus.Fields{
					"Sender":     env.Sender,
					"Recipients": env.Recipients,
					"Body":       string(env.Data),
				}).WithError(err).Warn("send mail error")
		}
	} else {
		log.WithFields(
			logrus.Fields{
				"Sender":     env.Sender,
				"Recipients": env.Recipients,
				"Body":       string(env.Data),
			}).Warn("Sender address not mapped to MS Graph account")
	}
}
