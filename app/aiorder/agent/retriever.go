package agent

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cloudwego/eino-ext/components/retriever/redis"
	"github.com/cloudwego/eino/components/retriever"
	"github.com/cloudwego/eino/schema"
	redisCli "github.com/whlxbd/gomall/app/aiorder/biz/dal/redis"

	redispkg "github.com/whlxbd/gomall/app/aiorder/pkg/redis"
	redisDoc "github.com/redis/go-redis/v9"
)

func defaultRedisRetrieverConfig(ctx context.Context) (*redis.RetrieverConfig, error) {
	config := &redis.RetrieverConfig{
		Client:       redisCli.RedisClient,
		Index:        fmt.Sprintf("%s%s", redispkg.RedisPrefix, redispkg.IndexName),
		Dialect:      2,
		ReturnFields: []string{redispkg.ContentField, redispkg.MetadataField, redispkg.DistanceField},
		TopK:         8,
		VectorField:  redispkg.VectorField,
		DocumentConverter: func(ctx context.Context, doc redisDoc.Document) (*schema.Document, error) {
			resp := &schema.Document{
				ID:       doc.ID,
				Content:  "",
				MetaData: map[string]any{},
			}
			for field, val := range doc.Fields {
				if field == redispkg.ContentField {
					resp.Content = val
				} else if field == redispkg.MetadataField {
					resp.MetaData[field] = val
				} else if field == redispkg.DistanceField {
					distance, err := strconv.ParseFloat(val, 64)
					if err != nil {
						continue
					}
					resp.WithScore(1 - distance)
				}
			}

			return resp, nil
		},
	}

	emd, err := NewArkEmbedding(ctx)
	if err != nil {
		return nil, err
	}
	config.Embedding = emd
	return config, nil
}

func newRetriever(ctx context.Context) (rtr retriever.Retriever, err error) {
	// TODO Modify component configuration here.
	config, err := defaultRedisRetrieverConfig(ctx)
	if err != nil {
		return nil, err
	}
	rtr, err = redis.NewRetriever(ctx, config)
	if err != nil {
		return nil, err
	}
	return rtr, nil
}
