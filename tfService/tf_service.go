package tfService

import(
	"context"
	"github.com/cihub/seelog"
	pb "simple-tensorflow-serving/predict_proto/proto/tensorflow_s"
)

type  TfService1 struct {
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
	return nil,nil
}

