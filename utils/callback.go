package utils

import (
	"encoding/json"
	"github.com/l-f-h/Service/conf"
	"github.com/l-f-h/Service/generated"
	"github.com/l-f-h/common"
	"log"
	"net/http"
	"strings"
)

func Callback(url string, response *pb_gen.JudgeResponse) {
	resp, err := json.Marshal(response)
	if err != nil {
		log.Printf("CallBack Marshal response failed: %v", err)
		response.ResCode = common.CallBackMarshalFailed
		return
	}
	log.Printf("CallBack Marshal response: %v", resp)
	callbackResp, err := http.Post(url, conf.CallBackContentType, strings.NewReader(string(resp)))
	if err != nil {
		log.Printf("CallBack Post failed: %v", err)
		response.ResCode = common.CallBackPostFailed
	}
	defer callbackResp.Body.Close()
	response.ResCode = common.CallBackSuccessful
}
