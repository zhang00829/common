package response

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Body struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	if err != nil {
		body := &Body{
			Code: 10086,
			Msg:  "error",
			Data: nil,
		}
		httpx.OkJson(w, body)
		return
	}
	httpx.OkJson(w, &Body{
		Code: 0,
		Msg:  "success",
		Data: resp,
	})

}
