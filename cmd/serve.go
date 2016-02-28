// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

	pb "github.com/rafax/grpc_rest/pb"
	"github.com/spf13/cobra"
)

const port = ":65432"

type server struct{}

func (s *server) LogIn(context.Context, *pb.LogInRequest) (*pb.AuthResponse, error) {
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
		fmt.Println("RUnning on ", port)
		s.Serve(lis)
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
