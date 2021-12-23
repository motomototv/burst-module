package module

import (
	"context"
)

type core interface {
	Context() context.Context
	RepoPath() string
	Identity() string
	Register(plugin Module)
	State() NodeState
}
