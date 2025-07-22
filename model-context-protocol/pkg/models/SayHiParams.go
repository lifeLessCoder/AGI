package models

type SayHiParams struct {
	Name string `json:"name" jsonschema:"Person to greet"`
}
