package plugin

import (
	"context"
)

type Plugin interface {
	Name() string
	Version() int
	Start(node Node) error
	Handler
	Sender
}

type Node interface {
	Context() context.Context
	//Config() *BurstConfig
	RepoPath() string
	Identity() string
	RegisterPlugin(plugin Plugin)
	State() NodeState
}
