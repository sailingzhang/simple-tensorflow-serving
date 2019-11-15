// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types.proto

package tensorflow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// (== suppress_warning documentation-presence ==)
// LINT.IfChange
type DataType int32

const (
	// Not a legal value for DataType.  Used to indicate a DataType field
	// has not been set.
	DataType_DT_INVALID DataType = 0
	// Data types that all computation devices are expected to be
	// capable to support.
	DataType_DT_FLOAT      DataType = 1
	DataType_DT_DOUBLE     DataType = 2
	DataType_DT_INT32      DataType = 3
	DataType_DT_UINT8      DataType = 4
	DataType_DT_INT16      DataType = 5
	DataType_DT_INT8       DataType = 6
	DataType_DT_STRING     DataType = 7
	DataType_DT_COMPLEX64  DataType = 8
	DataType_DT_INT64      DataType = 9
	DataType_DT_BOOL       DataType = 10
	DataType_DT_QINT8      DataType = 11
	DataType_DT_QUINT8     DataType = 12
	DataType_DT_QINT32     DataType = 13
	DataType_DT_BFLOAT16   DataType = 14
	DataType_DT_QINT16     DataType = 15
	DataType_DT_QUINT16    DataType = 16
	DataType_DT_UINT16     DataType = 17
	DataType_DT_COMPLEX128 DataType = 18
	DataType_DT_HALF       DataType = 19
	DataType_DT_RESOURCE   DataType = 20
	DataType_DT_VARIANT    DataType = 21
	DataType_DT_UINT32     DataType = 22
	DataType_DT_UINT64     DataType = 23
	// Do not use!  These are only for parameters.  Every enum above
	// should have a corresponding value below (verified by types_test).
	DataType_DT_FLOAT_REF      DataType = 101
	DataType_DT_DOUBLE_REF     DataType = 102
	DataType_DT_INT32_REF      DataType = 103
	DataType_DT_UINT8_REF      DataType = 104
	DataType_DT_INT16_REF      DataType = 105
	DataType_DT_INT8_REF       DataType = 106
	DataType_DT_STRING_REF     DataType = 107
	DataType_DT_COMPLEX64_REF  DataType = 108
	DataType_DT_INT64_REF      DataType = 109
	DataType_DT_BOOL_REF       DataType = 110
	DataType_DT_QINT8_REF      DataType = 111
	DataType_DT_QUINT8_REF     DataType = 112
	DataType_DT_QINT32_REF     DataType = 113
	DataType_DT_BFLOAT16_REF   DataType = 114
	DataType_DT_QINT16_REF     DataType = 115
	DataType_DT_QUINT16_REF    DataType = 116
	DataType_DT_UINT16_REF     DataType = 117
	DataType_DT_COMPLEX128_REF DataType = 118
	DataType_DT_HALF_REF       DataType = 119
	DataType_DT_RESOURCE_REF   DataType = 120
	DataType_DT_VARIANT_REF    DataType = 121
	DataType_DT_UINT32_REF     DataType = 122
	DataType_DT_UINT64_REF     DataType = 123
)

