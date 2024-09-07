package counter

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
}

func ResetCounterActivity(ctx context.Context) error {
    err := rdb.Set(ctx, "globalCounter", 0, 0).Err()
    if err != nil {
        return fmt.Errorf("Failed to set global counter: %v", err)
    }
    return err
}

func GetCounterActivity(ctx context.Context) (int, error) {
    counter, err := rdb.Get(ctx, "globalCounter").Int()
    if err != nil {
        return 0, fmt.Errorf("Failed to get global counter: %v", err)
    }
    return counter, err
}

func UpdateCounterActivity(ctx context.Context, workerID int) error {
	err := rdb.Incr(ctx, "globalCounter").Err()
	if err != nil {
		return fmt.Errorf("Failed to increment global counter: %v", err)
	}
	return err
}
