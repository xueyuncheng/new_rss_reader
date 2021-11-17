package ecode

var ErrOK Response = Response{
	Code:    0,
	Message: "成功",
	Data:    nil,
}

var ErrInternal = Response{
	Code:    -1,
	Message: "内部错误",
	Data:    nil,
}

var ErrInvalidParams = Response{
	Code:    -2,
	Message: "参数错误",
	Data:    nil,
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r Response) WithData(data interface{}) Response {
	r.Data = data
	return r
}
