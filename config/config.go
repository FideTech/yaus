package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	yaml "gopkg.in/yaml.v2"

	"github.com/FideTech/yaus/models"
)

//Config contains a reference to the configuration
var Config *config

type config struct {
	System    systemConfig    `yaml:"system" json:"system"`
	Hardcoded hardcodedConfig `yaml:"hardcoded" json:"hardcoded"`
}

type systemConfig struct {
	Port    int    `yaml:"port" json:"port"`
	BaseURL string `yaml:"baseUrl" json:"baseUrl"`
}

type hardcodedConfig struct {
	Error []models.ShortLink `yaml:"error" json:"error"`
	Info  []models.ShortLink `yaml:"info" json:"info"`
}

func (c *config) load(filePath string) error {
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(yamlFile, c); err != nil {
		return err
	}

	return nil
}

func (c *config) verifySettings() error {
	if c.System.BaseURL == "" {
		return errors.New("invalid system.baseUrl, can not be empty")
	}

	if strings.HasSuffix(c.System.BaseURL, "/") {
		return errors.New("system.baseUrl must NOT end with '/'")
	}

	return nil
}

//HTTPHandler outputs the config as a json file
func (c *config) HTTPHandler(w http.ResponseWriter) {
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
