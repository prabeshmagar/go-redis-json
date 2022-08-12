package service

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/nitishm/go-rejson/v4"
)

// Name - student name
type Name struct {
	First  string `json:"first,omitempty"`
	Middle string `json:"middle,omitempty"`
	Last   string `json:"last,omitempty"`
}

// Student - student object
type Student struct {
	Name Name `json:"name,omitempty"`
	Rank int  `json:"rank,omitempty"`
}

type RedigoService struct {
	Client  redis.Conn
	Handler *rejson.Handler
}

func RedigoExample(url string) RedigoService {
	conn, err := redis.Dial("tcp", url)
	rh := rejson.NewReJSONHandler()

	if err != nil {
		log.Fatalf("Failed to connect to redis-server @ %s", url)
		panic(err)
	}

	rh.SetRedigoClient(conn)

	return RedigoService{
		Client:  conn,
		Handler: rh,
	}
}

func (r RedigoService) SetJson(key string, value interface{}) error {
	res, err := r.Handler.JSONSet(key, ".", value)
	if err != nil {
		log.Fatalf("Failed to JSONSet")
		return err
	}

	if res.(string) == "OK" {
		fmt.Printf("Redigo Success: %s\n", res)
	}
	return nil
}

func (r RedigoService) GetJson(key string) []byte {

	studentJSON, err := redis.Bytes(r.Handler.JSONGet(key, "."))
	if err != nil {
		log.Fatalf("Failed to JSONGet")
		return nil
	}

	return studentJSON
}
