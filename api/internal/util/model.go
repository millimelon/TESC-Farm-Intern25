package util

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Mode   string
	Host   string
	Port   string `yaml:"Internal_Port"`
	DBConn string `yaml:"Database_Connection"`
}

func (s *Config) Load(prodconf string, devconf string) {
	configfile := devconf
	if dev, _ := os.LookupEnv("PRODUCTION"); dev != "" {
		configfile = prodconf
		s.Mode = "production"
		s.Host = "127.0.0.1"
		log.Println("Running in production mode.")
	} else {
		s.Mode = "development"
		s.Host = "0.0.0.0"
		log.Println("Running in development mode.")
	}
	content, err := ioutil.ReadFile(configfile)
	Check(err, "Error reading config file")
	content = []byte(os.ExpandEnv(string(content)))
	err = yaml.Unmarshal(content, &s)
	Check(err, "Error parsing configuration")
}
