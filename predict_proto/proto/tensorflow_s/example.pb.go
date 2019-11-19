// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example.proto

package tensorflow_s

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Example struct {
	Features *Features `protobuf:"bytes,1,opt,name=features" json:"features"`
}

func (m *Example) Reset()                    { *m = Example{} }
func (m *Example) String() string            { return proto.CompactTextString(m) }
func (*Example) ProtoMessage()               {}
func (*Example) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Example) GetFeatures() *Features {
	if m != nil {
		return m.Features
	}
	return nil
}

type SequenceExample struct {
	Context      *Features     `protobuf:"bytes,1,opt,name=context" json:"context"`
	FeatureLists *FeatureLists `protobuf:"bytes,2,opt,name=feature_lists,json=featureLists" json:"feature_lists"`
}

func (m *SequenceExample) Reset()                    { *m = SequenceExample{} }
func (m *SequenceExample) String() string            { return proto.CompactTextString(m) }
func (*SequenceExample) ProtoMessage()               {}
func (*SequenceExample) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *SequenceExample) GetContext() *Features {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *SequenceExample) GetFeatureLists() *FeatureLists {
	if m != nil {
		return m.FeatureLists
	}
	return nil
}

func init() {
	proto.RegisterType((*Example)(nil), "tensorflow_s.Example")
	proto.RegisterType((*SequenceExample)(nil), "tensorflow_s.SequenceExample")
}

func init() { proto.RegisterFile("example.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x29, 0x49, 0xcd, 0x2b, 0xce,
	0x2f, 0x4a, 0xcb, 0xc9, 0x2f, 0x8f, 0x2f, 0x96, 0xe2, 0x4d, 0x4b, 0x4d, 0x2c, 0x29, 0x2d, 0x82,
	0x4a, 0x2a, 0xd9, 0x72, 0xb1, 0xbb, 0x42, 0x54, 0x0b, 0x19, 0x71, 0x71, 0x40, 0xe5, 0x8a, 0x25,
	0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0xc4, 0xf4, 0x90, 0xb5, 0xea, 0xb9, 0x41, 0x65, 0x83, 0xe0,
	0xea, 0x94, 0x5a, 0x18, 0xb9, 0xf8, 0x83, 0x53, 0x0b, 0x4b, 0x53, 0xf3, 0x92, 0x53, 0x61, 0xe6,
	0x18, 0x70, 0xb1, 0x27, 0xe7, 0xe7, 0x95, 0xa4, 0x56, 0x94, 0x10, 0x30, 0x06, 0xa6, 0x4c, 0xc8,
	0x9e, 0x0b, 0xe6, 0xaa, 0xf8, 0x9c, 0xcc, 0xe2, 0x92, 0x62, 0x09, 0x26, 0xb0, 0x3e, 0x29, 0xac,
	0xfa, 0x7c, 0x40, 0x2a, 0x82, 0x78, 0xd2, 0x90, 0x78, 0x4e, 0x0e, 0x5c, 0x62, 0xf9, 0x45, 0xe9,
	0x48, 0xca, 0xf5, 0xa0, 0x41, 0xe0, 0xc4, 0x0b, 0x75, 0x55, 0x00, 0xc8, 0xb7, 0xc5, 0x01, 0x8c,
	0x51, 0x42, 0x60, 0x7f, 0xeb, 0x23, 0x9b, 0xfc, 0x83, 0x91, 0x31, 0x89, 0x0d, 0x2c, 0x6c, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x1f, 0x38, 0xb4, 0x3c, 0x01, 0x00, 0x00,
}