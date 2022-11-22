package psql

import (
	"context"
	"log"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PsqlStore struct {
	db *gorm.DB
}

var store *PsqlStore
var once sync.Once

func CreateStore(uri string, dbName string, timeout time.Duration) (*PsqlStore, error) {
	db, err := gorm.Open(postgres.Open(uri+"/"+dbName), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return &PsqlStore{db}, nil
}

func (store *PsqlStore) Insert(ctx context.Context, model interface{}, data interface{}) error {
	log.Println("Try to insert")

	store.db.Model(model).Create(data)

	log.Println("Success to insert")

	return nil
}

func (store *PsqlStore) Update(ctx context.Context, model interface{}, query interface{}, update interface{}) error {
	ret := store.db.Model(model).UpdateColumns(update).Where(query)

	log.Println(ret)

	return nil
}

func (store *PsqlStore) FindAll(ctx context.Context, model interface{}, query interface{}, result interface{}) error {
	store.db.Model(model).Find(result, query)
	return nil
}

func (store *PsqlStore) FindOne(
	ctx context.Context,
	model interface{},
	query interface{},
	result interface{}) error {
	store.db.First(result, query)

	return nil
}

func (store *PsqlStore) Delete(
	ctx context.Context,
	model interface{},
	query interface{}) error {
	store.db.Model(model).Delete(query, query)

	return nil
}

func (store *PsqlStore) Migrate(model interface{}) error {
	return store.db.AutoMigrate(model)
}
