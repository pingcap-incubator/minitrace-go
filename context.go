// Copyright 2020 TiKV Project Authors. Licensed under Apache-2.0.

package minitrace

import (
    "context"
    "sync"
    "time"
)

type spanContext struct {
    parent context.Context

    tracingContext *tracingContext
    tracedSpans    *localSpans
    currentId      uint32
    currentGid     int64
}

type tracingKey struct{}

var activeTracingKey = tracingKey{}

func (s spanContext) Deadline() (deadline time.Time, ok bool) {
    return s.parent.Deadline()
}

func (s spanContext) Done() <-chan struct{} {
    return s.parent.Done()
}

func (s spanContext) Err() error {
    return s.parent.Err()
}

func (s spanContext) Value(key interface{}) interface{} {
    if key == activeTracingKey {
        return s
    } else {
        return s.parent.Value(key)
    }
}

// Represents a per goroutine buffer
type localSpans struct {
    spans        *bufferList
    createTimeNs uint64
    refCount     int
}

type tracingContext struct {
    maxId uint32

    mu             sync.Mutex
    collectedSpans []SpanSet
}
