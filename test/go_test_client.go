package main

import (
	"log"
	pb "simple-tensorflow-serving/predict_proto/proto/tensorflow_s"

	"github.com/cihub/seelog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func test1() {
	address := "127.0.0.1:2233"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewPredictionServiceClient(conn)
	req := pb.PredictRequest{}
	req.Inputs = make(map[string]*pb.TensorProto)
	req.ModelSpec = &pb.ModelSpec{Name: "/tmp/emdmodel/10001", SignatureName: "serve"}
	tensor1 := pb.TensorProto{}
	tensor2 := pb.TensorProto{}
	tensor1.Dtype = pb.DataType_DT_FLOAT
	tensor1.TensorShape = &pb.TensorShapeProto{}
	tensor1.TensorShape.Dim = append(tensor1.TensorShape.Dim, &pb.TensorShapeProto_Dim{Size: 10}, &pb.TensorShapeProto_Dim{Size: 160}, &pb.TensorShapeProto_Dim{Size: 160}, &pb.TensorShapeProto_Dim{Size: 3})
	// {&pb.TensorShapeProto_Dim{Size: 10}, &pb.TensorShapeProto_Dim{Size: 160}, &pb.TensorShapeProto_Dim{Size: 160}, &pb.TensorShapeProto_Dim{Size: 3}}
	// tensor1.FloatVal = append(tensor1.FloatVal, [10 * 160 * 160 * 3]float32{}[:]...)
	var arr [10 * 160 * 160 * 3]float32
	tmpslice := arr[:]
	tensor1.FloatVal = append(tensor1.FloatVal, tmpslice...)
	tensor2.Dtype = pb.DataType_DT_BOOL
	tensor2.TensorShape = &pb.TensorShapeProto{}
	// tensor2.TensorShape.Dim =
	tensor2.BoolVal = append(tensor2.BoolVal, false)

	req.Inputs["input"] = &tensor1
	req.Inputs["phase_train"] = &tensor2
	req.OutputFilter = append(req.OutputFilter, "embeddings")

	rsp, reqErr := client.Predict(context.Background(), &req)
	if nil != reqErr {
		seelog.Errorf("req err=%v", reqErr)
		return
	}
	seelog.Tracef("rsp=%v", rsp)

}

func main() {
	seelog.Infof("enter")
	defer seelog.Flush()
	test1()

}
