package infra

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var loader *ConfigLodaer
var once sync.Once

type ConfigKey string

const (
	ConfigKeyDbUrl  ConfigKey = "DB_URL"
	ConfigKeyDbName ConfigKey = "DB_NAME"
)

type ConfigLodaer struct {
}

func GetConfigLoader() (*ConfigLodaer, error) {
	var err error
	once.Do(func() {
		err = loadConfig()
	})

	if err != nil {
		return nil, err
	}

	return loader, nil
}

func (l *ConfigLodaer) Get(key ConfigKey) string {
	str := string(key)

	return os.Getenv(str)
}

func loadConfig() error {
	err := godotenv.Load()

	if err != nil {
		return err
	}

	loader = &ConfigLodaer{}

	return nil
}
