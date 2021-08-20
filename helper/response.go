package helper

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Errors  []error     `json:"errors"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct {
}

func BuildReponse(status int, message string, errors []error, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Errors:  errors,
		Data:    data,
	}

	return res
}
