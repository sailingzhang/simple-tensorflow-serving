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
	var err error
	rsp := pb.PredictResponse{}
	modelname := req.ModelSpec.GetName()
	tag := req.GetModelSpec().GetSignatureName()
	model,ok := this.modelMap[modelname]
	if !ok{
		model, err = tf.LoadSavedModel(modelname, []string{tag}, nil)
		if err != nil {
			seelog.Errorf("open model=%v err, err=%v",)
			return nil,nil
		}
	}
	inputTesorMap := make(map[tf.Output]*tf.Tensor)
	outputTesorMap := make(map[tf.Output]*tf.Tensor)
	for inkey,invalue :=range req.Inputs{
		// dimension := value.GetTensorShape().GetDim()
	
		inputTesorMap[model.Graph.Operation(inkey).Output(0)],_= tf.NewTensor(invalue.GetFloatVal())
	}
	for outValue := range req.OutputFilter{
		outputTesorMap
	}
	_, err = model.Session.Run(
		inputTesorMap,
		[]tf.Output{
			model.Graph.Operation("embeddings").Output(0), // Replace this with your output layer name
		},
		nil,
	)

	
	// defer model.Session.Close()
	return nil,nil
}

