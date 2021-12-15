package plugin

import (
	"context"
)

type core interface {
	Context() context.Context
	RepoPath() string
	Identity() string
	RegisterPlugin(plugin Plugin)
	State() NodeState
}

// Node ...
// @Description:
type Node interface {
	core
	Sender
	Reporter
}
