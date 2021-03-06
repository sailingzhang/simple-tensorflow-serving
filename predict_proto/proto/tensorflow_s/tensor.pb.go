// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensor.proto

package tensorflow_s

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Protocol buffer representing a tensor.
type TensorProto struct {
	Dtype DataType `protobuf:"varint,1,opt,name=dtype,enum=tensorflow_s.DataType" json:"dtype"`
	// Shape of the tensor.  TODO(touts): sort out the 0-rank issues.
	TensorShape *TensorShapeProto `protobuf:"bytes,2,opt,name=tensor_shape,json=tensorShape" json:"tensor_shape"`
	// Version number.
	//
	// In version 0, if the "repeated xxx" representations contain only one
	// element, that element is repeated to fill the shape.  This makes it easy
	// to represent a constant Tensor with a single value.
	VersionNumber int32 `protobuf:"varint,3,opt,name=version_number,json=versionNumber" json:"version_number"`
	// Serialized raw tensor content from either Tensor::AsProtoTensorContent or
	// memcpy in tensorflow::grpc::EncodeTensorToByteBuffer. This representation
	// can be used for all tensor types. The purpose of this representation is to
	// reduce serialization overhead during RPC call by avoiding serialization of
	// many repeated small items.
	TensorContent []byte `protobuf:"bytes,4,opt,name=tensor_content,json=tensorContent,proto3" json:"tensor_content"`
	// DT_HALF, DT_BFLOAT16. Note that since protobuf has no int16 type, we'll
	// have some pointless zero padding for each value here.
	HalfVal []int32 `protobuf:"varint,13,rep,packed,name=half_val,json=halfVal" json:"half_val"`
	// DT_FLOAT.
	FloatVal []float32 `protobuf:"fixed32,5,rep,packed,name=float_val,json=floatVal" json:"float_val"`
	// DT_DOUBLE.
	DoubleVal []float64 `protobuf:"fixed64,6,rep,packed,name=double_val,json=doubleVal" json:"double_val"`
	// DT_INT32, DT_INT16, DT_INT8, DT_UINT8.
	IntVal []int32 `protobuf:"varint,7,rep,packed,name=int_val,json=intVal" json:"int_val"`
	// DT_STRING
	StringVal [][]byte `protobuf:"bytes,8,rep,name=string_val,json=stringVal,proto3" json:"string_val"`
	// DT_COMPLEX64. scomplex_val(2*i) and scomplex_val(2*i+1) are real
	// and imaginary parts of i-th single precision complex.
	ScomplexVal []float32 `protobuf:"fixed32,9,rep,packed,name=scomplex_val,json=scomplexVal" json:"scomplex_val"`
	// DT_INT64
	Int64Val []int64 `protobuf:"varint,10,rep,packed,name=int64_val,json=int64Val" json:"int64_val"`
	// DT_BOOL
	BoolVal []bool `protobuf:"varint,11,rep,packed,name=bool_val,json=boolVal" json:"bool_val"`
	// DT_COMPLEX128. dcomplex_val(2*i) and dcomplex_val(2*i+1) are real
	// and imaginary parts of i-th double precision complex.
	DcomplexVal []float64 `protobuf:"fixed64,12,rep,packed,name=dcomplex_val,json=dcomplexVal" json:"dcomplex_val"`
	// DT_RESOURCE
	ResourceHandleVal []*ResourceHandleProto `protobuf:"bytes,14,rep,name=resource_handle_val,json=resourceHandleVal" json:"resource_handle_val"`
	// DT_VARIANT
	VariantVal []*VariantTensorDataProto `protobuf:"bytes,15,rep,name=variant_val,json=variantVal" json:"variant_val"`
	// DT_UINT32
	Uint32Val []uint32 `protobuf:"varint,16,rep,packed,name=uint32_val,json=uint32Val" json:"uint32_val"`
	// DT_UINT64
	Uint64Val []uint64 `protobuf:"varint,17,rep,packed,name=uint64_val,json=uint64Val" json:"uint64_val"`
}

func (m *TensorProto) Reset()                    { *m = TensorProto{} }
func (m *TensorProto) String() string            { return proto.CompactTextString(m) }
func (*TensorProto) ProtoMessage()               {}
func (*TensorProto) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

func (m *TensorProto) GetDtype() DataType {
	if m != nil {
		return m.Dtype
	}
	return DataType_DT_INVALID
}

func (m *TensorProto) GetTensorShape() *TensorShapeProto {
	if m != nil {
		return m.TensorShape
	}
	return nil
}

func (m *TensorProto) GetVersionNumber() int32 {
	if m != nil {
		return m.VersionNumber
	}
	return 0
}

func (m *TensorProto) GetTensorContent() []byte {
	if m != nil {
		return m.TensorContent
	}
	return nil
}

func (m *TensorProto) GetHalfVal() []int32 {
	if m != nil {
		return m.HalfVal
	}
	return nil
}

func (m *TensorProto) GetFloatVal() []float32 {
	if m != nil {
		return m.FloatVal
	}
	return nil
}

