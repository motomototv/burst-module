package plugin

// Handler ...
// @Description:
type Handler interface {
	Handle([]byte) error
}
