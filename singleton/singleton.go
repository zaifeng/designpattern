package singleton

import "sync"

type singleton struct{}

var single = new(singleton)

func GetInstance() *singleton {
	return single
}

var mu sync.Mutex

func GetInsLazy() *singleton {
	if single == nil {
		mu.Lock()
		defer mu.Unlock()
		if single == nil {
			single = new(singleton)
		}
	}
	return single
}

var once sync.Once

func GetInsOnce() *singleton {
	once.Do(func() {
		single = new(singleton)
	})
	return single
}
