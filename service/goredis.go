package service

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/nitishm/go-rejson/v4"
)

type GoRedis struct {
	Client  *redis.Client
	Handler *rejson.Handler
}

func GoRedisExample(url string) *GoRedis {
	client := redis.NewClient(&redis.Options{
		Addr: url,
	})

	rh := rejson.NewReJSONHandler()
	rh.SetGoRedisClient(client)

	return &GoRedis{
		Client:  client,
		Handler: rh,
	}
}

func (r GoRedis) GoRedisSetJSON(key string, value interface{}) error {
	res, err := r.Handler.JSONSet(key, ".", value)
	if err != nil {
		log.Fatalf("Failed to JSONSet with go redis")
		return err
	}

	if res.(string) == "OK" {
		fmt.Printf("GoRedis Success: %s\n", res)
	}
	return nil
}

func (r GoRedis) GoRedisGetJSON(key string) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	ctx := context.Background()

	val, err := r.Client.Do(ctx, "JSON.GET", key, ".name").Result()

	if err != nil {
		fmt.Println("unable to get data from redis")
	}

	er := enc.Encode(val)

	if er != nil {
		panic("unable to encode value")
	}

	fmt.Printf("Reading data from redis with goredis : %#v\n", val)
}
