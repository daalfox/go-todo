package utils

type Validator interface {
	Validate() (problems map[string]string)
}
