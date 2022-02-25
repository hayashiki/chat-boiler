package ds

import (
	"context"
	config2 "github.com/hayashiki/chat-boiler/server/config"
	"github.com/pkg/errors"
	"log"

	"go.mercari.io/datastore/boom"

	"cloud.google.com/go/datastore"
	"go.mercari.io/datastore/clouddatastore"

	w "go.mercari.io/datastore"
)

type DSFunc func(ctx context.Context) w.Client

func NewDSFactory(projectID string) DSFunc {
	return func(ctx context.Context) w.Client {
		cli, err := datastore.NewClient(ctx, projectID)
		if err != nil {
			panic(err)
		}

		client, err := clouddatastore.FromClient(ctx, cli)
		if err != nil {
			panic(err)
		}

		return client
	}
}

type Transaction func(ctx context.Context, fn func(tx *boom.Transaction) error) error

func NewTransaction(df DSFunc) Transaction {
	return func(ctx context.Context, fn func(tx *boom.Transaction) error) error {
		b := boom.FromClient(ctx, df(ctx))
		if _, err := b.RunInTransaction(fn); err != nil {
			return errors.WithStack(err)
		}
		return nil
	}
}

type Transactor interface {
	RunInTransaction(context.Context, func(tx *boom.Transaction) error,) error
}

type datastoreTransactor struct {
}

func NewDatastoreTransactor() Transactor {
	return &datastoreTransactor{}
}

// RunInTransaction represents datastore transaction
func (t *datastoreTransactor) RunInTransaction(ctx context.Context, fn func(tx *boom.Transaction) error) error {
	if _, err := FromContext(ctx).RunInTransaction(fn); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func FromContext(ctx context.Context) *boom.Boom {
	proj, err := config2.GetProject()

	cli, err := datastore.NewClient(ctx, "chat-boiler-t1")
	log.Println("proj", proj)
	if err != nil {
		log.Println("ds cli error", err)
		panic(err)
	}
	ds, err := clouddatastore.FromClient(ctx, cli)
	if err != nil {
		panic(err)
	}
	return boom.FromClient(ctx, ds)
}
