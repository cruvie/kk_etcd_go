package api_resp

import (
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_stage"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"log/slog"
	"net/http"
)

func Fail(stage *kku_stage.Stage, res *ApiResp, data proto.Message) *ApiResp {
	if res == nil && data == nil {
		res = &ApiResp{
			Code: http.StatusInternalServerError,
			Msg:  "System Error",
		}
		return res
	}
	if res == nil {
		res = &ApiResp{
			Code: http.StatusBadRequest,
			Msg:  "Error",
		}
	} else {
		if res.Code == 0 {
			res.Code = http.StatusBadRequest
		}
		if res.Msg == "" {
			res.Msg = "Error"
		}
	}
	res.Data = getAnyData(stage, data)
	return res
}

func Success(stage *kku_stage.Stage, res *ApiResp, data proto.Message) *ApiResp {
	if res == nil && data == nil {
		res := &ApiResp{
			Code: http.StatusOK,
			Msg:  "OK",
		}
		return res
	}
	if res == nil {
		res = &ApiResp{
			Code: http.StatusOK,
			Msg:  "OK",
		}
	} else {
		if res.Code == 0 {
			res.Code = http.StatusOK
		}
		if res.Msg == "" {
			res.Msg = "OK"
		}
	}
	res.Data = getAnyData(stage, data)

	return res
}

func getAnyData(stage *kku_stage.Stage, data proto.Message) *anypb.Any {
	if data == nil {
		return nil
	}
	anyData, err := anypb.New(data)
	if err != nil {
		logBody := kku_stage.NewLogBody().SetError(err).SetTraceId(stage.TraceId)
		slog.Error("create any data error", logBody.GetLogArgs()...)
		return nil
	}
	return anyData
}
