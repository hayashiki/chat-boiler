package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/hayashiki/gql-chat/server/src/app"
	"github.com/hayashiki/gql-chat/server/src/config"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const defaultPort = "8080"
const shutdownTimeout = 30 * time.Second

func main()  {
	//proj := os.Getenv("GCP_PROJECT")
	//
	//client, err := pubsub.NewClient(ctx, proj)
	//if err != nil {
	//	panic(err)
	//}
	//defer client.Close()

	rootCmd := &cobra.Command{
		Use:   "web",
		Short: "web server",
		Long:  "web server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Run()
		},
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := Run(); err != nil {
		os.Exit(1)
	}
}

func Run() error {
	port := os.Getenv("PORT")


	if port == "" {
		port = defaultPort
	}

	conf, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed to read config")
	}

	d := &app.Dependency{}
	d.Inject(conf)
	r := chi.NewRouter()
	app.Routing(r, d)


	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	log.Printf("Listening on port %s", port)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, os.Interrupt)
	<-sigCh

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("graceful shutdown failure: %s", err)
	}
	log.Printf("graceful shutdown successfully")
	return nil
}
