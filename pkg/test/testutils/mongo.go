package testutils

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoHost   = "localhost"
	mongoClient *MongoClient
	mongoOnce   sync.Once
)

func init() {
	host := os.Getenv("TEST_MONGODB_HOST")
	if host != "" {
		mongoHost = host
	}
	rand.Seed(time.Now().Unix())
}

type MongoClient struct {
	*mongo.Client
}

func (c *MongoClient) Tx(t *testing.T, fn func(db *mongo.Database)) {
	dbName := "test-" + time.Now().Format("20060102-150405") + "-" + strconv.Itoa(rand.Intn(10000-1000)+1000)
	db := c.Database(dbName)
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		err := db.Drop(ctx)
		cancel()
		if err != nil {
			t.Errorf("drop test database: %s fail, %v", dbName, err)
		}
	}()
	fn(db)
}

func GetMongoClient(t *testing.T) *MongoClient {
	mongoOnce.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:27017", mongoHost)))
		if err != nil {
			t.Skipf("connect mongo fail: %v", err)
		}
		err = client.Ping(ctx, nil)
		if err != nil {
			t.Skipf("connect mongo fail: %v", err)
		}
		mongoClient = &MongoClient{client}
	})

	if mongoClient == nil {
		t.Skip("connect mongo fail")
	}
	return mongoClient
}
