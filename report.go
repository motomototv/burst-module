package module

type Reporter interface {
	Report([]byte) error
}
