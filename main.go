package main

import (
	"fmt"
	"time"
	"github.com/yixincapital/gomemcache/memcache"
	"github.com/yixincapital/gomemcache/serialization"
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
func main() {
	address := []string{"192.168.145.6:11211", "192.168.145.5:11211"}
	memcache.Instance().SetClient(new(serialization.MsgpackProvider), 200 * time.Millisecond, address...)
	key := fmt.Sprintf("testkey_%v", 11)
	err := memcache.Instance().SetValue(key, "my value 1", 1000)
	fmt.Println(err)
	key2 := fmt.Sprintf("testkey_%v", 12)
	memcache.Instance().SetValue(key2, "my value 2", 1000)
	memcache.Instance().SetValue("hkh", nil, 1000)
	call := []ColorGroup{}
	err2 := memcache.Instance().GetValue("fsaf", &call)
	fmt.Println(err2)
	fmt.Println(call)
	for index := 0; index < 1; index++ {
		maps, err := memcache.Instance().GetMultiValue([]string{"test1", "test2", "test3"})
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(maps)
			fmt.Println(len(maps))
			for _, item := range maps {
				fmt.Println(string(item.Value))
			}
			fmt.Println("====================================")
		}
	}
}
