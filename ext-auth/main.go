package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"

	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	auth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
)

type authorizationServer struct {
	auth.UnimplementedAuthorizationServer
}

// Check implements authorization logic.
func (a *authorizationServer) Check(ctx context.Context, req *auth.CheckRequest) (*auth.CheckResponse, error) {
	log.Printf("Received auth request:")
	log.Printf("  Method: %s", req.Attributes.Request.Http.Method)
	log.Printf("  Path: %s", req.Attributes.Request.Http.Path)
	log.Printf("  Host: %s", req.Attributes.Request.Http.Host)

	// Print headers
	if req.Attributes.Request.Http.Headers != nil {
		log.Printf("  Headers:")
		for k, v := range req.Attributes.Request.Http.Headers {
			log.Printf("    %s: %s", k, v)
		}
	}

	// Always allow - return OK response
	return &auth.CheckResponse{
		Status: &status.Status{
			Code: int32(code.Code_OK),
		},
		HttpResponse: &auth.CheckResponse_OkResponse{
			OkResponse: &auth.OkHttpResponse{
				Headers: []*core.HeaderValueOption{
					{
						Header: &core.HeaderValue{
							Key:   "x-auth-result",
							Value: "allowed",
						},
					},
				},
			},
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	auth.RegisterAuthorizationServer(s, &authorizationServer{})

	log.Println("Starting gRPC external auth server on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
