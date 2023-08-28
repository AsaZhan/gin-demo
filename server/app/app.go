package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
)

type AppConfig struct {
	ServerConfig     ServerConfig
	DataSourceConfig DataSourceConfig
}

type ServerConfig struct {
	Port int
}

type DataSourceConfig struct {
	Username string
	Password string
	Database string
	Port     int
	IP       string
}

var (
	Cfg       *AppConfig
	configMux sync.RWMutex
)

func Config() *AppConfig {
	return Cfg
}

func InitConfigJSON(filePath string) error {
	var config AppConfig

	if file, err := ioutil.ReadFile(filePath); err != nil {
		log.Fatal("配置文件读取错误,找不到配置文件", err)
		return err
	} else {
		if err = json.Unmarshal(file, &config); err != nil {
			log.Fatal("配置文件转换错误", err)
			return err
		}
	}

	configMux.Lock()
	Cfg = &config
	configMux.Unlock()

	return nil
}
