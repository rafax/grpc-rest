package cmd

import (
	"log"

	pb "github.com/rafax/grpc-rest/pb"

	c "github.com/rafax/grpc-rest/client"
	"github.com/spf13/cobra"
)

const (
	grpcEndpoint = "localhost:65432"
	restEndpoint = "http://localhost:8080"
	defaultName  = "world"
)

// clientCmd represents the client command
var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client := c.NewGrpcClient(grpcEndpoint)
		name := defaultName
		r, err := client.LogIn(&pb.LogInRequest{Username: name, Password: "secret"})
		if err != nil {
			log.Fatalf("could not log in: %v", err)
		}
		log.Printf("Logged in: %s", r)
		v, err := client.Validate(&pb.CredentialsRequest{Token: r.Token})
		if err != nil {
			log.Fatalf("could not validate: %v", err)
		}
		log.Printf("Validated: %s", v)

	},
}

func init() {
	RootCmd.AddCommand(echoCmd)
}
