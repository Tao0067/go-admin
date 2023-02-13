package config

import (
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"sync"
)

var once sync.Once
var conf *Config

func LoadConfig() *Config {
	once.Do(func() {
		file, err := os.ReadFile("conf/conf.yaml")
		if err != nil {
			log.Fatalf("Fail to read 'conf/conf.yaml': %v", err)
		}

		conf = &Config{}
		err = yaml.Unmarshal(file, conf)

		if err != nil {
			log.Fatalf("Fail to parse 'conf/conf.yaml': %v", err)
		}
	})
	//fmt.Printf("db: %+v:%+s\n", conf.Server.Host, conf.Server.Port)
	return conf
}

func ViperConfig() *Config {
	v := viper.New()
	v.AddConfigPath("conf")
	v.SetConfigFile("conf")
	v.SetConfigType("yaml")

	v.Unmarshal(&conf)
	return conf
	// 监听文件
	//v.WatchConfig()
	//v.OnConfigChange(func(e fsnotify.Event) {
	//	v.Unmarshal(&conf)
	//	conf.Looca
	//})

}
