package module

//NodeState is the status of the service
//ENUM(invalid,starting,connecting,connectFailed,connected,syncing,syncingBoot,syncedBoot,syncBootFailed,syncingInfo,syncedInfo,syncInfoFailed,running, stopping, stopped,max)
type NodeState int32
