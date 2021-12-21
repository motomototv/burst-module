package module

// Module ...
// @Description:
type Module interface {
	Name() string
	Version() int
	Start(node Node) error
	ContentHandler
}
