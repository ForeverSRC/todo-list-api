package mongodb

import (
	"context"

	"github.com/ForeverSRC/todo-list-api/pkg/model"
	"github.com/ForeverSRC/todo-list-api/pkg/vo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *Storage) CreateMission(ctx context.Context, mission model.Mission) (string, error) {
	mongoCtx, cancel := wrapContextWithTimeout(ctx)
	defer cancel()

	res, err := s.Mission.InsertOne(mongoCtx, mission)
	if err != nil {
		return "", err
	}

	id := res.InsertedID.(primitive.ObjectID).Hex()
	return id, nil
}

func (s *Storage) GetMissions(ctx context.Context, query vo.MissionListQuery) (*model.MissionList, error) {
	total := query.Page * query.PageSize
	skip := total - query.PageSize

	result := make([]model.Mission, 0, query.PageSize)

	mongoCtx, cancel := wrapContextWithTimeout(ctx)
	defer cancel()

	filter := make(bson.D, 0)
	if query.State != nil {
		filter = append(filter, bson.E{Key: "state", Value: query.State})
	}

	count, err := s.Mission.CountDocuments(mongoCtx, filter)
	if err != nil {
		return nil, err
	}

	sort := make(bson.D, 0, 2)
	if query.Descending {
		sort = append(sort, bson.E{Key: "priority", Value: -1})
	} else {
		sort = append(sort, bson.E{Key: "priority", Value: 1})
	}

	sort = append(sort, bson.E{Key: "create_time", Value: -1})

	opts := &options.FindOptions{
		Limit: &query.PageSize,
		Skip:  &skip,
		Sort:  sort,
		Projection: bson.D{
			{Key: "_id", Value: 1},
			{Key: "title", Value: 1},
			{Key: "state", Value: 1},
			{Key: "priority", Value: 1},
		},
	}

	cur, err := s.Mission.Find(mongoCtx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cur.Close(mongoCtx)

	for cur.Next(mongoCtx) {
		var mission model.Mission
		err = cur.Decode(&mission)
		if err != nil {
			return nil, err
		}
		result = append(result, mission)
	}

	return &model.MissionList{
		Missions: result,
		NoMore:   count <= total,
	}, nil
}

func (s *Storage) GetMission(ctx context.Context, id string) (*model.Mission, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	mongoCtx, cancel := wrapContextWithTimeout(ctx)
	defer cancel()

	filter := bson.M{
		"_id": oid,
	}

	var mission model.Mission
	err = s.Mission.FindOne(mongoCtx, filter).Decode(&mission)
	if err != nil {
		return nil, err
	}

	return &mission, nil
}

func (s *Storage) UpdateMission(ctx context.Context, id string, mission model.Mission) error {
	mission.Id = ""
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	mongoCtx, cancel := wrapContextWithTimeout(ctx)
	defer cancel()

	_, err = s.Mission.UpdateByID(mongoCtx, oid, bson.M{"$set": mission})
	if err != nil {
		return err
	}

	return nil

}

func (s *Storage) AddItemToMission(ctx context.Context, mid string, itemId string) error {
	oid, err := primitive.ObjectIDFromHex(mid)
	if err != nil {
		return err
	}

	mongoCtx, cancel := wrapContextWithTimeout(ctx)
	defer cancel()

	_, err = s.Mission.UpdateByID(mongoCtx, oid, bson.M{"$push": bson.M{"items": itemId}})
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) DeleteItemFromMission(ctx context.Context, mid string, itemId string) error {
	oid, err := primitive.ObjectIDFromHex(mid)
	if err != nil {
		return err
	}

	mongoCtx, cancel := wrapContextWithTimeout(ctx)
	defer cancel()

	_, err = s.Mission.UpdateByID(mongoCtx, oid, bson.M{"$pull": bson.M{"items": itemId}})
	if err != nil {
		return err
	}

	return nil
}
