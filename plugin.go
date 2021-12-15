package plugin

// Plugin ...
// @Description:
type Plugin interface {
	Name() string
	Version() int
	Start(node Node) error
	ContentHandler
}
