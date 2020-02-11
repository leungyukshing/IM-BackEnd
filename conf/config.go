package conf

import (
	"gopkg.in/yaml.v2"
	"os"
)

var Conf *Configuration

const path = "conf.yml"

type Configuration struct {
	MysqlDsn    string `yaml:"mysqlDsn"`
	GraylogAddr string `yaml:"graylog"`
}

func InitConfig() error {
	Conf = &Configuration{}
	if f, err := os.Open(path); err != nil {
		return err
	} else {
		yaml.NewDecoder(f).Decode(Conf)
	}
	return nil
}
