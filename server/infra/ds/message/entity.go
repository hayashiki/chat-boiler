package message

import (
	"github.com/google/uuid"
	model2 "github.com/hayashiki/chat-boiler/server/graph/model"
)

const kind = "Message"

type Entity struct {
	Kind   string `boom:"kind,Message"`
	ID     string `datastore:"-" boom:"id" json:"id" `
	RoomID string `json:"room_id"`
	//ParentKey datastore.Key `datastore:"-" boom:"parent"`
	Text string `json:"text"`
}

func NewEntity(roomID string, text string) *Entity {
	uuid := uuid.New()
	return &Entity{
		ID:     uuid.String(),
		RoomID: roomID,
		Text:   text,
	}
}

func (e *Entity) ToModel() *model2.Message {
	return &model2.Message{
		ID:     e.ID,
		RoomID: e.RoomID,
		Text:   e.Text,
	}
}
