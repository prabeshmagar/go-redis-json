package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/prabeshmagar/go-redis-json/service"
)

var redisHost = "localhost:6379"

func main() {
	student := service.Student{
		Name: service.Name{
			First:  "Mark",
			Middle: "S",
			Last:   "Pronto",
		},
		Rank: 1,
	}

	// redigo example
	redigoService := service.RedigoExample(redisHost)

	err := redigoService.SetJson("student", student)

	if err != nil {
		fmt.Println("failed to store json in redis with redigo")
	}

	studentJSON := redigoService.GetJson("student")

	readStudent := service.Student{}

	err = json.Unmarshal(studentJSON, &readStudent)
	if err != nil {
		log.Fatalf("Failed to JSON Unmarshal")
	}

	fmt.Printf("Reading data from redis with redigo \n%#v\n", readStudent)

	//go-redis example
	anotherStudent := service.Student{
		Name: service.Name{
			First:  "Prabesh",
			Middle: "",
			Last:   "Magar",
		},
		Rank: 2,
	}

	// Goredis example
	goRedisService := service.GoRedisExample(redisHost)

	err = goRedisService.GoRedisSetJSON("another", anotherStudent)

	if err != nil {
		fmt.Println("failed to store json in redis with goredis")
	}

	goRedisService.GoRedisGetJSON("another")
}
