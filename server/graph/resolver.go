package graph

import (
	"encoding/json"
	"github.com/go-redis/redis"
	model2 "github.com/hayashiki/chat-boiler/server/graph/model"
	ds2 "github.com/hayashiki/chat-boiler/server/infra/ds"
	message2 "github.com/hayashiki/chat-boiler/server/infra/ds/message"
	room2 "github.com/hayashiki/chat-boiler/server/infra/ds/room"
	"log"
	"sync"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	//pubsub *pubsub.Client
	transaction     ds2.Transactor
	messageRepo     message2.Repository
	roomRepo        room2.Repository
	redis           *redis.Client
	mutex           sync.Mutex
	messageChannels map[string]chan *model2.Message
}

func NewResolver(
	transactor ds2.Transactor,
	messageRepo message2.Repository,
	roomRepo room2.Repository,
	redis *redis.Client,
	) *Resolver {
	res := &Resolver{
		transaction: transactor,
		messageRepo: messageRepo,
		roomRepo:    roomRepo,
		redis: redis,
		mutex: sync.Mutex{},
		messageChannels: map[string]chan *model2.Message{},
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
				m := model2.Message{}
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
