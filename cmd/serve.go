package cmd

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/rafax/grpc-rest/pb"
	server "github.com/rafax/grpc-rest/server"

	"github.com/spf13/cobra"
)

const port = ":65432"

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
		pb.RegisterAuthenticatorServer(s, server.NewAuthServer())
		fmt.Println("Running on ", port)
		s.Serve(lis)
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}
