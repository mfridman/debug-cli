package debugcli

import "github.com/charmbracelet/huh"

type Fooer struct{}

func (f Fooer) Foo() string {
	_ = huh.Confirm{}
	return "foo"
}
