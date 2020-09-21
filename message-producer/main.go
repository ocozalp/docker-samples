package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)
import "github.com/google/uuid"

func main() {
	fmt.Println("Creating client")
	client := redis.NewClient(&redis.Options{Addr: ":12345"})
	ctx := context.Background()
	for {
		uuid := uuid.New().String()
		err := client.Set(ctx, "test_uuid", uuid, time.Hour)

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}

		time.Sleep(time.Second*5)
	}
}
