package plugin

type Reporter interface {
	Report([]byte) error
}
