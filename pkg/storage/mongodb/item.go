package mongodb

import (
	"context"

	itemlisting "github.com/ForeverSRC/todo-list-api/pkg/dto"
	"github.com/ForeverSRC/todo-list-api/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *Storage) InsertItem(ctx context.Context, item model.Item) error {
	mongoCtx, cancel := wrapContextWithTimeout(ctx)
	defer cancel()

	_, err := s.Item.InsertOne(mongoCtx, item)
	if err != nil {
		return err
	}

	return nil

}

func (s *Storage) FetchItems(ctx context.Context, query *itemlisting.ItemListQuery) (model.ItemList, error) {
	skip := query.PageSize * (query.Page - 1)
	filter := bson.D{
		{Key: "state", Value: query.State},
	}

	result := make(model.ItemList, 0, query.PageSize)

	mongoCtx, cancel := wrapContextWithTimeout(ctx)
	defer cancel()

	count, err := s.Item.CountDocuments(mongoCtx, filter)
	if err != nil {
		return nil, err
	}

	if count < skip {
		return result, nil
	}

	opts := &options.FindOptions{
		Limit: &query.PageSize,
		Skip:  &skip,
		Sort:  bson.D{{Key: "create_time", Value: -1}},
	}

	cur, err := s.Item.Find(mongoCtx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cur.Close(mongoCtx)

	for cur.Next(mongoCtx) {
		var item model.Item
		err = cur.Decode(&item)
		if err != nil {
			return nil, err
		}
		result = append(result, item)
	}

	return result, nil
}

func (s *Storage) GetItem(ctx context.Context, id string) (*model.Item, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	mongoCtx, cancel := wrapContextWithTimeout(ctx)
	defer cancel()

	filter := bson.M{
		"_id": oid,
	}

	var item model.Item
	err = s.Item.FindOne(mongoCtx, filter).Decode(&item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (s *Storage) UpdateItem(ctx context.Context, id string, item model.Item) error {
	item.Id = ""
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	mongoCtx, cancel := wrapContextWithTimeout(ctx)
	defer cancel()

	_, err = s.Item.UpdateByID(mongoCtx, oid, item)
	if err != nil {
		return err
	}

	return nil

}
