package ds

import (
	"context"
	"github.com/hayashiki/gql-chat/server/src/config"
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
	cli, err := datastore.NewClient(ctx, config.GetProject())
	log.Println("cli", cli)
	if err != nil {
		log.Println("cli", err)
		panic(err)
	}
	ds, err := clouddatastore.FromClient(ctx, cli)
	if err != nil {
		panic(err)
	}
	return boom.FromClient(ctx, ds)
}



////////////////////


