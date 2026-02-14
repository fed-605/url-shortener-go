package redis

import (
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Small unit test with mini redis
func TestInsertionAndRetrieval(t *testing.T) {

	miniRedis, err := miniredis.Run()
	require.NoError(t, err)
	defer miniRedis.Close()

	redClient := redis.NewClient(&redis.Options{
		Addr: miniRedis.Addr(),
	})

	red := &RedisService{
		redisClient: redClient,
	}

	url := "https://www.youtube.com"
	alias := "YT"

	err = red.SaveUrlMapping(url, alias)
	require.NoError(t, err)

	ans, err := red.RetrieveUrl(alias)
	require.NoError(t, err)

	assert.Equal(t, url, ans)
}
