package main

import "fmt"
import "gopkg.in/redis.v5"

func main() {
	fmt.Printf("Hello world")
	ExampleNewClient()
	fmt.Printf("Hello World2")
}

func ExampleNewClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}
