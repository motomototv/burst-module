// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package exchange

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Buffer struct {
	_tab flatbuffers.Table
}

func GetRootAsBuffer(buf []byte, offset flatbuffers.UOffsetT) *Buffer {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Buffer{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsBuffer(buf []byte, offset flatbuffers.UOffsetT) *Buffer {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Buffer{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Buffer) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Buffer) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Buffer) Security() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Buffer) Type() Type {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return Type(rcv._tab.GetUint32(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Buffer) MutateType(n Type) bool {
	return rcv._tab.MutateUint32Slot(6, uint32(n))
}

func (rcv *Buffer) From() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Buffer) To() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Buffer) Data() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Buffer) Last() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Buffer) MutateLast(n int64) bool {
	return rcv._tab.MutateInt64Slot(14, n)
}

func BufferStart(builder *flatbuffers.Builder) {
	builder.StartObject(6)
}
func BufferAddSecurity(builder *flatbuffers.Builder, security flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(security), 0)
}
func BufferAddType(builder *flatbuffers.Builder, type_ Type) {
	builder.PrependUint32Slot(1, uint32(type_), 0)
}
func BufferAddFrom(builder *flatbuffers.Builder, from flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(from), 0)
}
func BufferAddTo(builder *flatbuffers.Builder, to flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(to), 0)
}
func BufferAddData(builder *flatbuffers.Builder, data flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(data), 0)
}
func BufferAddLast(builder *flatbuffers.Builder, last int64) {
	builder.PrependInt64Slot(5, last, 0)
}
func BufferEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
