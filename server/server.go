package server

import (
	"errors"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/rafax/grpc-rest/pb"
	"golang.org/x/net/context"
)

// AuthServer stores tokens issued for users
type AuthServer struct {
	tokensMutex *sync.RWMutex
	tokens      map[string]time.Time
}

// NewAuthServer creates a new AuthServer
func NewAuthServer() AuthServer {
	var tm sync.RWMutex
	return AuthServer{tokens: make(map[string]time.Time), tokensMutex: &tm}
}

// LogIn logs issues a token valid for 24 hours
func (s AuthServer) LogIn(c context.Context, lir *pb.LogInRequest) (*pb.AuthResponse, error) {
	t := strconv.Itoa(rand.Int())
	validUntil := time.Now().Add(time.Duration(24) * time.Hour)
	s.tokensMutex.Lock()
	s.tokens[t] = validUntil
	s.tokensMutex.Unlock()
	return &pb.AuthResponse{Token: t, ValidUntil: validUntil.String()}, nil
}

// Validate checks if token is valid and was issued by this server
func (s AuthServer) Validate(ctx context.Context, in *pb.CredentialsRequest) (*pb.AuthResponse, error) {
	s.tokensMutex.RLock()
	validUntil, ok := s.tokens[in.Token]
	s.tokensMutex.RUnlock()
	if ok {
		if time.Now().After(validUntil) {
			s.tokensMutex.Lock()
			delete(s.tokens, in.Token)
			s.tokensMutex.Unlock()
			return nil, errors.New("Token expired")
		}
		return &pb.AuthResponse{Token: in.Token, ValidUntil: validUntil.String()}, nil
	}
	return nil, errors.New("Invalid token")
}
