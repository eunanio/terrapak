package metadata

type ApiResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewApiResponse(code int, message string) *ApiResponse {
	return &ApiResponse{Code: code, Message: message}
}