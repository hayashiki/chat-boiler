package app

import (
	config2 "github.com/hayashiki/chat-boiler/server/config"
	graph2 "github.com/hayashiki/chat-boiler/server/graph"
	ds2 "github.com/hayashiki/chat-boiler/server/infra/ds"
	message2 "github.com/hayashiki/chat-boiler/server/infra/ds/message"
	room2 "github.com/hayashiki/chat-boiler/server/infra/ds/room"
	redis2 "github.com/hayashiki/chat-boiler/server/infra/redis"
	"net/http"
)

type Dependency struct {
	roomRepo       room2.Repository
	messageRepo    message2.Repository
	transactor     ds2.Transactor
	resolver       *graph2.Resolver
	graphQLHandler http.Handler
	//redis redisClient
}

func (d *Dependency) Inject(conf config2.Config) {
	// infrastructure
	//dsCli, err := ds.NewClient(context.Background(), config.GetProject())
	//if err != nil {
	//	log.Fatalf("failed to read datastore client")
	//}
	//dsCliFunc := ds.NewDSFactory(config.GetProject())

	//tran := ds.NewTransaction(dsCliFunc)

	//redisUrl := os.Getenv("REDIS_URL")
	redisUrl := "redis://localhost:6379"
	redisClient, err := redis2.NewRedisClient(redisUrl)
	if err != nil {
		panic(err)
	}

	d.transactor = ds2.NewDatastoreTransactor()
	// repository
	d.roomRepo = room2.NewRepository()
	d.messageRepo = message2.NewRepository()

	resolver := graph2.NewResolver(d.transactor, d.messageRepo, d.roomRepo, redisClient)
	d.resolver = resolver
	d.graphQLHandler = NewGraphQLHandler(d.resolver)
}
