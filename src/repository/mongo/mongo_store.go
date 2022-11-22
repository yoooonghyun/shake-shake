package mongo

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"shake-shake/src/utils/type_utils"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type MongoStore struct {
	db *mongo.Database
}

type BsonMarshaler interface {
	ToBsonM() bson.M
}

func CreateStore(uri string, dbName string, timeout time.Duration) (*MongoStore, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, err
	}
	return &MongoStore{client.Database(dbName)}, nil
}

func (store *MongoStore) Insert(ctx context.Context, model interface{}, data interface{}) error {
	collection := type_utils.GetStructNameSnake(model)
	log.Println("Try to insert")
	_, err := store.db.Collection(collection).InsertOne(ctx, data)
	if err != nil {
		log.Println("Failed to insert")
		return err
	}

	log.Println("Success to insert")

	return nil
}

func (store *MongoStore) Update(ctx context.Context, model interface{}, query interface{}, update interface{}) error {
	collection := type_utils.GetStructNameSnake(model)
	if _, err := store.db.Collection(collection).UpdateOne(ctx, query, update); err != nil {
		return err
	}

	return nil
}

func (store *MongoStore) FindAll(ctx context.Context, model interface{}, query interface{}, result interface{}) error {
	collection := type_utils.GetStructNameSnake(model)
	where, err := queryToBsonM(query)

	log.Print(result)
	log.Print(where)

	if err != nil {
		return err
	}
	cur, err := store.db.Collection(collection).Find(ctx, where)
	if err != nil {
		return fmt.Errorf("finding collection: %w", err)
	}
	defer func() {
		_ = cur.Close(ctx)
	}()

	if err := cur.All(ctx, result); err != nil {
		return fmt.Errorf("fetching result: %w", err)
	}

	return cur.Err()
}

func (store *MongoStore) FindOne(
	ctx context.Context, model interface{},
	query interface{},
	result interface{}) error {
	collection := type_utils.GetStructNameSnake(model)
	where, err := queryToBsonM(query)

	if err != nil {
		return err
	}

	return store.db.Collection(collection).
		FindOne(
			ctx,
			where).Decode(result)
}

func (store *MongoStore) Delete(ctx context.Context, model interface{}, query interface{}) error {
	collection := type_utils.GetStructNameSnake(model)
	where, err := queryToBsonM(query)

	if err != nil {
		return err
	}

	_, err = store.db.Collection(collection).DeleteMany(ctx, where)

	if err != nil {
		return err
	}
	return nil
}

func (store *MongoStore) Migrate(model interface{}) error {
	return nil
}

func queryToBsonM(query interface{}) (bson.M, error) {
	value := reflect.ValueOf(query)
	if value.Kind() != reflect.Struct {
		return nil, fmt.Errorf("invalid model type")
	}

	bsonM := bson.M{}
	typeOfModel := value.Type()

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		filedType := field.Kind()

		if filedType == reflect.Interface || filedType == reflect.Struct {
			continue
		}

		columnName := typeOfModel.Field(i).Tag.Get("bson")

		if columnName[0] < 'a' || columnName[0] > 'z' {
			continue
		}

		val := field.Interface()

		if reflect.DeepEqual(val, reflect.Zero(field.Type()).Interface()) {
			continue
		}

		bsonM[columnName] = val
	}

	return bsonM, nil
}
