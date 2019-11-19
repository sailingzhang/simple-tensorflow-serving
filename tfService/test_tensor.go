package tfService
// package main

import (
	"fmt"
	"bytes"
	"github.com/cihub/seelog"
	tg "github.com/galeone/tfgo"

	"time"

	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"github.com/tensorflow/tensorflow/tensorflow/go/op"
	// "io/ioutil"
)

// func LOG_INIT(logconfigxml string) {
// 	var logConfigFile = ""
// 	dir, fileErr := filepath.Abs(filepath.Dir(os.Args[0]))
// 	if fileErr != nil {
// 		seelog.Errorf("err=%v", fileErr)
// 		logConfigFile = logconfigxml
// 	} else {
// 		logConfigFile = path.Join(dir, logconfigxml)
// 	}
// 	seelog.Tracef("file=%s", logConfigFile)
// 	logger, err := seelog.LoggerFromConfigAsFile(logConfigFile)
// 	if nil != err {
// 		seelog.Critical("err parsing config log file", err)
// 		return
// 	}
// 	seelog.ReplaceLogger(logger)
// }

// func my_tg() {
//     model := tg.LoadModel("test_models/export", []string{"tag"}, nil)

//     fakeInput, _ := tf.NewTensor([1][28][28][1]float32{})
//     results := model.Exec([]tf.Output{
//             model.Op("LeNetDropout/softmax_linear/Identity", 0),
//     }, map[tf.Output]*tf.Tensor{
//             model.Op("input_", 0): fakeInput,
//     })

//     predictions := results[0].Value().([][]float32)
//     fmt.Println(predictions)
// }

func my_tf() {
	seelog.Tracef("enter")
	// replace myModel and myTag with the appropriate exported names in the chestrays-keras-binary-classification.ipynb
	// model, err := tf.LoadSavedModel("/tmp/emdmodel/10001", []string{"myTag"}, nil)
	model, err := tf.LoadSavedModel("/tmp/emdmodel/10001", []string{"serve"}, nil)
	if err != nil {
		fmt.Printf("Error loading saved model: %s\n", err.Error())
		return
	}
	defer model.Session.Close()

	// bytes, err := ioutil.ReadFile("/home/sailingzhang/winshare/develop/source/mygit/tmp/ml_pic/to_compare/to_picface_0.png")
	// if err != nil {
	// 	return
	// }
	// DecodeJpeg uses a scalar String-valued tensor as input.
	// tensor, err := tf.NewTensor(string(bytes))
	// if err != nil {
	// 	return
	// }

	fakeInput, _ := tf.NewTensor([10][160][160][3]float32{})
	istrainInput, _ := tf.NewTensor(true)

	for {
		t1 := time.Now()
		_, err := model.Session.Run(
			map[tf.Output]*tf.Tensor{
				model.Graph.Operation("input").Output(0):       fakeInput,
				model.Graph.Operation("phase_train").Output(0): istrainInput, // Replace this with your input layer name
			},
			[]tf.Output{
				model.Graph.Operation("embeddings").Output(0), // Replace this with your output layer name
			},
			nil,
		)

		if err != nil {
			fmt.Printf("Error running the session with input, err: %s\n", err.Error())
			return
		}
		t2 := time.Now()
		fmt.Printf("tf cost=%v\n", t2.Sub(t1))
		// fmt.Printf("Result value: %v \n", result[0].Value())
	}

}


func my_tf2() {
	seelog.Tracef("enter")
	// replace myModel and myTag with the appropriate exported names in the chestrays-keras-binary-classification.ipynb
	// model, err := tf.LoadSavedModel("/tmp/emdmodel/10001", []string{"myTag"}, nil)
	model, err := tf.LoadSavedModel("/tmp/emdmodel/10001", []string{"serve"}, nil)
	if err != nil {
		fmt.Printf("Error loading saved model: %s\n", err.Error())
		return
	}
	defer model.Session.Close()

	// bytes, err := ioutil.ReadFile("/home/sailingzhang/winshare/develop/source/mygit/tmp/ml_pic/to_compare/to_picface_0.png")
	// if err != nil {
	// 	return
	// }
	// DecodeJpeg uses a scalar String-valued tensor as input.
	// tensor, err := tf.NewTensor(string(bytes))
	// if err != nil {
	// 	return
	// }
	originalinput :=[10*160*160*3]float32{}
	// fakeInput =[10][160][160][3]float32(originalinput)
	// covertinpult :=make([10][160][160][3]float32{})
	// fakeInput, _ := tf.NewTensor([10*160*160*3]float32{})
	r := bytes.NewReader([]byte(originalinput))
	shape :=make([]int64)
	shape =append(shape,10,160,160,3)
	fakeInput,_=tf.ReadTensor(tf.DataType.float32,shape,r)
	// fakeInput, _ := tf.NewTensor([10][160][160][3]float32{})
	
	fakeInput.Shape()
	
	istrainInput, _ := tf.NewTensor(true)

	for {
		t1 := time.Now()
		_, err := model.Session.Run(
			map[tf.Output]*tf.Tensor{
				model.Graph.Operation("input").Output(0):       fakeInput,
				model.Graph.Operation("phase_train").Output(0): istrainInput, // Replace this with your input layer name
			},
			[]tf.Output{
				model.Graph.Operation("embeddings").Output(0), // Replace this with your output layer name
			},
			nil,
		)

		if err != nil {
			fmt.Printf("Error running the session with input, err: %s\n", err.Error())
			return
		}
		t2 := time.Now()
		fmt.Printf("tf cost2=%v\n", t2.Sub(t1))
		// fmt.Printf("Result value: %v \n", result[0].Value())
	}

}