func (m *TensorProto) GetDoubleVal() []float64 {
	if m != nil {
		return m.DoubleVal
	}
	return nil
}

func (m *TensorProto) GetIntVal() []int32 {
	if m != nil {
		return m.IntVal
	}
	return nil
}

func (m *TensorProto) GetStringVal() [][]byte {
	if m != nil {
		return m.StringVal
	}
	return nil
}

func (m *TensorProto) GetScomplexVal() []float32 {
	if m != nil {
		return m.ScomplexVal
	}
	return nil
}

func (m *TensorProto) GetInt64Val() []int64 {
	if m != nil {
		return m.Int64Val
	}
	return nil
}

func (m *TensorProto) GetBoolVal() []bool {
	if m != nil {
		return m.BoolVal
	}
	return nil
}

func (m *TensorProto) GetDcomplexVal() []float64 {
	if m != nil {
		return m.DcomplexVal
	}
	return nil
}

func (m *TensorProto) GetResourceHandleVal() []*ResourceHandleProto {
	if m != nil {
		return m.ResourceHandleVal
	}
	return nil
}

func (m *TensorProto) GetVariantVal() []*VariantTensorDataProto {
	if m != nil {
		return m.VariantVal
	}
	return nil
}

func (m *TensorProto) GetUint32Val() []uint32 {
	if m != nil {
		return m.Uint32Val
	}
	return nil
}

func (m *TensorProto) GetUint64Val() []uint64 {
	if m != nil {
		return m.Uint64Val
	}
	return nil
}

// Protocol buffer representing the serialization format of DT_VARIANT tensors.
type VariantTensorDataProto struct {
	// Name of the type of objects being serialized.
	TypeName string `protobuf:"bytes,1,opt,name=type_name,json=typeName" json:"type_name"`
	// Portions of the object that are not Tensors.
	Metadata []byte `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata"`
	// Tensors contained within objects being serialized.
	Tensors []*TensorProto `protobuf:"bytes,3,rep,name=tensors" json:"tensors"`
}

func (m *VariantTensorDataProto) Reset()                    { *m = VariantTensorDataProto{} }
func (m *VariantTensorDataProto) String() string            { return proto.CompactTextString(m) }
func (*VariantTensorDataProto) ProtoMessage()               {}
func (*VariantTensorDataProto) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{1} }

func (m *VariantTensorDataProto) GetTypeName() string {
	if m != nil {
		return m.TypeName
	}
	return ""
}

func (m *VariantTensorDataProto) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *VariantTensorDataProto) GetTensors() []*TensorProto {
	if m != nil {
		return m.Tensors
	}
	return nil
}

func init() {
	proto.RegisterType((*TensorProto)(nil), "tensorflow_s.TensorProto")
	proto.RegisterType((*VariantTensorDataProto)(nil), "tensorflow_s.VariantTensorDataProto")
}

