// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tracker.proto

/*
Package anomalytracker is a generated protocol buffer package.

It is generated from these files:
	tracker.proto

It has these top-level messages:
	Anomaly
*/
package anomalytracker

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Anomaly struct {
	Id         string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	System     string                     `protobuf:"bytes,2,opt,name=system" json:"system,omitempty"`
	Type       string                     `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Name       string                     `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	InternalId int64                      `protobuf:"varint,5,opt,name=internal_id,json=internalId" json:"internal_id,omitempty"`
	UserId     int64                      `protobuf:"varint,6,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	GroupId    int64                      `protobuf:"varint,7,opt,name=group_id,json=groupId" json:"group_id,omitempty"`
	Created    *google_protobuf.Timestamp `protobuf:"bytes,8,opt,name=created" json:"created,omitempty"`
	Editable   bool                       `protobuf:"varint,9,opt,name=editable" json:"editable,omitempty"`
}

func (m *Anomaly) Reset()                    { *m = Anomaly{} }
func (m *Anomaly) String() string            { return proto.CompactTextString(m) }
func (*Anomaly) ProtoMessage()               {}
func (*Anomaly) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Anomaly) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Anomaly) GetSystem() string {
	if m != nil {
		return m.System
	}
	return ""
}

func (m *Anomaly) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Anomaly) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Anomaly) GetInternalId() int64 {
	if m != nil {
		return m.InternalId
	}
	return 0
}

func (m *Anomaly) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Anomaly) GetGroupId() int64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *Anomaly) GetCreated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *Anomaly) GetEditable() bool {
	if m != nil {
		return m.Editable
	}
	return false
}

func init() {
	proto.RegisterType((*Anomaly)(nil), "anomalytracker.Anomaly")
}

func init() { proto.RegisterFile("tracker.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0x31, 0x6f, 0x83, 0x30,
	0x10, 0x85, 0x65, 0x92, 0x62, 0x72, 0x51, 0x33, 0xdc, 0xd0, 0xba, 0x2c, 0x41, 0x9d, 0x98, 0x88,
	0xd4, 0xf6, 0x0f, 0x74, 0x64, 0x45, 0xdd, 0x2b, 0x13, 0x5f, 0x91, 0x55, 0x8c, 0x91, 0x31, 0x03,
	0x5b, 0x7f, 0x7a, 0x65, 0x13, 0xb2, 0xdd, 0xfb, 0xbe, 0x7b, 0xc3, 0x83, 0x47, 0xef, 0xe4, 0xf5,
	0x97, 0x5c, 0x35, 0x3a, 0xeb, 0x2d, 0x9e, 0xe4, 0x60, 0x8d, 0xec, 0x97, 0x1b, 0xcd, 0xcf, 0x9d,
	0xb5, 0x5d, 0x4f, 0x97, 0x68, 0xdb, 0xf9, 0xe7, 0xe2, 0xb5, 0xa1, 0xc9, 0x4b, 0x33, 0xae, 0x85,
	0xd7, 0xbf, 0x04, 0xf8, 0xe7, 0xda, 0xc1, 0x13, 0x24, 0x5a, 0x09, 0x56, 0xb0, 0xf2, 0xd0, 0x24,
	0x5a, 0xe1, 0x13, 0xa4, 0xd3, 0x32, 0x79, 0x32, 0x22, 0x89, 0xec, 0x96, 0x10, 0x61, 0xef, 0x97,
	0x91, 0xc4, 0x2e, 0xd2, 0x78, 0x07, 0x36, 0x48, 0x43, 0x62, 0xbf, 0xb2, 0x70, 0xe3, 0x19, 0x8e,
	0x7a, 0xf0, 0xe4, 0x06, 0xd9, 0x7f, 0x6b, 0x25, 0x1e, 0x0a, 0x56, 0xee, 0x1a, 0xd8, 0x50, 0xad,
	0xf0, 0x19, 0xf8, 0x3c, 0x91, 0x0b, 0x32, 0x8d, 0x32, 0x0d, 0xb1, 0x56, 0xf8, 0x02, 0x59, 0xe7,
	0xec, 0x3c, 0x06, 0xc3, 0xa3, 0xe1, 0x31, 0xd7, 0x0a, 0x3f, 0x80, 0x5f, 0x1d, 0x49, 0x4f, 0x4a,
	0x64, 0x05, 0x2b, 0x8f, 0x6f, 0x79, 0xb5, 0x6e, 0xac, 0xb6, 0x8d, 0xd5, 0xd7, 0xb6, 0xb1, 0xd9,
	0x5e, 0x31, 0x87, 0x8c, 0x94, 0xf6, 0xb2, 0xed, 0x49, 0x1c, 0x0a, 0x56, 0x66, 0xcd, 0x3d, 0xb7,
	0x69, 0x2c, 0xbe, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0x7a, 0xdf, 0xb8, 0x68, 0x4b, 0x01, 0x00,
	0x00,
}
