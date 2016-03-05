package cmd

import (
	"fmt"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/gengo/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	"github.com/rafax/grpc-rest/pb"
)

// gatewayCmd represents the gateway command
var gatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		endpoint := "localhost:65432"
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{grpc.WithInsecure()}
		err := pb.RegisterAuthenticatorHandlerFromEndpoint(ctx, mux, endpoint, opts)
		if err != nil {
			fmt.Println("Error when registering gateway", err)
		}

		http.ListenAndServe(":8080", mux)
		fmt.Println("gateway called")
	},
}

func init() {
	RootCmd.AddCommand(gatewayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gatewayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gatewayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
