package util

import (
	"io/ioutil"
)

const (
	PRODCONFIG = "configs/config.yaml"
	DEVCONFIG  = "configs/config-dev.yaml"
)

type Config struct {
	Mode   string
	Port   string `yaml:"Internal_Port"`
	DBConn string `yaml:"Database_Connection"`
}

func (s *Config) Load() {
	configfile := DEVCONFIG
	if dev, _ := os.LookupEnv("PRODUCTION"); dev != "" {
		configfile = PRODCONFIG
		s.Mode = "production"
		log.Println("Running in production mode.")
	} else {
		s.Mode = "development"
		log.Println("Running in development mode.")
	}
	content, err := ioutil.ReadFile(configfile)
	check(err, "Error reading config file")
	content = []byte(os.ExpandEnv(string(content)))
	err = yaml.Unmarshal(content, &s)
	check(err, "Error parsing configuration")
}
