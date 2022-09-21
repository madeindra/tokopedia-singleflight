package getdata

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

func GetData(w http.ResponseWriter, r *http.Request) {
	var result string
	ctxGetData := context.Background()
	resultRedis := RedisLocal.Get(ctxGetData, "key_testing")
	if resultRedis.Err() != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := &ResponseData{
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
		initData()
	} else {
		result = resultRedis.String()
	}
	w.WriteHeader(http.StatusOK)
	resp := &ResponseData{
		ErrorCode: http.StatusOK,
		Result:    result,
	}
	b, err := json.Marshal(resp)
	if err != nil {
		return
	}
	w.Write(b)
}
