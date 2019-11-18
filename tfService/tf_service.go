package tfService

import(
	"context"
	"github.com/cihub/seelog"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"github.com/tensorflow/tensorflow/tensorflow/go/op"

	pb "simple-tensorflow-serving/predict_proto/proto/tensorflow_s"
)

type  TfService1 struct {
	modelMap map[string]*tf.SavedModel
}

func (this * TfService1)Classify(ctx context.Context,req *pb.ClassificationRequest) (*pb.ClassificationResponse, error){
	seelog.Tracef("dd")
	return nil,nil
}

func (this * TfService1)Regress(ctx context.Context, req *pb.RegressionRequest) (*pb.RegressionResponse, error){
	// seelog.
	// see
	
	seelog.Tracef("dgd")
	seelog.Debugf("dsg")

	return nil,nil

}

func (this * TfService1)Predict(ctx context.Context, req *pb.PredictRequest) (*pb.PredictResponse, error){
	rsp := pb.PredictResponse{}
	modelname := req.ModelSpec.GetName()
	tag := req.GetModelSpec().GetSignatureName()
	model, err := tf.LoadSavedModel(modelname, []string{tag}, nil)
	if err != nil {
		seelog.Errorf("open model=%v err, err=%v",)
		return nil,nil
	}
	
	defer model.Session.Close()
	return nil,nil
}

