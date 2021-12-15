package plugin

// Handler ...
// @Description:
type Handler interface {
	Handle(id string, bytes []byte) error
}
