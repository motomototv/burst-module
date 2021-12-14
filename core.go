package plugin

import (
	"context"
)

type Plugin interface {
	Name() string
	Version() int
	Start(node Node) error
	Handler
}

type core interface {
	Context() context.Context
	RepoPath() string
	Identity() string
	RegisterPlugin(plugin Plugin)
	State() NodeState
}

type Node interface {
	core
	Sender
	Reporter
}
