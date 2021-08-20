package helper

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Errors  []string    `json:"errors"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct {
}

func BuildResponse(status int, message string, errors []string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Errors:  errors,
		Data:    data,
	}

	return res
}
