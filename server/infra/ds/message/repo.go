package message

import (
	"context"
	ds2 "github.com/hayashiki/chat-boiler/server/infra/ds"
	"github.com/pkg/errors"
	"go.mercari.io/datastore/boom"
)

type Repository interface {
	Get(ctx context.Context, id string) (*Entity, error)
	GetAll(ctx context.Context) ([]*Entity, error)
	Put(tx *boom.Transaction, item *Entity) error
}

//dsCli *datastore.Client
func NewRepository() Repository {
	return &repository{}
}

type repository struct {}

func (r *repository) Get(ctx context.Context, id string) (*Entity, error) {
	panic("implement me")
}

func (r *repository) GetAll(ctx context.Context) ([]*Entity, error) {
	b := ds2.FromContext(ctx)
	q := b.Client.NewQuery(kind)

	var entities []*Entity
	if _, err := b.GetAll(q, &entities); err != nil {
		return nil, errors.WithStack(err)
	}
	return entities, nil
}

func (r *repository) Put(tx *boom.Transaction, item *Entity) error {
	if _, err := tx.Put(item); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
