package main

import (
	"log"
	"time"
	pb "simple-tensorflow-serving/predict_proto/proto/tensorflow_s"

	"github.com/cihub/seelog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"simple-tensorflow-serving/tfService"
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

	for i := 0; i < 100; i++ {
		t1 := time.Now()
		_, reqErr := client.Predict(context.Background(), &req)
		if nil != reqErr {
			seelog.Errorf("req err=%v", reqErr)
			return
		}
		seelog.Tracef("net pretic costime=%v",time.Since(t1))
	}

	// seelog.Tracef("rsp=%v", rsp)

}

func test2() {

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

	tfs := tfService.NewTfService1()

	for i := 0; i < 100; i++ {
		_, reqErr := tfs.Predict(context.Background(), &req)
		if nil != reqErr {
			seelog.Errorf("req err=%v", reqErr)
			return
		}
	}

	// for i := 0; i < 100; i++ {
	// 	_, reqErr := client.Predict(context.Background(), &req)
	// 	if nil != reqErr {
	// 		seelog.Errorf("req err=%v", reqErr)
	// 		return
	// 	}
	// }

	// seelog.Tracef("rsp=%v", rsp)

}

func mymodelTest() {
	scopname := "mytest/"
	req := pb.PredictRequest{}
	req.Inputs = make(map[string]*pb.TensorProto)
	req.ModelSpec = &pb.ModelSpec{Name: "golangtestmodel", SignatureName: "mytag"}
	tensor1 := pb.TensorProto{}
	tensor1.Dtype = pb.DataType_DT_FLOAT
	tensor1.TensorShape = &pb.TensorShapeProto{}
	tensor1.TensorShape.Dim = append(tensor1.TensorShape.Dim, &pb.TensorShapeProto_Dim{Size: 2}, &pb.TensorShapeProto_Dim{Size: 3})
	arr1 := [...]float32{1, 1, 1, 2, 2, 2}
	tmpslice := arr1[:]
	tensor1.FloatVal = tmpslice

	arr2 := [...]float32{3, 3, 3, 4, 4, 4}
	tmpslice2 := arr2[:]
	tensor2 := tensor1
	tensor2.FloatVal = tmpslice2

	req.Inputs[scopname+"input1"] = &tensor1
	req.Inputs[scopname+"input2"] = &tensor2
	req.OutputFilter = append(req.OutputFilter, scopname+"myadd", scopname+"mymatmul")

	tfs := tfService.NewTfService1()

	rsp, reqErr := tfs.Predict(context.Background(), &req)
	if nil != reqErr {
		seelog.Errorf("req err=%v", reqErr)
		return
	}

	for k, v := range rsp.Outputs {
		seelog.Tracef("one tensor,key=%v,v=%v", k, v)
	}

}

func mymodelTest2() {
	scopname := "mytest/"
	req := pb.PredictRequest{}
	req.Inputs = make(map[string]*pb.TensorProto)
	req.ModelSpec = &pb.ModelSpec{Name: "golangtestmodel", SignatureName: "mytag"}
	tensor1 := pb.TensorProto{}
	tensor1.Dtype = pb.DataType_DT_FLOAT
	tensor1.TensorShape = &pb.TensorShapeProto{}
	tensor1.TensorShape.Dim = append(tensor1.TensorShape.Dim, &pb.TensorShapeProto_Dim{Size: 2}, &pb.TensorShapeProto_Dim{Size: 3})
	arr1 := [...]float32{1, 1, 1, 2, 2, 2}
	tmpslice := arr1[:]
	tensor1.FloatVal = tmpslice

	arr2 := [...]float32{3, 3, 3, 4, 4, 4}
	tmpslice2 := arr2[:]
	tensor2 := tensor1
	tensor2.FloatVal = tmpslice2

	req.Inputs[scopname+"input1"] = &tensor1
	req.Inputs[scopname+"input2"] = &tensor2
	req.OutputFilter = append(req.OutputFilter, scopname+"myadd", scopname+"mymatmul")

	tfs := tfService.NewTfService1()

	rsp, reqErr := tfs.Predict(context.Background(), &req)
	if nil != reqErr {
		seelog.Errorf("req err=%v", reqErr)
		return
	}

	for k, v := range rsp.Outputs {
		seelog.Tracef("one tensor,key=%v,v=%v", k, v)
	}

}

func LOG_INIT(logConfigFile string) {
	logger, err := seelog.LoggerFromConfigAsFile(logConfigFile)
	if nil != err {
		seelog.Critical("err parsing config log file", err)
	}
	seelog.ReplaceLogger(logger)
}

func main() {
	LOG_INIT("linux_log_conf.xml")
	seelog.Infof("enter2")
	defer seelog.Flush()
	test1()
	// test2()
	// mymodelTest()
	// mymodelTest2()

}
