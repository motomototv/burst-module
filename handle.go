package module

// ContentHandler ...
// @Description:
type ContentHandler interface {
	Handle(id ID, bytes []byte, last int64) error
}
