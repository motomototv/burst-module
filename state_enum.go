// Code generated by go-enum DO NOT EDIT.
// Version: 0.3.9
// Revision: 5a95b1bbcaf1f8f32542725929d84acdf48e0259
// Build Date: 2021-11-05T17:00:39Z
// Built By: goreleaser

package module

import (
	"fmt"
)

const (
	// NodeStateInvalid is a NodeState of type Invalid.
	NodeStateInvalid NodeState = iota
	// NodeStateStarting is a NodeState of type Starting.
	NodeStateStarting
	// NodeStateConnecting is a NodeState of type Connecting.
	NodeStateConnecting
	// NodeStateConnectFailed is a NodeState of type ConnectFailed.
	NodeStateConnectFailed
	// NodeStateConnected is a NodeState of type Connected.
	NodeStateConnected
	// NodeStateSyncing is a NodeState of type Syncing.
	NodeStateSyncing
	// NodeStateSyncingBoot is a NodeState of type SyncingBoot.
	NodeStateSyncingBoot
	// NodeStateSyncedBoot is a NodeState of type SyncedBoot.
	NodeStateSyncedBoot
	// NodeStateSyncBootFailed is a NodeState of type SyncBootFailed.
	NodeStateSyncBootFailed
	// NodeStateSyncingInfo is a NodeState of type SyncingInfo.
	NodeStateSyncingInfo
	// NodeStateSyncedInfo is a NodeState of type SyncedInfo.
	NodeStateSyncedInfo
	// NodeStateSyncInfoFailed is a NodeState of type SyncInfoFailed.
	NodeStateSyncInfoFailed
	// NodeStateRunning is a NodeState of type Running.
	NodeStateRunning
	// NodeStateStopping is a NodeState of type Stopping.
	NodeStateStopping
	// NodeStateStopped is a NodeState of type Stopped.
	NodeStateStopped
	// NodeStateMax is a NodeState of type Max.
	NodeStateMax
)

const _NodeStateName = "invalidstartingconnectingconnectFailedconnectedsyncingsyncingBootsyncedBootsyncBootFailedsyncingInfosyncedInfosyncInfoFailedrunningstoppingstoppedmax"

var _NodeStateMap = map[NodeState]string{
	NodeStateInvalid:        _NodeStateName[0:7],
	NodeStateStarting:       _NodeStateName[7:15],
	NodeStateConnecting:     _NodeStateName[15:25],
	NodeStateConnectFailed:  _NodeStateName[25:38],
	NodeStateConnected:      _NodeStateName[38:47],
	NodeStateSyncing:        _NodeStateName[47:54],
	NodeStateSyncingBoot:    _NodeStateName[54:65],
	NodeStateSyncedBoot:     _NodeStateName[65:75],
	NodeStateSyncBootFailed: _NodeStateName[75:89],
	NodeStateSyncingInfo:    _NodeStateName[89:100],
	NodeStateSyncedInfo:     _NodeStateName[100:110],
	NodeStateSyncInfoFailed: _NodeStateName[110:124],
	NodeStateRunning:        _NodeStateName[124:131],
	NodeStateStopping:       _NodeStateName[131:139],
	NodeStateStopped:        _NodeStateName[139:146],
	NodeStateMax:            _NodeStateName[146:149],
}

// String implements the Stringer interface.
func (x NodeState) String() string {
	if str, ok := _NodeStateMap[x]; ok {
		return str
	}
	return fmt.Sprintf("NodeState(%d)", x)
}

var _NodeStateValue = map[string]NodeState{
	_NodeStateName[0:7]:     NodeStateInvalid,
	_NodeStateName[7:15]:    NodeStateStarting,
	_NodeStateName[15:25]:   NodeStateConnecting,
	_NodeStateName[25:38]:   NodeStateConnectFailed,
	_NodeStateName[38:47]:   NodeStateConnected,
	_NodeStateName[47:54]:   NodeStateSyncing,
	_NodeStateName[54:65]:   NodeStateSyncingBoot,
	_NodeStateName[65:75]:   NodeStateSyncedBoot,
	_NodeStateName[75:89]:   NodeStateSyncBootFailed,
	_NodeStateName[89:100]:  NodeStateSyncingInfo,
	_NodeStateName[100:110]: NodeStateSyncedInfo,
	_NodeStateName[110:124]: NodeStateSyncInfoFailed,
	_NodeStateName[124:131]: NodeStateRunning,
	_NodeStateName[131:139]: NodeStateStopping,
	_NodeStateName[139:146]: NodeStateStopped,
	_NodeStateName[146:149]: NodeStateMax,
}

// ParseNodeState attempts to convert a string to a NodeState
func ParseNodeState(name string) (NodeState, error) {
	if x, ok := _NodeStateValue[name]; ok {
		return x, nil
	}
	return NodeState(0), fmt.Errorf("%s is not a valid NodeState", name)
}
