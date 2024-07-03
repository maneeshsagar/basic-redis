package util

type Response struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Model interface{} `json:"model"`
}

func NewResponseWithModel(code int, msg string, model interface{}) Response {
	return Response{
		Code:  code,
		Msg:   msg,
		Model: model,
	}
}

