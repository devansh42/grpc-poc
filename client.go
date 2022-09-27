package main

import (
	"context"
	"io"
	"log"
	"math/rand"
	"strconv"
	"sync"

	"github.com/devansh42/grpc-poc/poc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type client struct{}

func startClient(port int) {
	conn, err := grpc.Dial("localhost:"+strconv.Itoa(port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Couldn't dial to server due to %v", err)
	}
	defer conn.Close()
	cli := poc.NewServiceClient(conn)

	ctx := context.Background()
	log.Print("Calling SayHello")
	resp, err := cli.SayHello(ctx, &poc.SingleRequest{
		Limit: 20,
	})
	if err != nil {
		log.Fatalf("Couldn't make call to sayHello due to %v", err)
	}
	log.Printf("Response: SayHello: %v", resp.Number)

	log.Print("Calling SayHelloWithClientStream")
	cstream, err := cli.SayHelloWithClientStream(ctx)
	if err != nil {
		log.Fatalf("Couldn't make call to sayHello due to %v", err)
	}

	for i := 0; i < 10; i++ {
		err = cstream.Send(&poc.SingleRequest{Limit: getRandomLimit()})
		if err != nil {
			log.Printf("Couldn't send stream message due to %v", err)
			return
		}
	}
	resp, err = cstream.CloseAndRecv()
	if err != nil {
		log.Printf("Couldn't close the stream due to %v", err)
		return
	}
	log.Printf("Response: SayHelloWithClientStream: %v", resp.Number)

	log.Print("Calling SayHelloWithServerStream")

	sstream, err := cli.SayHelloWithServerStream(ctx, &poc.MultipleRequest{
		Limit: getRandomLimit(),
		Count: 10,
	})
	if err != nil {
		log.Fatalf("Couldn't make call to SayHelloWithServerStream due to %v", err)
	}
	for {
		resp, err = sstream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Printf("Done with Stream")
				break
			}
			log.Printf("Couldn't recv from stream due to %v", err)
		}
		log.Printf("Response: Received Number %v", resp.GetNumber())
	}

	log.Print("Calling SayHelloWithClientServerStream")

	csstream, err := cli.SayHelloWithClientServerStream(ctx)
	if err != nil {
		log.Fatalf("Couldn't make call to SayHelloWithClientServerStream due to %v", err)
	}
	var wg sync.WaitGroup
	wg.Add(1)

	go biStreamProcessor(&wg, csstream)

	for i := 0; i < 10; i++ {
		err = csstream.Send(&poc.SingleRequest{
			Limit: getRandomLimit(),
		})
		if err != nil {
			log.Printf("Couldn't send request to server due to %v", err)
			break
		}

	}
	err = csstream.CloseSend()
	if err != nil {
		log.Printf("Couldn't close the stream properly due to %v", err)
	}
	wg.Wait()
}

func biStreamProcessor(wg *sync.WaitGroup, csstream poc.Service_SayHelloWithClientServerStreamClient) {
	defer wg.Done()
	for {
		msg, err := csstream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("Couldn't recieve messages due to %v", err)
			return
		}
		log.Printf("Response: Received Number: %v", msg.Number)
	}
}

func getRandomLimit() int32 {
	return int32(rand.Intn(100000))
}
