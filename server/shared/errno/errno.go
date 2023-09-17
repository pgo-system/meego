package errno

type Response struct {
	Code    int16  `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (r *Response) Success(data any) {
	r.Data = data
	r.Code = Success
	r.Message = SuccessMessage
}

// 通用错误码
var (
	Success        int16 = 0
	SuccessMessage       = "ok"

	BusinessError int16 = 500

	DatabaseError        int16 = 10001
	DatabaseErrorMessage       = "数据库错误"

	ParamsError         int16 = 10002
	ParamsErrorMessage        = "参数错误"
	NoExistError        int16 = 10003
	NoExistErrorMessage       = "数据不存在"
)
