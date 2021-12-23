package module

// Node ...
// @Description:
type Node interface {
	core
	Sender
	Reporter
	Module(t string) Module
}
