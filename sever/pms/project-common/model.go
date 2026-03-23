package common

type BusinessCode int

type Result struct {
	Code BusinessCode `json:"code"`
	Msg  string       `json:"msg"`
	Data any          `json:"data"`
}

func (r *Result) Sucess(data any) *Result {
	r.Code = 200
	r.Msg = "success"
	r.Data = data
	return r
}
