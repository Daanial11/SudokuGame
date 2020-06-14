package main

import (
	"sync"
)

type SafeBool struct {
	val bool
	m   sync.Mutex
}

//Locks when getting the value and unlocks after
func (i *SafeBool) Get() bool {
	i.m.Lock()
	defer i.m.Unlock()
	return i.val
}

//Locks when setting the value and unlocks after
func (i *SafeBool) Set(val bool) {
	i.m.Lock()
	defer i.m.Unlock()
	i.val = val
}

