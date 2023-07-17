package main

import (
	"flag"
	"os"
	"time"

	"github.com/aleksa11010/HarnessTriggerAPI/harness"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.Stamp,
	})
	logrus.SetOutput(os.Stdout)
	if os.Getenv("DEBUG") == "true" {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
}

func main() {
	configFile := flag.String("config", "", "Path to the Account Config YAML file")
	triggerFile := flag.String("trigger", "", "Path to the Trigger YAML file")

	flag.Parse()

	apiClient := harness.APIRequest{
		BaseURL: harness.BaseURL,
		Client:  resty.New(),
	}

	accountDetails := harness.Config{}
	accountDetails.ReadConfig(*configFile)

	trigger := harness.ReadTriggerYaml(*triggerFile, &accountDetails)

	resp, err := apiClient.Client.R().
		SetHeader("x-api-key", accountDetails.ApiKey).
		SetHeader("Content-Type", "application/json").
		SetQueryParams(map[string]string{
			"accountIdentifier": accountDetails.AccountIdentifier,
			"orgIdentifier":     accountDetails.OrgIdentifier,
			"projectIdentifier": accountDetails.ProjectIdentifier,
			"targetIdentifier":  accountDetails.TargetIdentifier,
		}).
		SetBody(trigger).
		Post(apiClient.BaseURL + "/pipeline/api/triggers")

	if err != nil {
		logrus.Error(err)
	}

	logrus.Info(resp)
}
