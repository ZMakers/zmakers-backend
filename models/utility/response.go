package utility

type Response struct {
	Code uint64 `json:"code"`
	Data interface{} `json:"data"`
}
