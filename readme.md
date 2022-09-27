# POC to implement GRPC 


## Working

1. Every request sends "Limit" which means the upper bound for a random number e.g. [0,upper_bound)
2. Some stream related request may also contain a param name "Count" which means number of response message required for a given "Limit"

### Methods

1. SayHello
    Client sends a "Limit" and sever returns a random number i.e. [0,Limit)
2. SayHelloWithClientStream
    Client sends an arbitrary number of requests to generate random number upto the max of "Limit"s specified by all requests and finally responds with [0,max(all_Limits)) 
3. SayHelloWithServerStream
    Client sends number of response wanted and an upper "Limit" and server responds with "Count" number of response with a random number i.e. [0,Limit)    
4. SayHelloWithClientServerStream
    Client sends an arbitrary number of requests to generate random number upto the "Limit" specified by the that request and server responds immedietly for that request untils the incoming stream is closed

## How to run
1. Build with `go build .`
2. Invoke the generated binary