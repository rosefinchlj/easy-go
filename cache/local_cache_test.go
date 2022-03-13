package cache

import (
    "testing"
    "time"
)

func TestLocalCacheGetSet(t *testing.T) {
    localCache := NewLocalCache()

    _, found := localCache.Get("a")
    if found {
        t.FailNow()
    }

    localCache.Set("a", "a_string", 5*time.Second)

    time.Sleep(2 * time.Second)
    value, found := localCache.Get("a")
    if !found && "a_string" != value.(string) {
        t.FailNow()
    }
    t.Log(value)

    time.Sleep(3 * time.Second)

    _, found = localCache.Get("a")
    if found {
        t.FailNow()
    }
}

func TestLocalCacheCommon(t *testing.T) {
    localCache := NewLocalCache()

    // 不过期
    localCache.Set("b", 1, DefaultExpiration)

    err := localCache.Increment("b", 1)
    if err != nil {
        t.FailNow()
    }

    value, found := localCache.Get("b")
    if !found && 2 != value.(int) {
        t.FailNow()
    }
    t.Log(value)

    err = localCache.Replace("b", 10, 2*time.Second)
    if err != nil {
        t.FailNow()
    }
    value, expire, found := localCache.GetWithExpiration("b")
    if !found && 10 != value.(int) && expire.Sub(time.Now()).Seconds() > 2 {
        t.FailNow()
    }
}
