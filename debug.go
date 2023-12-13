package debugcli

type Fooer struct{}

func (f Fooer) Foo() string {
	return "foo"
}
