package main

import (
    "sync"
    "time"
)

const (
    Default = 10
)

type bucket struct {
    Value int
}

type numRoll struck {
    buckets map[int64]*bucket
    mutex   sync.RWMutex
}

