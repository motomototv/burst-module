package plugin

type Plugin interface {
	Name() string
	Version() int
	Handler
	Sender
}
