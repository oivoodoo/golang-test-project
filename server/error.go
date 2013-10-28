package server

type Error struct {
	Messages []string
}

var ERROR *Error = &Error {}

func (error *Error) Add(message string) {
	error.Messages = append(error.Messages, message)
}

