package store

import (
	"context"
	"net/url"
	"shake-shake/src/infra"
	"shake-shake/src/repository/mongo"
	"shake-shake/src/repository/psql"
	"sync"
)

var s Store
var once sync.Once

type Store interface {
	Insert(ctx context.Context, model interface{}, data interface{}) error
	Update(ctx context.Context, model interface{}, query interface{}, update interface{}) error
	FindAll(ctx context.Context, model interface{}, query interface{}, result interface{}) error
	FindOne(
		ctx context.Context,
		model interface{},
		query interface{},
		result interface{}) error
	Delete(
		ctx context.Context,
		model interface{},
		query interface{},
	) error
	Migrate(model interface{}) error
}

func GetStore(model interface{}) (Store, error) {
	var err error
	once.Do(func() {
		err = loadStore()
	})

	if err != nil {
		return nil, err
	}

	err = s.Migrate(model)

	if err != nil {
		return nil, err
	}

	return s, nil
}

func loadStore() error {
	config, err := infra.GetConfigLoader()

	if err != nil {
		return err
	}

	var dbUrl = config.Get(infra.ConfigKeyDbUrl)
	var dbName = config.Get(infra.ConfigKeyDbName)
	parsedURL, err := url.Parse(dbUrl)

	if err != nil {
		return err
	}

	switch parsedURL.Scheme {
	case "mongodb":
		s, err = mongo.CreateStore(dbUrl, dbName, 3000)
	case "postgres":
		s, err = psql.CreateStore(dbUrl, dbName, 3000)
	}

	if err != nil {
		return err
	}

	return nil
}
