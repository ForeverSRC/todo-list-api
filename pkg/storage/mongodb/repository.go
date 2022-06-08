package mongodb

import (
	itemlisting "github.com/ForeverSRC/todo-list-api/pkg/item/listing"
	"github.com/ForeverSRC/todo-list-api/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *Storage) InsertItem(item model.Item) error {
	ctx, cancel := defaultMongoContext()
	defer cancel()

	_, err := s.Item.InsertOne(ctx, item)
	if err != nil {
		return err
	}

	return nil

}

func (s *Storage) FetchItems(q *itemlisting.ItemListQuery) (itemlisting.ItemList, error) {
	skip := q.PageSize * (q.Page - 1)
	filter := bson.D{
		{Key: "uid", Value: q.Uid},
		{Key: "state", Value: q.State},
	}

	result := make(itemlisting.ItemList, 0, q.PageSize)

	ctx, cancel := defaultMongoContext()
	defer cancel()

	count, err := s.Item.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}

	if count < skip {
		return result, nil
	}

	opts := &options.FindOptions{
		Limit: &q.PageSize,
		Skip:  &skip,
		Sort:  bson.D{{Key: "create_time", Value: -1}},
	}

	cur, err := s.Item.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var item model.Item
		err = cur.Decode(&item)
		if err != nil {
			return nil, err
		}
		result = append(result, item)
	}

	return result, nil
}
