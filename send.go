package plugin

// Sender ...
// @Description:
type Sender interface {
	Send(id string, msg []byte) error
}
