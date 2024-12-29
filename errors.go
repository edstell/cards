package cards

type cardErr string

func (e cardErr) Error() string {
	return string(e)
}

const (
	ErrSyntax = cardErr("invalid syntax")
	ErrRange  = cardErr("input out of range")
)

type ParseErr struct {
	Func  string
	Input string
	Err   error
}

func (e ParseErr) Error() string {
	return "card." + e.Func + ": parsing " + quote(e.Input) + ": " + e.Err.Error()
}

func (e ParseErr) Unwrap() error {
	return e.Err
}

func quote(s string) string {
	return `"` + s + `"`
}
