package cache

import (
	lru "github.com/hashicorp/golang-lru"
	"github.com/unicornultrafoundation/go-hashgraph/native/idx"
)

const (
	maxValidatorsCacheSize = int(32)
)

type ValidatorCache struct {
	ValidatorCache *lru.Cache
}

type Validators struct {
	Count              uint64
	Key                []byte
	Indices            []idx.ValidatorID
	TotalStakedBalance uint64
}

func NewValidatorCache() *ValidatorCache {
	lruCache, _ := lru.New(maxValidatorsCacheSize)
	return &ValidatorCache{
		ValidatorCache: lruCache,
	}
}

func (c *ValidatorCache) Set(v *Validators) {
	c.ValidatorCache.Add(v.Key, v)
}

func (c *ValidatorCache) Get(key []byte) (*Validators, error) {
	obj, exists := c.ValidatorCache.Get(key)
	if !exists {
		return nil, nil
	}
	return obj.(*Validators), nil
}
