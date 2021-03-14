package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	yaml "gopkg.in/yaml.v2"

	"github.com/FideTech/yaus/models"
	"github.com/FideTech/yaus/utils"
)

//Config contains a reference to the configuration
var Config *config

type config struct {
	System    systemConfig    `yaml:"system" json:"system"`
	Hardcoded hardcodedConfig `yaml:"hardcoded" json:"hardcoded"`
}

type systemConfig struct {
	Router   routerConfig   `yaml:"router" json:"router"`
	Database databaseConfig `yaml:"database" json:"database"`
}

type routerConfig struct {
	Port int `yaml:"port" json:"port"`
}

type databaseConfig struct {
	Path string `yaml:"path" json:"path"`
}

type hardcodedConfig struct {
	Error []models.ShortLink `yaml:"error" json:"error"`
	Info  []models.ShortLink `yaml:"info" json:"info"`

	Errors map[string]models.ShortLink `yaml:"-" json:"-"`
	Infos  map[string]models.ShortLink `yaml:"-" json:"-"`
}

func (c *config) load(filePath string) error {
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(yamlFile, c); err != nil {
		return err
	}

	return c.loadHardcodedValues()
}

func (c *config) loadHardcodedValues() error {
	c.Hardcoded.Errors = map[string]models.ShortLink{}
	for index, e := range c.Hardcoded.Error {
		if e.Key == "" && e.URL == "" {
			log.Printf("hardcoded error short link at position #%d is not valid (missing key and url)", index+1)
			continue
		}

		if e.Key == "" {
			return fmt.Errorf("hardcoded short links require a non-empty 'key'. check position #%d of the 'error' short links", index+1)
		}

		if _, found := c.Hardcoded.Errors[e.Key]; found {
			return fmt.Errorf("duplicate hardcoded short 'error' short link found at position #%d with the key \"%s\"", index+1, e.Key)
		}

		if e.URL == "" {
			return fmt.Errorf("hardcoded short links require a non-empty 'url'. check position #%d (\"%s\") of the 'error' short links", index+1, e.Key)
		}

		if !utils.IsValidUrl(e.URL) {
			return fmt.Errorf("the hardcoded 'error' short link \"%s\" (%s) is an invalid url", e.URL, e.Key)
		}

		c.Hardcoded.Errors[e.Key] = e
	}

	c.Hardcoded.Infos = map[string]models.ShortLink{}
	for index, i := range c.Hardcoded.Info {
		if i.Key == "" && i.URL == "" {
			log.Printf("hardcoded info short link at position #%d is not valid (missing key and url)", index+1)
			continue
		}

		if i.Key == "" {
			return fmt.Errorf("hardcoded short links require a non-empty 'key'. check position #%d of the 'info' short links", index+1)
		}

		if _, found := c.Hardcoded.Infos[i.Key]; found {
			return fmt.Errorf("duplicate hardcoded short 'info' short link found at position #%d with the key \"%s\"", index+1, i.Key)
		}

		if i.URL == "" {
			return fmt.Errorf("hardcoded short links require a non-empty 'url'. check position #%d (\"%s\") of the 'info' short links", index+1, i.Key)
		}

		if !utils.IsValidUrl(i.URL) {
			return fmt.Errorf("the hardcoded 'info' short link \"%s\" (%s) is an invalid url", i.URL, i.Key)
		}

		c.Hardcoded.Infos[i.Key] = i
	}

	return nil
}

func (c *config) verifySettings() error {
	//TODO: check if the router port is open and usable
	//TODO: check if the database path is accessible and writable

	return nil
}

//HTTPHandler outputs the config as a json file
func (c *config) HTTPHandler(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

//Load tries to load the configuration file
func Load(filePath string) error {
	Config = new(config)

	if err := Config.load(filePath); err != nil {
		return err
	}

	return Config.verifySettings()
}
