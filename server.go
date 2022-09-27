package main

import (
	"context"
	"io"
	"log"
	"math/rand"
	"net"
	"strconv"

	"github.com/devansh42/grpc-poc/poc"
	"google.golang.org/grpc"
)

type server struct {
	poc.UnimplementedServiceServer
}

func (server) SayHello(c context.Context, req *poc.SingleRequest) (*poc.Response, error) {
	l := req.Limit
	log.Printf("Server: Received Limit %v", l)
	return &poc.Response{Number: getRandomNumber(l)}, nil
}

func (server) SayHelloWithClientStream(s poc.Service_SayHelloWithClientStreamServer) error {
	var max int32
	for {
		msg, err := s.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("Couldn't read messages from client stream")
			return err
		}
		log.Printf("Server: Received Limit %v", msg.Limit)

		if msg.Limit > max {
			max = msg.Limit
		}
	}
	log.Printf("Server: Received Max Limit %v", max)

	err := s.SendAndClose(&poc.Response{Number: getRandomNumber(max)})
	if err != nil {
		log.Printf("Couldn't send and close the connection due to %v", err)
	}
	return err
}

func (server) SayHelloWithServerStream(req *poc.MultipleRequest, s poc.Service_SayHelloWithServerStreamServer) error {
	var i int32
	var err error
	log.Printf("Server: Received Count %v", req.Count)
	log.Printf("Server: Received Limit %v", req.Limit)

	for i = 0; i < req.Count; i++ {
		err = s.Send(&poc.Response{Number: getRandomNumber(req.Limit)})
		if err != nil {
			log.Printf("Couldn't send msg to client due to %v", err)
			break
		}
	}
	return err
}

func (server) SayHelloWithClientServerStream(s poc.Service_SayHelloWithClientServerStreamServer) error {

	for {
		msg, err := s.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("Couldn't read messages from client stream")
			return err
		}
		log.Printf("Server: Received Limit %v", msg.Limit)

		err = s.Send(&poc.Response{Number: getRandomNumber(msg.Limit)})
		if err != nil {
			log.Printf("Couldn't send msg to client due to %v", err)
			return err
		}
	}
	return nil
}

func startServer(port int) {
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("Couldn't start tcp server due to %v", err)
	}
	serv := grpc.NewServer()
	var s server
	poc.RegisterServiceServer(serv, s)
	log.Println("Starting RPC Server...")
	log.Fatal(serv.Serve(listener))
}

func getRandomNumber(limit int32) int32 {
	return int32(rand.Intn(int(limit)))
}
