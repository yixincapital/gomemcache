package memcache

import (
	"fmt"
	"testing"
	"time"

	"github.com/dropbox/godropbox/memcache"
	"github.com/reckcn/cache/serialization"
	msgpack "gopkg.in/vmihailenco/msgpack.v2"
)

type ColorGroup struct {
	Id     int
	Name   string
	Colors []string
}

var group = ColorGroup{
	Id:     1,
	Name:   "Reds",
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

func init() {
	address := []string{"192.168.145.6:11211", "192.168.145.5:11211"}
	Instance().SetClient(new(serialization.MsgpackProvider), 200 * time.Millisecond, address...)
}

func BenchmarkByGomemcacheSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("testkey2_%v", i)
		Instance().SetValue(key, group, 60 * 3)
	}
}

func BenchmarkByGomemcacheGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("testkey2_%v", i)
		Instance().GetValue(key, &group)
	}
}