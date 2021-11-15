package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hayashiki/gql-chat/server/src/graph/generated"
	"github.com/hayashiki/gql-chat/server/src/graph/model"
	"github.com/hayashiki/gql-chat/server/src/infra/ds/message"
	"github.com/hayashiki/gql-chat/server/src/infra/ds/room"
	"go.mercari.io/datastore/boom"
)

func (r *mutationResolver) CreateRoom(ctx context.Context, name string, description *string) (*model.Room, error) {
	log.Println("called")
	newRoom := &room.Entity{}
	if err := r.transaction.RunInTransaction(ctx, func(tx *boom.Transaction) error {
		newRoom = room.NewEntity(name, *description)
		log.Println("called", newRoom)
		if err := r.roomRepo.Put(tx, newRoom); err != nil {
			log.Println("called err", err)
			return err
		}
		return nil
	}); err != nil {
		log.Println("err", err)
		return nil, err
	}
	return newRoom.ToModel(), nil
}

func (r *mutationResolver) PostMessage(ctx context.Context, roomID string, text string) (*model.Message, error) {
	log.Println("called")
	newMessage := &message.Entity{}
	if err := r.transaction.RunInTransaction(ctx, func(tx *boom.Transaction) error {
		newMessage = message.NewEntity(roomID, text)
		log.Println("called", newMessage)
		if err := r.messageRepo.Put(tx, newMessage); err != nil {
			log.Println("called err", err)
			return err
		}
		return nil
	}); err != nil {
		log.Println("err", err)
		return nil, err
	}
	mb, err := json.Marshal(newMessage)
	if err != nil {
		return nil, err
	}
	r.redis.Publish("room", mb)
	return newMessage.ToModel(), nil
}

func (r *queryResolver) Rooms(ctx context.Context) ([]*model.Room, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Messages(ctx context.Context, roomID string) ([]*model.Message, error) {
	entities, err := r.messageRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	messages := make([]*model.Message, len(entities))

	for i, e := range entities {
		messages[i] = e.ToModel()
	}

	return messages, nil
}

func (r *subscriptionResolver) MessagePosted(ctx context.Context, roomID string) (<-chan *model.Message, error) {
	log.Println("MessagePosted")
	messageChan := make(chan *model.Message, 1)
	log.Println("MessagePosted 0", messageChan)
	r.mutex.Lock()
	r.messageChannels[roomID] = messageChan
	r.mutex.Unlock()

	log.Println("MessagePosted 1", messageChan)
	go func() {
		<-ctx.Done()
		log.Println("MessagePosted 2", messageChan)
		r.mutex.Lock()
		r.mutex.Unlock()
	}()

	log.Println("MessagePosted 3", messageChan)
	return messageChan, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
