package cmd

import (
	"log"
	"os"

	pb "github.com/rafax/grpc-rest/pb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/spf13/cobra"
)

const (
	address     = "localhost:65432"
	defaultName = "world"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewAuthenticatorClient(conn)

		// Contact the server and print out its response.
		name := defaultName
		if len(os.Args) > 1 {
			name = os.Args[1]
		}
		r, err := c.LogIn(context.Background(), &pb.LogInRequest{Username: name, Password: "secret"})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Logged in: %s", r)
	},
}

func init() {
	RootCmd.AddCommand(clientCmd)
}
