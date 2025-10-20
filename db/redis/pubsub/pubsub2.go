package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func main() {
	// // Establish a connection to Redis
	// conn, err := redis.Dial("tcp", "localhost:6379",redis.DialReadTimeout(time.Second))
	// if err != nil {
	// 	log.Fatalf("Unable to connect to redis: %v", err)
	// }
	// defer conn.Close()

	// Create a Redis connection pool
	pool = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "localhost:6379")
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	defer pool.Close()

	var err error


	psc := redis.PubSubConn{Conn: pool.Get()}
	// defer psc.Close()

	// Set the notify-keyspace-events config
	_, err = psc.Conn.Do("CONFIG", "SET", "notify-keyspace-events", "KEA")
	if err != nil {
		log.Fatalf("Unable to set redis config to notify expired events: %v", err)
	}

	// Subscribe to the keyevent notifications
	err = psc.PSubscribe("__keyevent@*__:expired")
	if err != nil {
		log.Fatalf("Unable to subscribe: %v", err)
	}

	// Start a goroutine to handle messages
	go func() {
		for {
			switch msg := psc.ReceiveWithTimeout(0).(type) {
			case redis.Message:
				key := string(msg.Data)
				fmt.Printf("Received message: [%v][%v]\n", msg.Channel, key)
			case error:
				log.Fatalf("Error: %v", msg)
			}
		}
	}()

	// Keep the program running to receive messages
	for {
		time.Sleep(1 * time.Second)
	}
}

// // terminal 1: start docker container for redis with port exposed
// docker run -p 6379:6379 redis:alpine &

// // terminal 1: connect to redis and set the config to notify expired events, and then start monitoring (i.e. print all the commands redis receives)
// redis-cli -h localhost -p 6379
// monitor

// // terminal 2: start the go code to listen to expired pubsub messages
// go run pubsub2.go

// // terminal 3:
// redis-cli -h localhost -p 6379
// set name srini EX 5
