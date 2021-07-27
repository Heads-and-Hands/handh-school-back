package models

type Error struct {
	Message string `json:"message"`
}

func NewErrorByWrapping(err error) Error {
	return Error{Message: err.Error()}
}
