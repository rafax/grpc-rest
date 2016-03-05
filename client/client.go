package client

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	ctx "golang.org/x/net/context"

	pb "github.com/rafax/grpc-rest/pb"
	"google.golang.org/grpc"
)

// AuthClient abstracts grpc and rest clients
type AuthClient interface {
	// LogIn logs in
	LogIn(r *pb.LogInRequest) (*pb.AuthResponse, error)
	// Validate validates
	Validate(r *pb.CredentialsRequest) (*pb.AuthResponse, error)
}

// GrpcClient is a GRPC implementation of AuthClient
type GrpcClient struct {
	client pb.AuthenticatorClient
}

// NewGrpcClient creates a new grpc client for given endpoint
func NewGrpcClient(endpoint string) AuthClient {
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return GrpcClient{pb.NewAuthenticatorClient(conn)}
}

// LogIn logs in
func (g GrpcClient) LogIn(r *pb.LogInRequest) (*pb.AuthResponse, error) {
	return g.client.LogIn(ctx.Background(), r)
}

// Validate validates
func (g GrpcClient) Validate(r *pb.CredentialsRequest) (*pb.AuthResponse, error) {
	return g.client.Validate(ctx.Background(), r)
}

// RestClient is a REST implementation of AuthClient
type RestClient struct {
	endpoint string
}

// NewRestClient creates a new rest client for given endpoint
func NewRestClient(endpoint string) AuthClient {
	return RestClient{endpoint: endpoint}
}

// LogIn logs in
func (c RestClient) LogIn(r *pb.LogInRequest) (*pb.AuthResponse, error) {
	req, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(c.endpoint+"/v1/login", "application/json", bytes.NewReader(req))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var ar pb.AuthResponse
	json.NewDecoder(resp.Body).Decode(&ar)
	return &ar, nil
}

// Validate validates
func (c RestClient) Validate(r *pb.CredentialsRequest) (*pb.AuthResponse, error) {
	req, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(c.endpoint+"/v1/validate", "application/json", bytes.NewReader(req))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var ar pb.AuthResponse
	err = json.NewDecoder(resp.Body).Decode(&ar)
	if err != nil {
		return nil, err
	}
	return &ar, nil
}
