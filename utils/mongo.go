package utils

import (
	"github.com/chentihe/mongodb-api/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ToOpts(paginate *types.MongoPaginate) *options.FindOptions {
	limit := paginate.Limit
	skip := paginate.Page*paginate.Limit - paginate.Limit
	return &options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
	}
}

func ToDoc(v interface{}) (doc *bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return nil, err
	}

	if err = bson.Unmarshal(data, &doc); err != nil {
		return nil, err
	}

	return doc, nil
}
