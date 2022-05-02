package main

import (
	"context"
	"net/http"

	pb "github.com/grpc-ecosystem/grpc-gateway/v2/examples/internal/helloworld"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func main() {
	ctx := context.TODO()
	mux := runtime.NewServeMux()
	// Register generated routes to mux
	err := pb.RegisterGreeterHandlerServer(ctx, mux, &GreeterServer{})
	if err != nil {
		panic(err)
	}
	// Register custom route for  GET /hello/{name}
	err = mux.HandlePath("GET", "/hello/{name}", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		w.Write([]byte("hello " + pathParams["name"]))
	})
	if err != nil {
		panic(err)
	}
	http.ListenAndServe(":8080", mux)
}

// GreeterServer is the server API for Greeter service.
type GreeterServer struct {
}

// SayHello implement to say hello
func (h *GreeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: "hello " + req.Name,
	}, nil
}

func (s *serviceServer) Check(ctx context.Context, in *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{Status: health.HealthCheckResponse_SERVING}, nil
}
