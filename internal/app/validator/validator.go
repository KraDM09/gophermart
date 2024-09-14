package validator

//go:generate mockery --name=Validator
type Validator interface {
	Struct(s interface{}) error
	Initialize()
}
