package plugin

// Handler ...
// @Description:
type Handler interface {
	// Handle is the main entry point for the plugin.
	//
	// The Handle is passed to the plugin main() function.
	//
	// The Handle is responsible for creating and managing the plugin's
	// connection to the plugin process.
	//
	// The Handle is responsible for reading and writing to the plugin
	// process's stdin and stdout.
	//
	// The Handle is responsible for reading and writing to the plugin
	// process's stderr.
	//
	// The Handle is responsible for exiting the plugin process.
	Handle([]byte) error
}
