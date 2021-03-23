package model

type Response struct {
	Data data `json:"data"`
}

type ResponseWithPagination struct {
	Data       data       `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type data interface {
}

type Pagination struct {
	Page          int `json:"page"`
	Size          int `json:"size"`
	TotalElements int `json:"totalElements"`
}

func CreateResponse(data interface{}) *Response {
	return &Response{
		data,
	}
}

func CreateResponseWithPagination(data interface{}, page int, size int, total int) *ResponseWithPagination {
	return &ResponseWithPagination{
		data, Pagination{
			page,
			size,
			total,
		},
	}
}
