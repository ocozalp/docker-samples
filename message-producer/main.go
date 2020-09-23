package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

func main() {
	fmt.Println("Creating client")
	client := redis.NewClient(&redis.Options{Addr: "redis:6379"})
	ctx := context.Background()
	for {
		uuid := uuid.New().String()
		result := client.Set(ctx, "test_uuid", uuid, time.Hour)

		if result.Err() != nil {
			fmt.Printf("Error: %v\n", result.Err())
		} else {
			fmt.Println("Key successfully updated...")
		}

		time.Sleep(time.Second * 5)
	}
}
