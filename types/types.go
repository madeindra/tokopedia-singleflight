package types

import "github.com/go-redis/redis/v8"

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
