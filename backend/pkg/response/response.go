package response

import "backend/pkg/pagination"

type Meta = pagination.Meta

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`

	Data interface{} `json:"data"`

	Meta *Meta `json:"meta,omitempty"`
}

func Success(
	message string,
	data interface{},
) Response {

	return Response{
		Status:  true,
		Message: message,
		Data:    data,
	}
}

func SuccessWithMeta(
	message string,
	data interface{},
	meta Meta,
) Response {

	return Response{
		Status:  true,
		Message: message,
		Data:    data,
		Meta:    &meta,
	}
}

func Error(
	message string,
) Response {

	return Response{
		Status:  false,
		Message: message,
		Data:    nil,
	}
}
