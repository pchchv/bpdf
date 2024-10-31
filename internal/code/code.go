package code

// codeInstance is a singleton of code.
// It is used to ensure that it is not instantiated more than once,
// as it is not necessary since the code is stateless.
var codeInstance *code = nil

type code struct{}

// New create a Code (Singleton).
func New() *code {
	if codeInstance == nil {
		codeInstance = &code{}
	}
	return codeInstance
}
