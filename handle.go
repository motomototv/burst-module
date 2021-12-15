package plugin

// ContentHandler ...
// @Description:
type ContentHandler interface {
	Handle(id string, bytes []byte) error
}
