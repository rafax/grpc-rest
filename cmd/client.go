package cmd

import (
	"log"

	pb "github.com/rafax/grpc-rest/pb"

	c "github.com/rafax/grpc-rest/client"
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
		client := c.GrpcClient{}
		client.Init(address)
		name := defaultName
		r, err := client.LogIn(&pb.LogInRequest{Username: name, Password: "secret"})
		if err != nil {
			log.Fatalf("could not log in: %v", err)
		}
		log.Printf("Logged in: %s", r)
	},
}

func init() {
	RootCmd.AddCommand(clientCmd)
}
