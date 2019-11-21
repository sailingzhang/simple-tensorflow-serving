package tfService

import (
	"bytes"
	"context"
	"github.com/cihub/seelog"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"

	pb "simple-tensorflow-serving/predict_proto/proto/tensorflow_s"
)

func getSliceFromTensorProto(req *pb.TensorProto) interface{} {
	switch req.Dtype {
	case pb.DataType_DT_FLOAT:
		return req.GetFloatVal()
	case pb.DataType_DT_DOUBLE:
		return req.GetDoubleVal()
	case pb.DataType_DT_BOOL:
		return req.GetBoolVal()
	case pb.DataType_DT_INT32:
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
func getShapeFroTensorProto(req *pb.TensorProto) []int64 {
	rsp := []int64{}
	dim := req.GetTensorShape().GetDim()
	for _, v := range dim {
		rsp = append(rsp, v.Size)
	}
	return rsp
}

func getTensorProtoFromTensor(req *tf.Tensor) *pb.TensorProto {
	rsp := pb.TensorProto{}
	shapeproto := pb.TensorShapeProto{}
	for _, v := range req.Shape() {
		dim := pb.TensorShapeProto_Dim{}
		dim.Size = v
		shapeproto.Dim = append(shapeproto.Dim, &dim)
	}
	datetypeProto := getTensorProtoTypeFormTensor(req)

	req.DataType()

	req.Shape()

}

type TfService1 struct {
	modelMap map[string]*tf.SavedModel //should use  sync.map later,grpc work in multiple goroutine
	// modelMap sync.Map[string]*tf.SavedModel

}

func (this *TfService1) Classify(ctx context.Context, req *pb.ClassificationRequest) (*pb.ClassificationResponse, error) {
	seelog.Tracef("dd")
	return nil, nil
}

func (this *TfService1) Regress(ctx context.Context, req *pb.RegressionRequest) (*pb.RegressionResponse, error) {
	// seelog.
	// see

	seelog.Tracef("dgd")
	seelog.Debugf("dsg")
	return nil, nil

}

func (this *TfService1) Predict(ctx context.Context, req *pb.PredictRequest) (*pb.PredictResponse, error) {
	var err error
	rsp := pb.PredictResponse{}

	modelname := req.ModelSpec.GetName()
	tag := req.GetModelSpec().GetSignatureName()
	model, ok := this.modelMap[modelname]
	if !ok {
		model, err = tf.LoadSavedModel(modelname, []string{tag}, nil)
		if err != nil {
			seelog.Errorf("open model=%v err, err=%v")
			return &rsp, nil
		}
	}
	inputTesorMap := make(map[tf.Output]*tf.Tensor)
	outputTesorSlice := []tf.Output{}
	for inkey, invalue := range req.Inputs {
		// dimension := value.GetTensorShape().GetDim()
		// tmpTensor, _ := tf.NewTensor(invalue.GetDtype().EnumDescriptor())
		shapeSlice := getShapeFroTensorProto(invalue)
		datatype := invalue.Dtype
		tensorSlice := getSliceFromTensorProto(invalue)
		tmpTensor, _ := tf.NewTensor(tensorSlice)
		buf := new(bytes.Buffer)
		if _, err := tmpTensor.WriteContentsTo(buf); nil != err {
			seelog.Errorf("write buf err")
			return &rsp, nil
		}
		fakeInput, readErr := tf.ReadTensor(tmpTensor.DataType(), shapeSlice, buf)
		if nil != readErr {
			seelog.Errorf("read tensor err")
			return &rsp, nil
		}
		inputTesorMap[model.Graph.Operation(inkey).Output(0)] = fakeInput
	}

	for _, outvalue := range req.GetOutputFilter() {
		outputTesorSlice = append(outputTesorSlice, model.Graph.Operation(outvalue).Output(0))
	}
	ret, err := model.Session.Run(inputTesorMap, outputTesorSlice, nil)
	for _, v := range ret {

	}

	// defer model.Session.Close()
	return nil, nil
}
