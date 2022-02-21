package app

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"github.com/gorilla/websocket"
	graph2 "github.com/hayashiki/chat-boiler/server/graph"
	generated2 "github.com/hayashiki/chat-boiler/server/graph/generated"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
)

// NewGraphQLHandler returns GraphQL Server.
func NewGraphQLHandler(
	resolver *graph2.Resolver,
) http.Handler {
	srv := handler.New(generated2.NewExecutableSchema(
		generated2.Config{
			Resolvers:  resolver,
			Directives: generated2.DirectiveRoot{},
			Complexity: generated2.ComplexityRoot{},
		},
	))
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			//				origin := r.Header.Get("Origin")
			//				origin = strings.ToLower(origin)
			//				for _, o := range corsOpts.AllowedOrigins {
			//					if o == origin {
			//						return true
			//					}
			//				}
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},

		InitFunc: func(ctx context.Context, initPayload transport.InitPayload) (context.Context, error) {
			token := initPayload["authToken"]
			if token == nil {
				return ctx, errors.New("Unauthorized")
			}
			log.Println("ws token", token)
			//parsedToken, err := auth.ParseAndValidateJWT(token.(string))
			//if err != nil {
			//	return ctx, errors.New("Unauthorized")
			//}
			//authCtx := auth.SetUserContext(ctx, *parsedToken)
			//return authCtx, nil
			return ctx, nil
		},

		KeepAlivePingInterval: 10 * time.Second,
	})


	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})
	srv.SetQueryCache(lru.New(1000))
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})
	srv.SetErrorPresenter(errorPresenter)
	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		return &gqlerror.Error{
			Message: "server error",
			Extensions: map[string]interface{}{
				"type": "Internal",
				"code": "Unknown",
			},
		}
	})

	return srv
}

func errorPresenter(ctx context.Context, e error) *gqlerror.Error {
	err := graphql.DefaultErrorPresenter(ctx, e)
	return err
}

// NewRootHandler returns GraphQL Playground.
func NewRootHandler() http.Handler {
	return playground.Handler("GraphQL playground", "/query")
}
