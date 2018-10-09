package models

type JsonResult struct {
	Code int64 `json:"code"`
	Msg  string               `json:"msg"`
	Obj  interface{}          `json:"obj"`
}
