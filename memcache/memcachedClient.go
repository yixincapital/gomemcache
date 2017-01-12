package memcache

import (
	"sync"
	"time"
	serialization "github.com/reckcn/cache/serialization"
)

var (
	_instance *MemcachedClient
	mutex sync.Mutex
)

type MemcachedClient struct {
	Client  *Client
	factory serialization.SerializationFactory
}

func Instance() *MemcachedClient {
	mutex.Lock()
	if _instance == nil {
		_instance = new(MemcachedClient)
	}
	defer mutex.Unlock()
	return _instance
}

func (self *MemcachedClient) SetClient(factory serialization.SerializationFactory, time time.Duration, server ...string) {
	mc := New(server...)
	mc.Timeout = time
	self.Client = mc
	self.factory = factory
}

func (self *MemcachedClient) SetValue(key string, v interface{}, expiration int32) error {
	data, err := self.factory.Marshal(v)
	if err != nil {
		return err
	}
	e := self.Client.Set(&Item{Key: key, Value: data, Expiration: expiration})
	return e
}

func (self *MemcachedClient) GetValue(key string, v ...interface{}) error {
	item, err := self.Client.Get(key)
	if err != nil {
		return err
	}
	e := self.factory.Unmarshal(item.Value, v...)
	return e
}

// func (self *MemcachedClient) GetMultiValue(keys []string, v1 interface{}, v2 map[string]interface{}) error {
// 	maps := v2
// 	items, err := self.Client.GetMulti(keys)
// 	if err != nil {
// 		return err
// 	}
// 	for _, item := range items {
// 		e := self.factory.Unmarshal(item.Value, v1)
// 		if e != nil {
// 			fmt.Println(e)
// 		}
// 		maps[item.Key] = v1
// 	}
// 	return nil
// }

func (self *MemcachedClient) GetMultiValue(keys []string) (map[string]*Item, error) {
	items, err := self.Client.GetMulti(keys)
	return items, err
}
