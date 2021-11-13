package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main()  {
	ctx := context.Background()
	proj := os.Getenv("GCP_PROJECT")

	client, err := pubsub.NewClient(ctx, proj)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	rootCmd := &cobra.Command{
		Use:   "sub",
		Short: "subscriber",
		Long:  "subscriber",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := server.Run()
			return err
		},
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
