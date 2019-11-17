package main

import (
    "github.com/cihub/seelog"
    // "context"
    "log"
    "net"
    "google.golang.org/grpc"
    pb "simple-tensorflow-serving/predict_proto/proto/tensorflow_s"
    "simple-tensorflow-serving/tfService"
)



func main() {
    seelog.Tracef("htt\n")
    seelog.Debugf("heeloo\n")
    lis, err := net.Listen("tcp", ":5555")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterPredictionServiceServer(s, &tfService.TfService1{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}