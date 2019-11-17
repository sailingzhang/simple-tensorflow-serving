module simple-tensorflow-serving

go 1.13

// replace (
// 	simple-tensorflow-serving/tfService => ./simple-tensorflow-serving/tfService
// 	“simple-tensorflow-serving/predict_proto/proto/tensorflow_s” => ./simple-tensorflow-serving/predict_proto/proto/tensorflow_s
// )

require (
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575
	github.com/golang/protobuf v1.3.2
	google.golang.org/grpc v1.25.1
)
