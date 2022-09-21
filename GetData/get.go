package getdata

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bagusandrian/dummy_app/types"
)

func GetData(w http.ResponseWriter, r *http.Request) {
	var result string
	ctxGetData := context.Background()
	resultRedis := types.RedisLocal.Get(ctxGetData, "key_testing")
	if resultRedis.Err() != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := &types.ResponseData{
			ErrorCode:    http.StatusInternalServerError,
			ErrorMessage: resultRedis.Err().Error(),
		}
		b, err := json.Marshal(resp)
		if err != nil {
			return
		}
		w.Write(b)
		return
	}
	if resultRedis.String() == "" {
		result = "Hello Data Founded"
		time.Sleep(1 * time.Second)
		SetRedis()
	} else {
		result = resultRedis.String()
	}
	w.WriteHeader(http.StatusOK)
	resp := &types.ResponseData{
		ErrorCode: http.StatusOK,
		Result:    result,
	}
	b, err := json.Marshal(resp)
	if err != nil {
		return
	}
	w.Write(b)
}

func SetRedis() {
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
