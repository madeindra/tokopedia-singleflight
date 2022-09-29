package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bagusandrian/dummy_app/getdata"
	"github.com/bagusandrian/dummy_app/types"
	"github.com/go-redis/redis/v8"
)

func init() {
	initRedis()
	initData()
}

func main() {
	http.HandleFunc("/get_data", func(w http.ResponseWriter, r *http.Request) {
		// get data redis
		w.Header().Set("Content-Type", "application/json")
		getdata.GetData(w, r)
	})
	log.Println("running server on port :8080")
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func initRedis() {
	var redisHost = "localhost:6379"
	var redisPassword = ""

	types.RedisLocal = redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       0,
	})
	checkConnection := types.RedisLocal.Ping(context.Background())
	if checkConnection.Err() != nil {
		log.Panic("Failed create connection redis")
	}
	fmt.Println("redis client initialized")
}

func initData() {
	key := "key_testing"
	data := "Hello Data Founded"
	ttl := time.Duration(3) * time.Minute

	// store data using SET command
	op1 := types.RedisLocal.Set(context.Background(), key, data, ttl)
	if err := op1.Err(); err != nil {
		fmt.Printf("unable to SET data. error: %v", err)
		return
	}
	log.Println("finish init Redis")
}
