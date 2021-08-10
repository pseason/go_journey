package model

const (
	SuccessCode int = 200
	ErrorCode   int = 100
)

type ResponseData struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

// Success ResponseData
func SuccessResponseData(data interface{}) ResponseData {
	return ResponseData{
		Code: SuccessCode,
		Data: data,
	}
}

// Error ResponseData
func ErrorResponseData(data interface{}) ResponseData {
	switch dt := data.(type) {
	case error:
		return ResponseData{
			Code: SuccessCode,
			Data: dt.Error(),
		}
	default:
		return ResponseData{
			Code: SuccessCode,
			Data: data,
		}
	}
}
