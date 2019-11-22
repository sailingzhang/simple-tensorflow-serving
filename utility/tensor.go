package utility

import (
	"bytes"

	"github.com/cihub/seelog"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"

	pb "simple-tensorflow-serving/predict_proto/proto/tensorflow_s"
)

func getSliceFromTensorProto(req *pb.TensorProto) interface{} {
	switch req.Dtype {
	case pb.DataType_DT_FLOAT:
		seelog.Debugf("float type")
		return req.GetFloatVal()
	case pb.DataType_DT_DOUBLE:
		seelog.Debugf("double type")
		return req.GetDoubleVal()
	case pb.DataType_DT_BOOL:
		seelog.Debugf("bool type")
		return req.GetBoolVal()
	case pb.DataType_DT_INT32:
		seelog.Debugf("int type")
		return req.GetIntVal()
	default:
		seelog.Errorf("not support such type={}", req.Dtype)
		return nil
	}
	return nil
}

func getTensorProtoTypeFormTensor(req *tf.Tensor) pb.DataType {
	switch req.DataType() {
	case tf.Float:
		return pb.DataType_DT_FLOAT
	case tf.Double:
		return pb.DataType_DT_DOUBLE
	case tf.Bool:
		return pb.DataType_DT_BOOL
	case tf.Int32:
		return pb.DataType_DT_INT32
	default:
		seelog.Errorf("no support such type=%v,defualt float", req.DataType())
		return pb.DataType_DT_FLOAT
	}
}
func getShapeFromTensorProto(req *pb.TensorProto) []int64 {
	rsp := []int64{}
	dim := req.GetTensorShape().GetDim()
	for _, v := range dim {
		rsp = append(rsp, v.Size)
	}
	return rsp
}

func PrototensorToTensor(req *pb.TensorProto) (*tf.Tensor, error) {
	shapeSlice := getShapeFromTensorProto(req)
	tensorSlice := getSliceFromTensorProto(req)
	if nil == tensorSlice {
		return nil, seelog.Errorf("get slice err")
	}
	tmpTensor, newErr := tf.NewTensor(tensorSlice)
	if nil != newErr {
		return nil, seelog.Errorf("new tensor err=%v", newErr)
	}
	buf := new(bytes.Buffer)
	if _, err := tmpTensor.WriteContentsTo(buf); nil != err {
		return nil, seelog.Errorf("write buf err")
	}
	newTensor, readErr := tf.ReadTensor(tmpTensor.DataType(), shapeSlice, buf)
	if nil != readErr {
		return nil, seelog.Errorf("read tensor err")
	}
	return newTensor, nil
}

func TensorToProtoTensor(req *tf.Tensor) (*pb.TensorProto, error) {
	rsp := pb.TensorProto{}

	shapeproto := pb.TensorShapeProto{}
	sliceLen := int64(1)
	for _, len := range req.Shape() {
		dim := pb.TensorShapeProto_Dim{}
		dim.Size = len
		Dim := []*pb.TensorShapeProto_Dim{&dim}
		shapeproto.Dim = append(Dim, shapeproto.Dim...)
		sliceLen *= len
	}
	rsp.TensorShape = &shapeproto
	converShape := []int64{sliceLen}
	seelog.Debugf("converShape=%v", converShape)
	buf := new(bytes.Buffer)
	if _, err := req.WriteContentsTo(buf); nil != err {
		return nil, seelog.Errorf("write buf err")
	}
	newTensor, readErr := tf.ReadTensor(req.DataType(), converShape, buf)
	if nil != readErr {
		return nil, seelog.Errorf("read tensor err")
	}

	if tf.Float == newTensor.DataType() {
		v, ok := newTensor.Value().([]float32)
		if !ok {
			return nil, seelog.Errorf("assert bad")
		}
		rsp.Dtype = pb.DataType_DT_FLOAT
		rsp.FloatVal = v
	} else if tf.Double == newTensor.DataType() {
		v, ok := newTensor.Value().([]float64)
		if !ok {
			return nil, seelog.Errorf("assert bad")
		}
		rsp.Dtype = pb.DataType_DT_DOUBLE
		rsp.DoubleVal = v
	} else if tf.Bool == newTensor.DataType() {
		v, ok := newTensor.Value().([]bool)
		if !ok {
			return nil, seelog.Errorf("assert bad")
		}
		rsp.Dtype = pb.DataType_DT_BOOL
		rsp.BoolVal = v
	} else if tf.Int32 == newTensor.DataType() {
		v, ok := newTensor.Value().([]int32)
		if !ok {
			return nil, seelog.Errorf("assert bad")
		}
		rsp.Dtype = pb.DataType_DT_INT32
		rsp.IntVal = v
	} else {
		return nil, seelog.Errorf("no support such type=%v", req.DataType())
	}
	return &rsp, nil
}
