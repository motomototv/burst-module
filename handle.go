package module

// ContentHandler ...
// @Description:
type ContentHandler interface {
	Handle(id string, bytes []byte) error
}
