package api

import "fmt"

type MissingParam string

func (e MissingParam) Error() string {
	return fmt.Sprintf("Missing [%s]", string(e))
}

type BadParam string

func (e BadParam) Error() string {
	return fmt.Sprintf("BadParam [%s]", string(e))
}

type NotFound string

func (e NotFound) Error() string {
	return fmt.Sprintf("%s not found", string(e))
}

type Timeout struct{}

func (e Timeout) Error() string {
	return "timeout"
}
