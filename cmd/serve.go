package cmd

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"

	"golang.org/x/net/context"

	"google.golang.org/grpc"

	pb "github.com/rafax/grpc-rest/pb"
	"github.com/spf13/cobra"
)

const port = ":65432"

type server struct{}

func (s *server) LogIn(c context.Context, lir *pb.LogInRequest) (*pb.AuthResponse, error) {
	//fmt.Println("Logged in ", lir)
	return &pb.AuthResponse{Token: strconv.Itoa(rand.Int()), ValidUntil: time.Now().String()}, nil
}

func (s *server) Validate(ctx context.Context, in *pb.CredentialsRequest) (*pb.AuthResponse, error) {
	return &pb.AuthResponse{Token: strconv.Itoa(rand.Int()), ValidUntil: time.Now().String()}, nil
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterAuthenticatorServer(s, &server{})
		fmt.Println("Running on ", port)
		s.Serve(lis)
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}
