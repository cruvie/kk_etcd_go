package api_resp

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"log"
	"net/http"
)

//====================Fail

func Fail() *ApiResp {
	apiRes := ApiResp{
		Code: http.StatusBadRequest,
		Msg:  "System Error",
	}

	return &apiRes
}

func FailMsg(msg string) *ApiResp {
	apiRes := ApiResp{
		Code: http.StatusBadRequest,
		Msg:  msg,
	}

	return &apiRes
}

func FailCodeMsg(code int32, msg string) *ApiResp {
	apiRes := ApiResp{
		Code: code,
		Msg:  msg,
	}

	return &apiRes
}

func FailMsgData(msg string, data proto.Message) *ApiResp {
	anyData, err := anypb.New(data)
	if err != nil {
		log.Println("error serializing response data", err)
		return nil
	}
	apiRes := ApiResp{
		Code: http.StatusBadRequest,
		Msg:  msg,
		Data: anyData,
	}

	return &apiRes
}

//====================Success

func Success() *ApiResp {
	apiRes := ApiResp{
		Code: http.StatusOK,
		Msg:  "OK",
	}

	return &apiRes
}

func SuccessMsg(msg string) *ApiResp {
	apiRes := ApiResp{
		Code: http.StatusOK,
		Msg:  msg,
	}
	return &apiRes
}

func SuccessData(data proto.Message) *ApiResp {
	anyData, err := anypb.New(data)
	if err != nil {
		log.Println("error serializing response data", err)
		return nil
	}
	apiRes := ApiResp{
		Code: http.StatusOK,
		Msg:  "OK",
		Data: anyData,
	}
	return &apiRes
}

func SuccessCodeMsg(code int32, msg string) *ApiResp {
	apiRes := ApiResp{
		Code: code,
		Msg:  msg,
	}
	return &apiRes
}

func SuccessMsgData(msg string, data proto.Message) *ApiResp {
	anyData, err := anypb.New(data)
	if err != nil {
		log.Println("error serializing response data", err)
		//return nil
	}
	apiRes := ApiResp{
		Code: http.StatusOK,
		Msg:  msg,
		Data: anyData,
	}
	return &apiRes
}

func FullCodeMsgData(code int32, msg string, data proto.Message) *ApiResp {
	anyData, err := anypb.New(data)
	if err != nil {
		log.Println("error serializing response data", err)
		return nil
	}
	apiRes := ApiResp{
		Code: code,
		Msg:  msg,
		Data: anyData,
	}

	return &apiRes
}