func my_gf() {
	model := tg.LoadModel("/tmp/emdmodel/10001", []string{"serve"}, nil)

	// bytes, err := ioutil.ReadFile("/home/sailingzhang/winshare/develop/source/mygit/tmp/ml_pic/to_compare/to_picface_0.png")
	// if err != nil {
	// 	return
	// }
	// fakeInput, _ := tf.NewTensor(string(bytes))

	fakeInput, _ := tf.NewTensor([10][160][160][3]float32{})
	istrainInput, _ := tf.NewTensor(true)
	// results := model.Exec([]tf.Output{
	// 	model.Op("embeddings", 0),
	// }, map[tf.Output]*tf.Tensor{
	// 	model.Op("input", 0): fakeInput,
	// })

	for {
		t1 := time.Now()
		results := model.Exec([]tf.Output{
			model.Op("embeddings", 0),
		}, map[tf.Output]*tf.Tensor{
			model.Op("input", 0): fakeInput, model.Op("phase_train", 0): istrainInput,
		})
		_ = results[0].Value().([][]float32)
		//predictions := results[0].Value().([][]float32)
		//fmt.Println(predictions)
		t2 := time.Now()
		fmt.Printf("cost=%v\n", t2.Sub(t1))
	}

	// results := model.Exec([]tf.Output{
	// 	model.Op("embeddings", 0),
	// }, map[tf.Output]*tf.Tensor{
	// 	model.Op("input", 0): fakeInput,model.Op("phase_train",0):istrainInput,
	// })

	// predictions := results[0].Value().([][]float32)
	// fmt.Println(predictions)
}

// func makeTensorFromImage(filename string) (*tf.Tensor, error) {
// 	bytes, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// DecodeJpeg uses a scalar String-valued tensor as input.
// 	tensor, err := tf.NewTensor(string(bytes))
// 	if err != nil {
// 		return nil, err
// 	}
// 	// Construct a graph to normalize the image
// 	graph, input, output, err := constructGraphToNormalizeImage()
// 	if err != nil {
// 		return nil, err
// 	}
// 	// Execute that graph to normalize this one image
// 	session, err := tf.NewSession(graph, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer session.Close()
// 	normalized, err := session.Run(
// 		map[tf.Output]*tf.Tensor{input: tensor},
// 		[]tf.Output{output},
// 		nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return normalized[0], nil
// }

func my_batchtf() {
	// replace myModel and myTag with the appropriate exported names in the chestrays-keras-binary-classification.ipynb
	model, err := tf.LoadSavedModel("/tmp/emdmodel/10001", []string{"myTag"}, nil)
	if err != nil {
		fmt.Printf("Error loading saved model: %s\n", err.Error())
		return
	}
	defer model.Session.Close()

	// batchShape := []int64{len(images), 224, 224, 3}
	// batch, err := tf.ReadTensor(tf.Float, batchShape, &buf)
	// if err != nil {
	// // Handle error
	// }
	tensor, _ := tf.NewTensor([1][160][160][3]float32{})

	result, err := model.Session.Run(
		map[tf.Output]*tf.Tensor{
			model.Graph.Operation("inputLayer_input").Output(0): tensor, // Replace this with your input layer name
		},
		[]tf.Output{
			model.Graph.Operation("inferenceLayer/Sigmoid").Output(0), // Replace this with your output layer name
		},
		nil,
	)

	if err != nil {
		fmt.Printf("Error running the session with input, err: %s\n", err.Error())
		return
	}

	fmt.Printf("Result value: %v \n", result[0].Value())

}

func test1() {
	// Construct a graph with an operation that produces a string constant.
	s := op.NewScope()
	c := op.Const(s, "Hello from TensorFlow version "+tf.Version())
	graph, err := s.Finalize()
	if err != nil {
		panic(err)
	}

	// Execute the graph in a session.
	sess, err := tf.NewSession(graph, nil)
	if err != nil {
		panic(err)
	}
	output, err := sess.Run(nil, []tf.Output{c}, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(output[0].Value())
}

func main() {
	// return 
	// my_tf()
	my_tf2()
	// my_gf()
}
