package room

import (
	"github.com/google/uuid"
	"github.com/hayashiki/gql-chat/server/src/graph/model"
)

type Entity struct {
	Kind       string `boom:"kind,Room"`
	ID          string     `boom:"id" json:"id" `
	Name        string     `json:"name"`
	Description string    `json:"description"`
}

func NewEntity(name string, description string) *Entity {
	uuid := uuid.New()
	return &Entity{
		ID: uuid.String(),
		Name:        name,
		Description: description,
	}
}

func (e *Entity) ToModel() *model.Room {
	return &model.Room{
		ID:     e.ID,
		Name: e.Name,
		Description: &e.Description,
	}
}
