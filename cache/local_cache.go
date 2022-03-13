package cache

import goCache "github.com/patrickmn/go-cache"

type LocalCache goCache.Cache

const (
    NoExpiration      = goCache.NoExpiration
    DefaultExpiration = goCache.DefaultExpiration
)

func NewLocalCache() *LocalCache {
    return (*LocalCache)(goCache.New(DefaultExpiration, 0))
}
