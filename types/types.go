package types

import "github.com/tokopedia/tdk/go/redis"

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
