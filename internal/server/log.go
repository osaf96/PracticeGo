package server

import (
	"fmt"
	"sync"
)

type Log struct {
	mu      sync.Mutex
	records []Record
}

// NewLog creates a new log
func NewLog() *Log {
	// & is used to create a pointer to the struct
	return &Log{}
}

func (c *Log) Append(record Record) (uint64, error) {
	// Lock the mutex
	c.mu.Lock()
	// Unlock the mutex when the function returns
	defer c.mu.Unlock()
	record.Offset = uint64(len(c.records))
	c.records = append(c.records, record)
	return record.Offset, nil
}

func (c *Log) Read(offset uint64) (Record, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if offset >= uint64(len(c.records)) {
		return Record{}, fmt.Errorf("offset out of range")
	}
	return c.records[offset], nil
}

type Record struct {
	Value  []byte
	Offset uint64
}

var ErrOffsetNotFound = fmt.Errorf("offset not found")
