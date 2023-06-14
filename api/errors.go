package api

import "net/http"

type Error struct {
	Code 	int			`json:"code"`
	Err 	string	`json:"error"`
}

func (e Error) Error() string {
	return e.Err
}

func ErrBadRequest(res string) error {
	return Error{
		Code: http.StatusBadRequest,
		Err: res + " request invalid",
	}
}

func ErrInvalidID(id string) Error {
	return Error{
		Code: http.StatusBadRequest,
		Err: id + " is an Invalid ID",
	}
}

func ErrNotFound(res string) Error {
	return Error{
		Code: http.StatusNotFound,
		Err: res + " resource not found",
	}
}