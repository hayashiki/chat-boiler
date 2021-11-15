package graph

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/hayashiki/gql-chat/server/src/graph/model"
	"github.com/hayashiki/gql-chat/server/src/infra/ds"
	"github.com/hayashiki/gql-chat/server/src/infra/ds/message"
	"github.com/hayashiki/gql-chat/server/src/infra/ds/room"
	"log"
	"sync"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	//pubsub *pubsub.Client
	transaction  ds.Transactor
	messageRepo message.Repository
	roomRepo room.Repository
	redis *redis.Client
	mutex sync.Mutex
	messageChannels map[string]chan *model.Message
}

func NewResolver(
	transactor ds.Transactor,
	messageRepo message.Repository,
	roomRepo room.Repository,
	redis *redis.Client,
	) *Resolver {
	res := &Resolver{
		transaction: transactor,
		messageRepo: messageRepo,
		roomRepo:    roomRepo,
		redis: redis,
		mutex: sync.Mutex{},
		messageChannels: map[string]chan *model.Message{},
	}
	res.startSubscribingRedis()
	return res
}

func (r *Resolver) startSubscribingRedis() {
	log.Println("start sub...")

	go func() {
		pubsub := r.redis.Subscribe("room")
		defer pubsub.Close()

		for {
			msgi, err := pubsub.Receive()
			if err != nil {
				panic(err)
			}

			switch msg := msgi.(type) {
			case *redis.Message:
				m := model.Message{}
				if err := json.Unmarshal([]byte(msg.Payload), &m); err != nil {
					log.Println(err)
					continue
				}
				r.mutex.Lock()
				for _, ch := range r.messageChannels {
					ch <- &m
				}
				r.mutex.Unlock()
			default:
			}
		}
	}()
}