var DataType_name = map[int32]string{
	0:   "DT_INVALID",
	1:   "DT_FLOAT",
	2:   "DT_DOUBLE",
	3:   "DT_INT32",
	4:   "DT_UINT8",
	5:   "DT_INT16",
	6:   "DT_INT8",
	7:   "DT_STRING",
	8:   "DT_COMPLEX64",
	9:   "DT_INT64",
	10:  "DT_BOOL",
	11:  "DT_QINT8",
	12:  "DT_QUINT8",
	13:  "DT_QINT32",
	14:  "DT_BFLOAT16",
	15:  "DT_QINT16",
	16:  "DT_QUINT16",
	17:  "DT_UINT16",
	18:  "DT_COMPLEX128",
	19:  "DT_HALF",
	20:  "DT_RESOURCE",
	21:  "DT_VARIANT",
	22:  "DT_UINT32",
	23:  "DT_UINT64",
	101: "DT_FLOAT_REF",
	102: "DT_DOUBLE_REF",
	103: "DT_INT32_REF",
	104: "DT_UINT8_REF",
	105: "DT_INT16_REF",
	106: "DT_INT8_REF",
	107: "DT_STRING_REF",
	108: "DT_COMPLEX64_REF",
	109: "DT_INT64_REF",
	110: "DT_BOOL_REF",
	111: "DT_QINT8_REF",
	112: "DT_QUINT8_REF",
	113: "DT_QINT32_REF",
	114: "DT_BFLOAT16_REF",
	115: "DT_QINT16_REF",
	116: "DT_QUINT16_REF",
	117: "DT_UINT16_REF",
	118: "DT_COMPLEX128_REF",
	119: "DT_HALF_REF",
	120: "DT_RESOURCE_REF",
	121: "DT_VARIANT_REF",
	122: "DT_UINT32_REF",
	123: "DT_UINT64_REF",
}
var DataType_value = map[string]int32{
	"DT_INVALID":        0,
	"DT_FLOAT":          1,
	"DT_DOUBLE":         2,
	"DT_INT32":          3,
	"DT_UINT8":          4,
	"DT_INT16":          5,
	"DT_INT8":           6,
	"DT_STRING":         7,
	"DT_COMPLEX64":      8,
	"DT_INT64":          9,
	"DT_BOOL":           10,
	"DT_QINT8":          11,
	"DT_QUINT8":         12,
	"DT_QINT32":         13,
	"DT_BFLOAT16":       14,
	"DT_QINT16":         15,
	"DT_QUINT16":        16,
	"DT_UINT16":         17,
	"DT_COMPLEX128":     18,
	"DT_HALF":           19,
	"DT_RESOURCE":       20,
	"DT_VARIANT":        21,
	"DT_UINT32":         22,
	"DT_UINT64":         23,
	"DT_FLOAT_REF":      101,
	"DT_DOUBLE_REF":     102,
	"DT_INT32_REF":      103,
	"DT_UINT8_REF":      104,
	"DT_INT16_REF":      105,
	"DT_INT8_REF":       106,
	"DT_STRING_REF":     107,
	"DT_COMPLEX64_REF":  108,
	"DT_INT64_REF":      109,
	"DT_BOOL_REF":       110,
	"DT_QINT8_REF":      111,
	"DT_QUINT8_REF":     112,
	"DT_QINT32_REF":     113,
	"DT_BFLOAT16_REF":   114,
	"DT_QINT16_REF":     115,
	"DT_QUINT16_REF":    116,
	"DT_UINT16_REF":     117,
	"DT_COMPLEX128_REF": 118,
	"DT_HALF_REF":       119,
	"DT_RESOURCE_REF":   120,
	"DT_VARIANT_REF":    121,
	"DT_UINT32_REF":     122,
	"DT_UINT64_REF":     123,
}

func (x DataType) String() string {
	return proto.EnumName(DataType_name, int32(x))
}
func (DataType) EnumDescriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

func init() {
	proto.RegisterEnum("tensorflow.DataType", DataType_name, DataType_value)
}

