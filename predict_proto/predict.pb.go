// Code generated by protoc-gen-go. DO NOT EDIT.
// source: predict.proto

package framework

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// PredictRequest specifies which TensorFlow model to run, as well as
// how inputs are mapped to tensors and how outputs are filtered before
// returning to user.
type PredictRequest struct {
	// Model Specification. If version is not specified, will use the latest
	// (numerical) version.
	ModelSpec *ModelSpec `protobuf:"bytes,1,opt,name=model_spec,json=modelSpec" json:"model_spec"`
	// Input tensors.
	// Names of input tensor are alias names. The mapping from aliases to real
	// input tensor names is stored in the SavedModel export as a prediction
	// SignatureDef under the 'inputs' field.
	Inputs map[string]*TensorProto `protobuf:"bytes,2,rep,name=inputs" json:"inputs" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Output filter.
	// Names specified are alias names. The mapping from aliases to real output
	// tensor names is stored in the SavedModel export as a prediction
	// SignatureDef under the 'outputs' field.
	// Only tensors specified here will be run/fetched and returned, with the
	// exception that when none is specified, all tensors specified in the
	// named signature will be run/fetched and returned.
	OutputFilter []string `protobuf:"bytes,3,rep,name=output_filter,json=outputFilter" json:"output_filter"`
}

func (m *PredictRequest) Reset()                    { *m = PredictRequest{} }
func (m *PredictRequest) String() string            { return proto.CompactTextString(m) }
func (*PredictRequest) ProtoMessage()               {}
func (*PredictRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *PredictRequest) GetModelSpec() *ModelSpec {
	if m != nil {
		return m.ModelSpec
	}
	return nil
}

func (m *PredictRequest) GetInputs() map[string]*TensorProto {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *PredictRequest) GetOutputFilter() []string {
	if m != nil {
		return m.OutputFilter
	}
	return nil
}

// Response for PredictRequest on successful run.
type PredictResponse struct {
	// Effective Model Specification used to process PredictRequest.
	ModelSpec *ModelSpec `protobuf:"bytes,2,opt,name=model_spec,json=modelSpec" json:"model_spec"`
	// Output tensors.
	Outputs map[string]*TensorProto `protobuf:"bytes,1,rep,name=outputs" json:"outputs" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *PredictResponse) Reset()                    { *m = PredictResponse{} }
func (m *PredictResponse) String() string            { return proto.CompactTextString(m) }
func (*PredictResponse) ProtoMessage()               {}
func (*PredictResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *PredictResponse) GetModelSpec() *ModelSpec {
	if m != nil {
		return m.ModelSpec
	}
	return nil
}

func (m *PredictResponse) GetOutputs() map[string]*TensorProto {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func init() {
	proto.RegisterType((*PredictRequest)(nil), "tensorflow.serving.PredictRequest")
	proto.RegisterType((*PredictResponse)(nil), "tensorflow.serving.PredictResponse")
}

func init() { proto.RegisterFile("predict.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0xd9, 0x04, 0x2b, 0x99, 0xa4, 0x2a, 0x7b, 0x31, 0x04, 0x84, 0x50, 0x2f, 0xb9, 0xb8,
	0x48, 0xbd, 0x88, 0x78, 0x12, 0x2c, 0x28, 0x88, 0x65, 0xeb, 0xbd, 0x68, 0x3a, 0x95, 0x60, 0x9a,
	0x5d, 0xf7, 0x4f, 0xa5, 0x9f, 0xc2, 0xaf, 0xea, 0xd1, 0xa3, 0x64, 0xd7, 0x4a, 0xfc, 0x83, 0x20,
	0x78, 0x9b, 0x4c, 0xde, 0x9b, 0x79, 0xf3, 0x63, 0xa1, 0x2f, 0x15, 0xce, 0xaa, 0xd2, 0x30, 0xa9,
	0x84, 0x11, 0x94, 0x1a, 0x6c, 0xb4, 0x50, 0xf3, 0x5a, 0x3c, 0x31, 0x8d, 0x6a, 0x59, 0x35, 0xf7,
	0x59, 0xe2, 0x7b, 0x5e, 0x91, 0xc5, 0x0b, 0x31, 0xc3, 0xda, 0x7f, 0x0c, 0x9e, 0x03, 0xd8, 0x1a,
	0xfb, 0x01, 0x1c, 0x1f, 0x2d, 0x6a, 0x43, 0x4f, 0x01, 0x9c, 0x62, 0xaa, 0x25, 0x96, 0x29, 0xc9,
	0x49, 0x11, 0x0f, 0xf7, 0xd8, 0xf7, 0xb1, 0xec, 0xaa, 0x55, 0x4d, 0x24, 0x96, 0x3c, 0x5a, 0xac,
	0x4b, 0x3a, 0x82, 0x5e, 0xd5, 0x48, 0x6b, 0x74, 0x1a, 0xe4, 0x61, 0x11, 0x0f, 0xd9, 0x4f, 0xce,
	0xcf, 0x1b, 0xd9, 0x85, 0x33, 0x9c, 0x37, 0x46, 0xad, 0xf8, 0xbb, 0x9b, 0xee, 0x43, 0x5f, 0x58,
	0x23, 0xad, 0x99, 0xce, 0xab, 0xda, 0xa0, 0x4a, 0xc3, 0x3c, 0x2c, 0x22, 0x9e, 0xf8, 0xe6, 0xc8,
	0xf5, 0x32, 0x0e, 0x71, 0xc7, 0x4b, 0x77, 0x20, 0x7c, 0xc0, 0x95, 0x8b, 0x1c, 0xf1, 0xb6, 0xa4,
	0x07, 0xb0, 0xb1, 0xbc, 0xad, 0x2d, 0xa6, 0x81, 0x3b, 0x63, 0xb7, 0x1b, 0xe6, 0xc6, 0x95, 0xe3,
	0x16, 0x03, 0xf7, 0xaa, 0x93, 0xe0, 0x98, 0x0c, 0x5e, 0x08, 0x6c, 0x7f, 0xe4, 0xd3, 0x52, 0x34,
	0x1a, 0xbf, 0x20, 0x09, 0xfe, 0x88, 0xe4, 0x12, 0x36, 0x7d, 0x6a, 0x9d, 0x12, 0xc7, 0xe4, 0xf0,
	0x57, 0x26, 0x7e, 0x27, 0xbb, 0xf6, 0x16, 0x4f, 0x65, 0x3d, 0x20, 0x9b, 0x40, 0xd2, 0xfd, 0xf1,
	0x2f, 0x27, 0x9f, 0x85, 0xaf, 0x84, 0xdc, 0xf5, 0xdc, 0x83, 0x38, 0x7a, 0x0b, 0x00, 0x00, 0xff,
	0xff, 0x9a, 0xb6, 0x05, 0x6a, 0x50, 0x02, 0x00, 0x00,
}