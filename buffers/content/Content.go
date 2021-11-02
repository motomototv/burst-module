// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package content

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Content struct {
	_tab flatbuffers.Table
}

func GetRootAsContent(buf []byte, offset flatbuffers.UOffsetT) *Content {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Content{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsContent(buf []byte, offset flatbuffers.UOffsetT) *Content {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Content{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Content) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Content) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Content) Version() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Content) Type() Type {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return Type(rcv._tab.GetUint32(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Content) MutateType(n Type) bool {
	return rcv._tab.MutateUint32Slot(6, uint32(n))
}

func (rcv *Content) Message(obj *Message) *Message {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Message)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Content) Ext(obj *Ext, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Content) ExtLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Content) Node(obj *Node) *Node {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Node)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func ContentStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func ContentAddVersion(builder *flatbuffers.Builder, version flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(version), 0)
}
func ContentAddType(builder *flatbuffers.Builder, type_ Type) {
	builder.PrependUint32Slot(1, uint32(type_), 0)
}
func ContentAddMessage(builder *flatbuffers.Builder, message flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(message), 0)
}
func ContentAddExt(builder *flatbuffers.Builder, ext flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(ext), 0)
}
func ContentStartExtVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ContentAddNode(builder *flatbuffers.Builder, node flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(node), 0)
}
func ContentEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}