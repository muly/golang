package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx := context.Background()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Unable to connect to redis: %v", err)
	}

	if err := rdb.ConfigSet(ctx, "notify-keyspace-events", "KEA").Err(); err != nil { // eqivalent to running this on redis cli: CONFIG SET notify-keyspace-events Ex
		log.Fatalf("Unable to set redis config to notify expired events: %v", err)
	}

	/*
		K (Keyspace Events):
			Keyspace events are notifications about what's happening to the database itself (e.g., a key was set, a key was deleted, but the message sent to the pub/sub will have the key name as the channel.

		E (Keyevent Events):
			This flag, when present, indicates that you want to receive keyevent notifications.
			Keyevent notifications are events related to specific operations performed on keys (e.g., a key is set, deleted, expired, etc.).
			Without the E flag, you won't get any keyevent notifications, even if you specify other event types in the value.
			E must be set if you want to be notified of key events.

		A (All Events)
			This flag is a combination of all the flags (g$lshztxek).

		x (Expired Event):
			This flag, when present, indicates that you specifically want to receive notifications about expired keys.
			An "expired" event occurs when a key with a set time-to-live (TTL) reaches its expiration time and is automatically deleted by Redis.
			If you didn't include x but you did include E, you would get keyevent notifications but not the expired event specifically. You might get set or delete events, but not expired events.

	*/

	// Subscribe to the keyevent notifications
	pubsub := rdb.PSubscribe(ctx, "__keyevent@*__:expired")
	defer pubsub.Close()

	// Get the channel for receiving messages
	ch := pubsub.Channel()

	pubsub.Receive(ctx)


	// Start a goroutine to handle messages
	go func() {
		for msg := range ch {
			key := msg.Payload
			fmt.Printf("Received message: [%v][%v]", msg.Channel, key)
			value, err := rdb.Get(ctx, key).Result()
			if err != nil {
				log.Printf("Error getting value for expired key: %v\n", err)
				continue // skip to the next key
			}
			fmt.Printf("deleted value is %s\n", value)
		}
	}()

	// // Example: Set a key with an expiration time
	// err := rdb.Set(ctx, "mykey", "myvalue", 10*time.Second).Err()
	// if err != nil {
	// 	log.Fatalf("Failed to set key: %v", err)
	// }
	// fmt.Println("Set key 'mykey' with value 'myvalue' and expiration of 10 seconds")

	// Keep the program running to receive messages

	// time.Sleep(100 * time.Second)
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
// go run pubsub.go

// // terminal 3:
// redis-cli -h localhost -p 6379
// set name srini EX 10
