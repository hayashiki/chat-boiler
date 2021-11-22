package app

import (
	"github.com/hayashiki/chat-boiler/server/src/config"
	"github.com/hayashiki/chat-boiler/server/src/graph"
	"github.com/hayashiki/chat-boiler/server/src/infra/ds"
	"github.com/hayashiki/chat-boiler/server/src/infra/ds/message"
	"github.com/hayashiki/chat-boiler/server/src/infra/ds/room"
	"github.com/hayashiki/chat-boiler/server/src/infra/redis"
	"net/http"
	"os"
)

type Dependency struct {
	roomRepo    room.Repository
	messageRepo      message.Repository
	transactor ds.Transactor
	resolver       *graph.Resolver
	graphQLHandler http.Handler
	//redis redisClient
}

func (d *Dependency) Inject(conf config.Config) {
	// infrastructure
	//dsCli, err := ds.NewClient(context.Background(), config.GetProject())
	//if err != nil {
	//	log.Fatalf("failed to read datastore client")
	//}
	//dsCliFunc := ds.NewDSFactory(config.GetProject())

	//tran := ds.NewTransaction(dsCliFunc)

	redisUrl := os.Getenv("REDIS_URL")
	redisClient, err := redis.NewRedisClient(redisUrl)
	if err != nil {
		panic(err)
	}

	d.transactor = ds.NewDatastoreTransactor()
	// repository
	d.roomRepo = room.NewRepository()
	d.messageRepo = message.NewRepository()

	resolver := graph.NewResolver(d.transactor, d.messageRepo, d.roomRepo, redisClient)
	d.resolver = resolver
	d.graphQLHandler = NewGraphQLHandler(d.resolver)
}
