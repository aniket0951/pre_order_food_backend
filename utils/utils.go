package utils

type Response struct {
	Status       bool        `json:"status,omitempty"`
	Message      string      `json:"message,omitempty"`
	Error        string      `json:"error,omitempty"`
	ResponseBody interface{} `json:"data,omitempty"`
}

func BuildSuccessResponse(msg string, responseData interface{}) *Response {
	return &Response{
		Status:       true,
		Message:      msg,
		ResponseBody: responseData,
	}
}

func BuildFailedResponse(errorMsg string) Response {

	response := Response{
		Status:  false,
		Message: "Failed to process the request",
		Error:   errorMsg,
	}

	return response
}