func init() { proto.RegisterFile("types.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x93, 0xdb, 0x8e, 0xd3, 0x40,
	0x0c, 0x86, 0x29, 0x87, 0x6e, 0xd7, 0xe9, 0xc1, 0x9d, 0xdd, 0x05, 0x9e, 0x81, 0x8b, 0xa2, 0xb4,
	0xd5, 0xa8, 0x57, 0x48, 0xe9, 0x26, 0x85, 0x48, 0x21, 0xd3, 0x66, 0xa7, 0x2b, 0xc4, 0x4d, 0x54,
	0xa4, 0x74, 0x81, 0x3d, 0xa4, 0x24, 0x81, 0x52, 0x78, 0x23, 0x9e, 0x90, 0x4b, 0x54, 0xc7, 0xed,
	0x0c, 0x97, 0xfe, 0xec, 0xf9, 0xfd, 0x8f, 0x2d, 0x83, 0x53, 0xed, 0x36, 0x59, 0x39, 0xd8, 0x14,
	0x79, 0x95, 0x0b, 0xa8, 0xb2, 0x87, 0x32, 0x2f, 0xd6, 0x77, 0xf9, 0xf6, 0xd5, 0x9f, 0x26, 0xb4,
	0xfc, 0x55, 0xb5, 0xd2, 0xbb, 0x4d, 0x26, 0xba, 0x00, 0xbe, 0x4e, 0xc3, 0xf8, 0xda, 0x8b, 0x42,
	0x1f, 0x1f, 0x89, 0x36, 0xb4, 0x7c, 0x9d, 0xce, 0x22, 0xe5, 0x69, 0x6c, 0x88, 0x0e, 0x9c, 0xfa,
	0x3a, 0xf5, 0xd5, 0x72, 0x1a, 0x05, 0xf8, 0x98, 0x93, 0x61, 0xac, 0x47, 0x43, 0x7c, 0xc2, 0xd1,
	0x32, 0x8c, 0xf5, 0x04, 0x9f, 0x9a, 0x9c, 0x2b, 0xf1, 0x99, 0x70, 0xe0, 0xa4, 0x8e, 0x26, 0xd8,
	0x64, 0x95, 0x2b, 0x9d, 0x84, 0xf1, 0x5b, 0x3c, 0x11, 0x08, 0x6d, 0x5f, 0xa7, 0x97, 0xea, 0xfd,
	0x3c, 0x0a, 0x3e, 0xc8, 0x31, 0xb6, 0xcc, 0x5b, 0x39, 0xc6, 0x53, 0x7e, 0x3b, 0x55, 0x2a, 0x42,
	0xe0, 0xd4, 0x82, 0x94, 0x1c, 0x56, 0x5a, 0xd4, 0x3d, 0xdb, 0x87, 0xb0, 0x36, 0xd4, 0x11, 0x3d,
	0x70, 0xf6, 0x0f, 0xc9, 0xbc, 0x2b, 0xb1, 0x6b, 0xe5, 0x5d, 0x89, 0x3d, 0xfe, 0x2b, 0xbd, 0x76,
	0x25, 0x22, 0xa7, 0x39, 0xec, 0x8b, 0x3e, 0x74, 0x8c, 0x2f, 0x77, 0x38, 0x41, 0xc1, 0x56, 0xde,
	0x79, 0xd1, 0x0c, 0xcf, 0x58, 0x3e, 0x09, 0xae, 0xd4, 0x32, 0xb9, 0x0c, 0xf0, 0x9c, 0xf5, 0xae,
	0xbd, 0x24, 0xf4, 0x62, 0x8d, 0x17, 0x96, 0xde, 0x68, 0x88, 0xcf, 0xad, 0x50, 0x8e, 0xf1, 0x05,
	0x7f, 0x9b, 0xcc, 0xa5, 0x49, 0x30, 0xc3, 0x8c, 0x1b, 0xd6, 0xd3, 0x25, 0xb4, 0xe6, 0x22, 0x52,
	0x20, 0x72, 0xc3, 0x84, 0x7e, 0x4c, 0xe4, 0xb3, 0xa9, 0x71, 0x25, 0x91, 0x2f, 0xec, 0xec, 0x58,
	0xf2, 0x95, 0x95, 0xeb, 0x89, 0x13, 0xba, 0x15, 0xe7, 0x80, 0xf6, 0xd4, 0x89, 0xde, 0x19, 0x2d,
	0x26, 0xf7, 0x87, 0x21, 0x2a, 0x15, 0x11, 0x78, 0xe0, 0x92, 0xc5, 0x51, 0x3d, 0x67, 0xf5, 0x85,
	0xf1, 0xb4, 0x39, 0x20, 0x63, 0xfc, 0x9b, 0x38, 0x83, 0x9e, 0xb5, 0x0d, 0x82, 0x85, 0x55, 0xc7,
	0xa8, 0x14, 0x02, 0xba, 0x66, 0x2b, 0xc4, 0x2a, 0x2e, 0xb3, 0xd0, 0x77, 0x71, 0x01, 0xfd, 0xff,
	0xb6, 0x43, 0xf8, 0x07, 0xdb, 0xdd, 0x6f, 0x88, 0xc0, 0x96, 0xdb, 0x1e, 0xb6, 0x44, 0xf0, 0x27,
	0xf7, 0xe0, 0x4d, 0x11, 0xdb, 0x59, 0x3d, 0xd8, 0xf2, 0x2f, 0x0b, 0xf1, 0x38, 0x7e, 0x4f, 0xdf,
	0xc0, 0xcb, 0xbc, 0xb8, 0x19, 0x98, 0xf3, 0x19, 0xac, 0x8b, 0xd5, 0x7d, 0xb6, 0xcd, 0x8b, 0xdb,
	0xa9, 0xb3, 0xbf, 0xa0, 0x72, 0xbe, 0x3f, 0xb0, 0x72, 0xde, 0xf8, 0x88, 0x74, 0x6a, 0xaf, 0x4d,
	0xe9, 0xdf, 0x46, 0xe3, 0x53, 0x93, 0xe0, 0xe8, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x37,
	0x7c, 0xba, 0x8e, 0x03, 0x00, 0x00,
}