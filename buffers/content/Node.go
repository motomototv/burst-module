// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package content

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Node struct {
	_tab flatbuffers.Table
}

func GetRootAsNode(buf []byte, offset flatbuffers.UOffsetT) *Node {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Node{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsNode(buf []byte, offset flatbuffers.UOffsetT) *Node {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Node{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Node) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Node) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Node) Pid() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Node) Cpuid() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Node) Addr() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Node) Data() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Node) Addrs() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func NodeStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func NodeAddPid(builder *flatbuffers.Builder, pid flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(pid), 0)
}
func NodeAddCpuid(builder *flatbuffers.Builder, cpuid flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(cpuid), 0)
}
func NodeAddAddr(builder *flatbuffers.Builder, addr flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(addr), 0)
}
func NodeAddData(builder *flatbuffers.Builder, data flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(data), 0)
}
func NodeAddAddrs(builder *flatbuffers.Builder, addrs flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(addrs), 0)
}
func NodeEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}