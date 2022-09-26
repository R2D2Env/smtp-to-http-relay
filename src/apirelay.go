package main

import (
	"encoding/json"
  "fmt"
	"strings"
  "os"
	//"io/ioutil"

	"github.com/sirupsen/logrus"

	msgraph "github.com/R2D2Env/go-msgraph"
	"github.com/chrj/smtpd"
)

var graphClients map[string]msgraph.GraphClient

type SMTPUser struct {
  Address    string `json: "address"`
  Hashedpass string `json: "hashedpass"`
  Username   string `json: "username"`
}

type GraphClientSMTPAuth struct {
  Users []SMTPUser `json: "users"`
}

func InitMSGraph() {
	// initialize the GraphClient via JSON-Load.
	// Specify JSON-Fields TenantID, ApplicationID and ClientSecret
	graphClients = make(map[string]msgraph.GraphClient)
	files := strings.Split(*apiCredentialFiles, ",")
	for _, file := range files {
		fileContents, err := os.ReadFile(file)
    fc := string(fileContents)
    log.WithFields(logrus.Fields{
      "fileContents": fc,
    }).Info("File contents")
		if err != nil {
			log.WithFields(logrus.Fields{
				"apiCredentialFiles": *apiCredentialFiles,
			}).WithError(err).Warn("could not read api credentials file")
		}

		// Unmarshel the file contents into objects
		// use the 0'th sender address in the auth credentials file as the key
		// for the map
		smtpmap := GraphClientSMTPAuth{}
    fmt.Println("%+v\n", smtpmap)
		gc := msgraph.GraphClient{}
		json.Unmarshal(fileContents, &smtpmap)
    fmt.Println("%+v\n", smtpmap)
		json.Unmarshal(fileContents, &gc)
		graphClients[smtpmap.Users[0].address] = gc
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
