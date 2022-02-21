package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	generated2 "github.com/hayashiki/chat-boiler/server/graph/generated"
	model2 "github.com/hayashiki/chat-boiler/server/graph/model"
	message2 "github.com/hayashiki/chat-boiler/server/infra/ds/message"
	room2 "github.com/hayashiki/chat-boiler/server/infra/ds/room"
	"log"

	"go.mercari.io/datastore/boom"
)

func (r *mutationResolver) CreateRoom(ctx context.Context, input model2.RoomInput) (*model2.Room, error) {
	newRoom := &room2.Entity{}
	if err := r.transaction.RunInTransaction(ctx, func(tx *boom.Transaction) error {
		newRoom = room2.NewEntity(input.Name, input.Description)
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

func (r *mutationResolver) CreateMessage(ctx context.Context, input model2.CreateMessageInput) (*model2.Message, error) {
	log.Println("called")
	newMessage := &message2.Entity{}
	if err := r.transaction.RunInTransaction(ctx, func(tx *boom.Transaction) error {
		newMessage = message2.NewEntity(input.RoomID, input.Text)
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

func (r *queryResolver) Rooms(ctx context.Context) ([]*model2.Room, error) {
	rooms, err := r.roomRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	mRooms := make([]*model2.Room, len(rooms))
	for i, r := range rooms {
		mRooms[i] = r.ToModel()
	}
	return mRooms, err
}

func (r *queryResolver) Messages(ctx context.Context, roomID string) ([]*model2.Message, error) {
	entities, err := r.messageRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	messages := make([]*model2.Message, len(entities))

	for i, e := range entities {
		messages[i] = e.ToModel()
	}

	return messages, nil
}

func (r *subscriptionResolver) MessagePosted(ctx context.Context, roomID string) (<-chan *model2.Message, error) {
	log.Println("MessagePosted")
	messageChan := make(chan *model2.Message, 1)
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
func (r *Resolver) Mutation() generated2.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated2.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated2.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
