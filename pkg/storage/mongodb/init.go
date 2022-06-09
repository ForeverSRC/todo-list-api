package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	arrayutil "github.com/ForeverSRC/todo-list-api/pkg/utils/array"
)

var collections = []string{"item"}

type Storage struct {
	Context context.Context
	Client  *mongo.Client
	Cancel  context.CancelFunc
	db      *mongo.Database

	Item *mongo.Collection
}

func NewStorage(connStr string, db string) *Storage {
	store, err := connect(connStr, db)
	if err != nil {
		panic(err)
	}

	if err = initCollections(store); err != nil {
		panic(err)
	}

	store.Item = store.db.Collection("item")

	return store
}

func connect(connStr string, db string) (*Storage, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connStr))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		cancel()
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		cancel()
		return nil, err
	}

	return &Storage{
		Context: ctx,
		Client:  client,
		Cancel:  cancel,
		db:      client.Database(db),
	}, nil
}

func initCollections(storage *Storage) error {
	cs, err := storage.db.ListCollectionNames(storage.Context, bson.M{})
	if err != nil {
		return err
	}

	if cs == nil {
		cs = []string{}
	}
	for _, c := range collections {
		if !arrayutil.Contains(collections, c) {
			if err = storage.db.CreateCollection(storage.Context, c); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *Storage) Close() error {
	defer s.Cancel()
	return s.Client.Disconnect(s.Context)
}

func defaultMongoContext() (ctx context.Context, cancel context.CancelFunc) {
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	return
}

func wrapContextWithTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	if ctx == nil {
		return defaultMongoContext()
	}
	return context.WithTimeout(ctx, 10*time.Second)
}
