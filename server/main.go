package main

import (
	"github.com/cihub/seelog"
	// "context"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "simple-tensorflow-serving/predict_proto/proto/tensorflow_s"
	"simple-tensorflow-serving/tfService"
)

func LOG_INIT(logConfigFile string) {
	logger, err := seelog.LoggerFromConfigAsFile(logConfigFile)
	if nil != err {
		seelog.Critical("err parsing config log file", err)
	}
	seelog.ReplaceLogger(logger)
}

func main() {
	LOG_INIT("linux_log_conf.xml")
	seelog.Tracef("enter")
	defer seelog.Flush()
	lis, err := net.Listen("tcp", ":2233")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPredictionServiceServer(s, tfService.NewTfService1())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
