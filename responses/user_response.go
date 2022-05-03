package responses

type Response struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func BuildResponse(statusCode int, message string, data interface{}) Response {
	response := Response{Status: statusCode, Message: message, Data: map[string]interface{}{"data": data}}
	return response
}
