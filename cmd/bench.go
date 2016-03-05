package cmd

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/rafax/grpc-rest/client"
	pb "github.com/rafax/grpc-rest/pb"

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
			grpcEndpoint = "localhost:65432"
			restEndpoint = "http://localhost:8080"
			defaultName  = "world"
			requests     = int64(1000)
			workers      = 1
			name         = defaultName
		)

		jobs := make(chan pb.LogInRequest)
		c := client.NewGrpcClient(grpcEndpoint)
		//c := client.NewRestClient(restEndpoint)
		var wg sync.WaitGroup

		for w := 0; w <= workers; w++ {
			go func() {
				for lr := range jobs {
					defer wg.Done()
					r, err := c.LogIn(&lr)
					if err != nil {
						log.Fatalf("could not log in: %v", err)
						return
					}
					_, err = c.Validate(&pb.CredentialsRequest{Token: r.Token})
					if err != nil {
						log.Fatalf("could not validate token: %v", err)
					}
				}
			}()
		}

		start := time.Now()

		for i := int64(0); i < requests; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				r, err := c.LogIn(&pb.LogInRequest{Username: name, Password: "secret"})
				if err != nil {
					log.Fatalf("could not log in: %v", err)
					return
				}
				_, err = c.Validate(&pb.CredentialsRequest{Token: r.Token})
				if err != nil {
					log.Fatalf("could not validate token: %v", err)
				}
			}()
		}
		close(jobs)
		wg.Wait()
		elapsed := time.Since(start)
		fmt.Printf("Executed %d requests in %s, %f ms/op", requests*2, elapsed, float64(elapsed.Nanoseconds())/float64(2*requests*1000000))
	},
}

func init() {
	RootCmd.AddCommand(benchCmd)
}
