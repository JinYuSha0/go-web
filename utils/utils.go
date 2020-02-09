package utils

import (
	"go-web/models"
	"log"

	"github.com/BurntSushi/toml"
)

var config *models.Config

// GetConfig 获取配置内容
func GetConfig() *models.Config {
	if config == nil {
		path := "./config/config.toml"
		if _, err := toml.DecodeFile(path, &config); err != nil {
			log.Fatal(err)
		}
	}

	return config
}
