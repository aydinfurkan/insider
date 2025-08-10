package api

type ApiResponse struct {
	Success bool              `json:"success"`
	Data    interface{}       `json:"data"`
	Error   *ApiResponseError `json:"error"`
}

type ApiResponseError struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Exception string `json:"exception"`
}

func NewSuccessResponse(data interface{}) *ApiResponse {
	return &ApiResponse{
		Success: true,
		Data:    data,
		Error:   nil,
	}
}

func NewErrorResponse(err error, message string, code int) *ApiResponse {
	return &ApiResponse{
		Success: false,
		Data:    nil,
		Error: &ApiResponseError{
			Code:      code,
			Message:   message,
			Exception: err.Error(),
		},
	}
}
