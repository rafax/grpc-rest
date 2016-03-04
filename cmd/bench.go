package cmd

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	pb "github.com/rafax/grpc-rest/pb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/spf13/cobra"
)

// clientCmd represents the client command
var benchCmd = &cobra.Command{
	Use:   "bench",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		const (
			address     = "localhost:65432"
			defaultName = "world"
		)

		requests := int64(10000)
		// TODO: Work your own magic here
		// Set up a connection to the server.
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewAuthenticatorClient(conn)
		name := defaultName
		if len(os.Args) > 1 {
			name = os.Args[1]
		}
		var wg sync.WaitGroup
		start := time.Now()
		for i := int64(0); i < requests; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				r, err := c.LogIn(context.Background(), &pb.LogInRequest{Username: name, Password: "secret"})
				if err != nil {
					log.Fatalf("could not log in: %v", err)
					return
				}
				_, err = c.Validate(context.Background(), &pb.CredentialsRequest{Token: r.Token})
				if err != nil {
					log.Fatalf("could not validate token: %v", err)
				}
			}()
		}
		wg.Wait()
		elapsed := time.Since(start)
		fmt.Printf("Executed %d requests in %s, %f ms/op", requests*2, elapsed, float64(elapsed.Nanoseconds())/float64(2*requests*1000000))
	},
}

func init() {
	RootCmd.AddCommand(benchCmd)
}
