// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensor_shape.proto

package tensorflow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Dimensions of a tensor.
type TensorShapeProto struct {
	// Dimensions of the tensor, such as {"input", 30}, {"output", 40}
	// for a 30 x 40 2D tensor.  If an entry has size -1, this
	// corresponds to a dimension of unknown size. The names are
	// optional.
	//
	// The order of entries in "dim" matters: It indicates the layout of the
	// values in the tensor in-memory representation.
	//
	// The first entry in "dim" is the outermost dimension used to layout the
	// values, the last entry is the innermost dimension.  This matches the
	// in-memory layout of RowMajor Eigen tensors.
	//
	// If "dim.size()" > 0, "unknown_rank" must be false.
	Dim []*TensorShapeProto_Dim `protobuf:"bytes,2,rep,name=dim" json:"dim"`
	// If true, the number of dimensions in the shape is unknown.
	//
	// If true, "dim.size()" must be 0.
	UnknownRank bool `protobuf:"varint,3,opt,name=unknown_rank,json=unknownRank" json:"unknown_rank"`
}

func (m *TensorShapeProto) Reset()                    { *m = TensorShapeProto{} }
func (m *TensorShapeProto) String() string            { return proto.CompactTextString(m) }
func (*TensorShapeProto) ProtoMessage()               {}
func (*TensorShapeProto) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func (m *TensorShapeProto) GetDim() []*TensorShapeProto_Dim {
	if m != nil {
		return m.Dim
	}
	return nil
}

func (m *TensorShapeProto) GetUnknownRank() bool {
	if m != nil {
		return m.UnknownRank
	}
	return false
}

// One dimension of the tensor.
type TensorShapeProto_Dim struct {
	// Size of the tensor in that dimension.
	// This value must be >= -1, but values of -1 are reserved for "unknown"
	// shapes (values of -1 mean "unknown" dimension).  Certain wrappers
	// that work with TensorShapeProto may fail at runtime when deserializing
	// a TensorShapeProto containing a dim value of -1.
	Size int64 `protobuf:"varint,1,opt,name=size" json:"size"`
	// Optional name of the tensor dimension.
	Name string `protobuf:"bytes,2,opt,name=name" json:"name"`
}

func (m *TensorShapeProto_Dim) Reset()                    { *m = TensorShapeProto_Dim{} }
func (m *TensorShapeProto_Dim) String() string            { return proto.CompactTextString(m) }
func (*TensorShapeProto_Dim) ProtoMessage()               {}
func (*TensorShapeProto_Dim) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0, 0} }

func (m *TensorShapeProto_Dim) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *TensorShapeProto_Dim) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*TensorShapeProto)(nil), "tensorflow.TensorShapeProto")
	proto.RegisterType((*TensorShapeProto_Dim)(nil), "tensorflow.TensorShapeProto.Dim")
}

func init() { proto.RegisterFile("tensor_shape.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0xb1, 0x4a, 0xc0, 0x30,
	0x10, 0x86, 0xb9, 0x46, 0x44, 0xaf, 0x0e, 0x35, 0x53, 0x70, 0x8a, 0x4e, 0x59, 0x8c, 0x50, 0xdf,
	0xa0, 0xf4, 0x01, 0x4a, 0x74, 0x72, 0x29, 0x11, 0x53, 0x0d, 0x31, 0x49, 0x49, 0x2a, 0x05, 0x9f,
	0xc4, 0x47, 0x75, 0x94, 0x46, 0xa1, 0xd2, 0xed, 0xbf, 0x8f, 0xef, 0xb8, 0xfb, 0x91, 0x2e, 0x26,
	0xe4, 0x98, 0xc6, 0xfc, 0xa6, 0x67, 0x23, 0xe7, 0x14, 0x97, 0x48, 0xf1, 0x97, 0x4d, 0xef, 0x71,
	0xbd, 0xf9, 0x02, 0x6c, 0x1e, 0xcb, 0xf8, 0xb0, 0x19, 0x43, 0x11, 0x5a, 0x24, 0x2f, 0xd6, 0xb3,
	0x8a, 0x13, 0x51, 0xb7, 0x5c, 0xee, 0xba, 0x3c, 0xaa, 0xb2, 0xb7, 0x5e, 0x6d, 0x32, 0xbd, 0xc6,
	0x8b, 0x8f, 0xe0, 0x42, 0x5c, 0xc3, 0x98, 0x74, 0x70, 0x8c, 0x70, 0x10, 0x67, 0xaa, 0xfe, 0x63,
	0x4a, 0x07, 0x77, 0x75, 0x8b, 0xa4, 0xb7, 0x9e, 0x52, 0x3c, 0xc9, 0xf6, 0xd3, 0x30, 0xe0, 0x20,
	0x88, 0x2a, 0x79, 0x63, 0x41, 0x7b, 0xc3, 0x2a, 0x0e, 0xe2, 0x5c, 0x95, 0xdc, 0xf5, 0xc8, 0x62,
	0x7a, 0xfd, 0x7f, 0x7d, 0x4a, 0xda, 0x9b, 0x35, 0x26, 0xd7, 0x5d, 0x1e, 0x1f, 0xc9, 0x03, 0x3c,
	0x35, 0xa5, 0xde, 0xdd, 0xbe, 0xf0, 0x0d, 0xf0, 0x7c, 0x5a, 0xe0, 0xfd, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xdb, 0x6f, 0x3f, 0x12, 0x09, 0x01, 0x00, 0x00,
}