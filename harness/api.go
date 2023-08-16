package harness

import (
	"log"
	"os"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type APIRequest struct {
	BaseURL string
	Client  *resty.Client
}

type Config struct {
	AccountIdentifier string   `yaml:"accountIdentifier"`
	OrgIdentifier     string   `yaml:"orgIdentifier"`
	ProjectIdentifier string   `yaml:"projectIdentifier"`
	TargetIdentifier  []string `yaml:"targetIdentifier"`
	ApiKey            string   `yaml:"apiKey"`
	Names             []string `yaml:"names"`
}

func GetAccountIDFromAPIKey(apiKey string) string {
	accountId := strings.Split(apiKey, ".")[1]
	if accountId == "" {
		logrus.Fatal("Failed to get account ID from API key")
	}

	return accountId
}

func (c *Config) ReadConfig(filepath string) *Config {
	yamlFile, err := os.ReadFile(filepath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
