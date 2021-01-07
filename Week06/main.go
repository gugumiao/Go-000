package main

import (
    "sync"
    "time"
)

const (
    Default = 10
)

type bucket struct {
    Value int64
}

type num struct {
    buckets map[int64]*bucket
    mutex   sync.RWMutex
    size    int64
}

func newRoll() *num {
    r := &num{
        buckets: make(map[int64]*bucket),
        mutex:   sync.RWMutex{},
        size:    Default,
    }
    return r
}

func (r *num) getCurrBucket() *bucket {
    now := time.Now().Unix()
    b := r.buckets[now]
    if b == nil {
        b = &bucket{}
        r.buckets[now] = b
    }
    return b
}

func (r* num) removeBucket() {
    b := time.Now().Unix() - r.size

    for timestamp := range r.buckets {
        if timestamp <= b {
            delete(r.buckets, timestamp)
        }
    }
}

func (r *num) increment(i int64) {
    if i == 0 {
        return
    }

    r.mutex.Lock()
    defer r.mutex.Unlock()

    b := r.getCurrBucket()
    b.Value += i
    r.removeBucket()
}

func (r *num) sum(now time.Time) (s int64) {
    r.mutex.RLock()
    defer r.mutex.RUnlock()

    for timestamp, bucket := range r.buckets {
        if timestamp >= now.Unix() - r.size {
            s += bucket.Value
        }
    }
    return s
}

func (r *num) avg(now time.Time) int64 {
    return r.sum(now) / r.size
}

