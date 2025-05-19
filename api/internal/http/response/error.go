package response

import "encoding/json"

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewErrorResponse(err error, code int) ErrorResponse {
	return ErrorResponse{
		Error:   err.Error(),
		Message: "An error occurred",
		Code:    code,
	}
}
func (e ErrorResponse) ToJSON() ([]byte, error) {
	return json.Marshal(e)
}
