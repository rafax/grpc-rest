package server

import (
	"testing"
	"time"

	"github.com/rafax/grpc-rest/pb"
	ctx "golang.org/x/net/context"
)

func TestExpiredTokenIsNotAcceptedAndRemoved(t *testing.T) {
	s := NewAuthServer()
	token := "test"
	s.tokens[token] = time.Now().Add(time.Duration(-1) * time.Hour)
	r, err := s.Validate(ctx.Background(), &pb.CredentialsRequest{Token: token})
	if err == nil {
		t.Error("Expected an error, got ", err)
	}
	if r != nil {
		t.Error("Expected nil response, got ", r)
	}
	if len(s.tokens) != 0 {
		t.Error("tokens should be empty, got ", s.tokens)
	}
}

func TestInvalidTokenIsNotAccepted(t *testing.T) {
	s := NewAuthServer()
	token := "invalid"
	r, err := s.Validate(ctx.Background(), &pb.CredentialsRequest{Token: token})
	if err == nil {
		t.Error("Expected an error, got ", err)
	}
	if r != nil {
		t.Error("Expected nil response, got ", r)
	}
	if len(s.tokens) != 0 {
		t.Error("tokens should be empty, got ", s.tokens)
	}
}

func TestValidTokenIsAccepted(t *testing.T) {
	s := NewAuthServer()
	token := "test"
	s.tokens[token] = time.Now().Add(time.Duration(1) * time.Hour)
	r, err := s.Validate(ctx.Background(), &pb.CredentialsRequest{Token: token})
	if err != nil {
		t.Error("Expected nil error, got ", err)
	}
	if r.Token != token {
		t.Error("Expected passed token in response, got ", r)
	}
}

func TestLoginPersistsAToken(t *testing.T) {
	s := NewAuthServer()
	auth := &pb.LogInRequest{Username: "user", Password: "pass"}
	r, err := s.LogIn(ctx.Background(), auth)
	if err != nil {
		t.Error("Expected nil error, got ", err)
	}
	if r == nil {
		t.Error("Expected token in response, got ", r)
	}
	valid, ok := s.tokens[r.Token]
	if !ok {
		t.Error("Expected token to be stored in server.tokens")
	}
	if r.ValidUntil != valid.String() {
		t.Error("Bad validity in response in response, got ", r.ValidUntil, " expected ", valid)
	}
}
