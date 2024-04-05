package httpresponse

type HttpResponse struct {
	Code     int                    `json:"code"`
	Message  string                 `json:"message"`
	Data     interface{}            `json:"data"`
	MetaData map[string]interface{} `json:"meta_data"`
}

func New(res HttpResponse) HttpResponse {
	return HttpResponse{
		Code:     res.Code,
		Message:  res.Message,
		Data:     res.Data,
		MetaData: res.MetaData,
	}
}
func NewCustomize(Code int, Message string, Data interface{}, MetaData map[string]interface{}) HttpResponse {
	return HttpResponse{
		Code:     Code,
		Message:  Message,
		Data:     Data,
		MetaData: MetaData,
	}
}