func init() { proto.RegisterFile("tensor.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 535 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0xd1, 0x6e, 0xd3, 0x3e,
	0x14, 0xc6, 0x95, 0x79, 0x6d, 0x13, 0x27, 0xed, 0x7f, 0xcb, 0x5f, 0x4c, 0xa1, 0xd3, 0xc0, 0x9b,
	0x98, 0xe4, 0x0b, 0x54, 0xa4, 0x16, 0x71, 0xbf, 0x02, 0x12, 0x57, 0xd3, 0x08, 0xd3, 0x2e, 0xb8,
	0x89, 0xdc, 0xc6, 0x5d, 0x23, 0x12, 0xbb, 0x72, 0xdc, 0x8e, 0xbd, 0x00, 0xef, 0xc8, 0x9b, 0x70,
	0x89, 0x7c, 0x4e, 0xda, 0xa5, 0x68, 0x97, 0xf9, 0xce, 0xcf, 0xdf, 0xe7, 0xe3, 0x73, 0x42, 0x23,
	0x2b, 0x55, 0xad, 0xcd, 0x68, 0x65, 0xb4, 0xd5, 0x71, 0xf3, 0xb5, 0x28, 0xf5, 0x43, 0x56, 0x0f,
	0x5f, 0x18, 0x59, 0xeb, 0xb5, 0x99, 0xcb, 0x6c, 0x29, 0x54, 0x5e, 0x4a, 0x84, 0x86, 0x31, 0x42,
	0x59, 0xbd, 0x14, 0xab, 0xad, 0x16, 0xda, 0xc7, 0x95, 0xac, 0xf1, 0xe3, 0xe2, 0x77, 0x87, 0x86,
	0xb7, 0xc0, 0xdc, 0x80, 0xeb, 0x5b, 0xda, 0xc9, 0x5d, 0x3d, 0xf1, 0x98, 0xc7, 0x07, 0xe3, 0x93,
	0x51, 0x3b, 0x65, 0xf4, 0x49, 0x58, 0x71, 0xfb, 0xb8, 0x92, 0x29, 0x42, 0xf1, 0xd5, 0xf6, 0x4e,
	0x18, 0x90, 0x1c, 0x30, 0x8f, 0x87, 0xe3, 0x57, 0xfb, 0x87, 0xd0, 0xfe, 0x9b, 0x03, 0x20, 0x23,
	0x0d, 0xed, 0x93, 0x12, 0x5f, 0xd2, 0xc1, 0x46, 0x9a, 0xba, 0xd0, 0x2a, 0x53, 0xeb, 0x6a, 0x26,
	0x4d, 0x42, 0x98, 0xc7, 0x3b, 0x69, 0xbf, 0x51, 0xaf, 0x41, 0x74, 0x58, 0x93, 0x34, 0xd7, 0xca,
	0x4a, 0x65, 0x93, 0x43, 0xe6, 0xf1, 0x28, 0xed, 0xa3, 0xfa, 0x11, 0xc5, 0xf8, 0x8c, 0xfa, 0x4b,
	0x51, 0x2e, 0xb2, 0x8d, 0x28, 0x93, 0x3e, 0x23, 0xbc, 0x33, 0x3d, 0x38, 0xf2, 0xd2, 0x9e, 0xd3,
	0xee, 0x44, 0x19, 0xbf, 0xa6, 0xc1, 0xa2, 0xd4, 0xc2, 0x42, 0xbd, 0xc3, 0x08, 0x3f, 0x80, 0xba,
	0x0f, 0xa2, 0x03, 0xce, 0x29, 0xcd, 0xf5, 0x7a, 0x56, 0x4a, 0x20, 0xba, 0x8c, 0x70, 0x0f, 0x88,
	0x00, 0x55, 0x87, 0x9c, 0xd2, 0x5e, 0xa1, 0xd0, 0xa1, 0xb7, 0x4b, 0xe8, 0x16, 0x0a, 0xce, 0x9f,
	0x51, 0x5a, 0x5b, 0x53, 0xa8, 0x7b, 0xa8, 0xfb, 0x8c, 0xf0, 0x28, 0x0d, 0x50, 0x71, 0xe5, 0x4b,
	0x1a, 0xd5, 0x73, 0x5d, 0xad, 0x4a, 0xf9, 0x13, 0x80, 0x60, 0x77, 0x85, 0x70, 0xab, 0x37, 0xd7,
	0x2c, 0x94, 0xfd, 0xf0, 0x1e, 0x18, 0xca, 0x08, 0x27, 0x78, 0x4d, 0x10, 0x31, 0xc6, 0x9f, 0x69,
	0x5d, 0x42, 0x3d, 0x64, 0x84, 0xfb, 0xd8, 0xa6, 0xd3, 0x9a, 0x98, 0xbc, 0x1d, 0x13, 0xed, 0xfa,
	0x08, 0xf3, 0x56, 0xcc, 0x57, 0xfa, 0xff, 0x3f, 0x5b, 0x03, 0xf4, 0x80, 0x11, 0x1e, 0x8e, 0xcf,
	0xf7, 0x87, 0x98, 0x36, 0xe0, 0x17, 0xe0, 0x70, 0x8e, 0xc7, 0x66, 0x4f, 0x74, 0x96, 0x9f, 0x69,
	0xb8, 0x11, 0xa6, 0x10, 0xcd, 0x03, 0xfd, 0x07, 0x56, 0x6f, 0xf6, 0xad, 0xee, 0x10, 0xc0, 0xb5,
	0x70, 0x1b, 0x85, 0x6e, 0xb4, 0x39, 0xd8, 0x8c, 0x61, 0x5d, 0x28, 0x3b, 0x19, 0x83, 0xcb, 0x11,
	0x23, 0xbc, 0x8f, 0x63, 0x40, 0xb5, 0x85, 0x34, 0x8f, 0x74, 0xcc, 0x08, 0x3f, 0x7c, 0x42, 0xe0,
	0x95, 0x2e, 0x7e, 0x79, 0xf4, 0xe4, 0xf9, 0xb0, 0xf8, 0x94, 0x06, 0x6e, 0x81, 0x33, 0x25, 0x2a,
	0x5c, 0xf5, 0x20, 0xf5, 0x9d, 0x70, 0x2d, 0x2a, 0x19, 0x0f, 0xa9, 0x5f, 0x49, 0x2b, 0x72, 0x61,
	0x05, 0x6c, 0x74, 0x94, 0xee, 0xbe, 0xe3, 0x09, 0xed, 0x61, 0x33, 0x75, 0x42, 0xa0, 0xb9, 0x97,
	0xcf, 0x2d, 0x3b, 0x76, 0xb4, 0x25, 0xa7, 0x57, 0x34, 0xd1, 0xe6, 0xbe, 0x05, 0x8e, 0x16, 0x46,
	0x54, 0xf2, 0x41, 0x9b, 0x1f, 0xd3, 0xa8, 0x75, 0xa2, 0xbe, 0xf1, 0xbe, 0xc7, 0xf0, 0x5f, 0xbe,
	0x6b, 0x9b, 0xfe, 0xf1, 0xbc, 0x59, 0x17, 0xe4, 0xc9, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x45,
	0xc9, 0x8b, 0x98, 0x04, 0x04, 0x00, 0x00,
}
