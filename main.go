package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	RedisLocal *redis.Client
)

type (
	ResponseData struct {
		ErrorCode    int    `json:"error_code"`
		ErrorMessage string `json:"error_message"`
		Result       string `json:"result"`
	}
)

func init() {
	initRedis()
	initData()
}
func main() {
	http.HandleFunc("/get_data", func(w http.ResponseWriter, r *http.Request) {
		// get data redis
		w.Header().Set("Content-Type", "application/json")
		getData(w, r)
	})
	log.Println("running server on port :8080")
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func initRedis() {
	var redisHost = "localhost:6379"
	var redisPassword = ""

	RedisLocal = redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       0,
	})
	checkConnection := RedisLocal.Ping(context.Background())
	if checkConnection.Err() != nil {
		log.Panic("Failed create connection redis")
	}
	fmt.Println("redis client initialized")
}

func initData() {
	key := "key_testing"
	data := "Hello Data Founded"
	ttl := time.Duration(3) * time.Hour

	// store data using SET command
	op1 := RedisLocal.Set(context.Background(), key, data, ttl)
	if err := op1.Err(); err != nil {
		fmt.Printf("unable to SET data. error: %v", err)
		return
	}
	log.Println("finish init Redis")
}

func getData(w http.ResponseWriter, r *http.Request) {
	ctxGetData := context.Background()
	result := RedisLocal.Get(ctxGetData, "key_testing")
	if result.Err() != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := &ResponseData{
			ErrorCode:    http.StatusInternalServerError,
			ErrorMessage: result.Err().Error(),
		}
		b, err := json.Marshal(resp)
		if err != nil {
			return
		}
		w.Write(b)
		return
	}
	w.WriteHeader(http.StatusOK)
	resp := &ResponseData{
		ErrorCode: http.StatusOK,
		Result:    result.String(),
	}
	b, err := json.Marshal(resp)
	if err != nil {
		return
	}
	w.Write(b)
}
