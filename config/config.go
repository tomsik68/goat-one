package config

import (
	"io/ioutil"
	"time"
	"github.com/go-yaml/yaml"
)

type Configuration struct {
	DateFrom time.Time
	DateTo time.Time
	Period time.Duration
	IncludeGroups []string
	ExcludeGroups []string
	GroupsFile string
	GoatServerIP string
	OneKey string
	OneEndpoint string
	ClientIdentifier string
}

func Load(filename string, c *Configuration) error {

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return err
	}
	return nil
}
