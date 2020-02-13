package assert

// Assertions provides assertion methods around the
// TestingT interface.
type Assertions struct {
	t TestingT
}

// New makes a new Assertions object for the specified TestingT.
func New(t TestingT) *Assertions {
	return &Assertions{
		t: t,
	}
}

//go:generate sh -c "cd ../.codegen && go build && cd - && ../.codegen/_codegen -output-package=assert -template=assertion_forward.go.tmpl -include-format-funcs"
