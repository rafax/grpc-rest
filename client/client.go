package client

import (
	"log"

	ctx "golang.org/x/net/context"

	pb "github.com/rafax/grpc-rest/pb"
	"google.golang.org/grpc"
)

// AuthClient abstracts grpc and rest clients
type AuthClient interface {
	// Init initializes client with given endpoint
	Init(endpoint string)
	// LogIn logs in
	LogIn(r *pb.LogInRequest) (*pb.AuthResponse, error)
	// Validate validates
	Validate(r *pb.CredentialsRequest) (*pb.AuthResponse, error)
}

// GrpcClient is a GRPC implementation of AuthClient
type GrpcClient struct {
	client pb.AuthenticatorClient
}

// Init initializes client with given endpoint
func (g *GrpcClient) Init(endpoint string) {
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	g.client = pb.NewAuthenticatorClient(conn)
}

// LogIn logs in
func (g *GrpcClient) LogIn(r *pb.LogInRequest) (*pb.AuthResponse, error) {
	return g.client.LogIn(ctx.Background(), r)
}

// Validate validates
func (g *GrpcClient) Validate(r *pb.CredentialsRequest) (*pb.AuthResponse, error) {
	return g.client.Validate(ctx.Background(), r)
}
