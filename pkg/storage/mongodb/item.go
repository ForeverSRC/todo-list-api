package mongodb

import (
	"context"

	"github.com/ForeverSRC/todo-list-api/pkg/model"
	"github.com/ForeverSRC/todo-list-api/pkg/vo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *Storage) InsertItem(ctx context.Context, item model.Item) (string, error) {
	mongoCtx, cancel := wrapContextWithTimeout(ctx)
	defer cancel()

	res, err := s.Item.InsertOne(mongoCtx, item)
	if err != nil {
		return "", err
	}

	id := res.InsertedID.(primitive.ObjectID).Hex()
	return id, nil

}

func (s *Storage) FetchItems(ctx context.Context, query *vo.ItemListQuery) (*model.ItemList, error) {
	total := query.PageSize * query.Page
	skip := total - query.PageSize

	result := make([]model.Item, 0, query.PageSize)

	mongoCtx, cancel := wrapContextWithTimeout(ctx)
	defer cancel()

	filter := make(bson.D, 0)
	if query.State != 0 {
		filter = append(filter, bson.E{Key: "state", Value: query.State})
	}

	count, err := s.Item.CountDocuments(mongoCtx, filter)
	if err != nil {
		return nil, err
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

	return &model.ItemList{
		Items:  result,
		NoMore: count <= total,
	}, nil
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

	_, err = s.Item.UpdateByID(mongoCtx, oid, bson.M{"$set": item})
	if err != nil {
		return err
	}

	return nil

}

func (s *Storage) DeleteItem(ctx context.Context, id string) (model.Item, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Item{}, err
	}

	filter := bson.M{"_id": oid}

	mongoCtx, cancel := wrapContextWithTimeout(ctx)
	defer cancel()

	res := s.Item.FindOneAndDelete(mongoCtx, filter)
	if res.Err() != nil {
		return model.Item{}, res.Err()
	}

	var item model.Item
	err = res.Decode(&item)
	if err != nil {
		return model.Item{}, err
	}

	return item, nil
}

func (s *Storage) FetchItemsByIds(ctx context.Context, ids []string) (map[string]model.Item, error) {
	numbers := len(ids)
	if numbers == 0 {
		return map[string]model.Item{}, nil
	}

	oids := make([]primitive.ObjectID, 0, numbers)

	for _, id := range ids {
		oid, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		oids = append(oids, oid)
	}

	filter := bson.M{"_id": bson.M{"$in": oids}}
	mongoCtx, cancel := wrapContextWithTimeout(ctx)
	defer cancel()

	cur, err := s.Item.Find(mongoCtx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(mongoCtx)

	items := make(map[string]model.Item, numbers)

	for cur.Next(mongoCtx) {
		var item model.Item
		err = cur.Decode(&item)
		if err != nil {
			return nil, err
		}
		items[item.Id] = item
	}

	return items, nil
}
