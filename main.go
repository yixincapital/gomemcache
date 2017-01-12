package main

import (
	"fmt"
	"time"
	"github.com/yixincapital/gomemcache/memcache"
	"github.com/yixincapital/gomemcache/serialization"
)

func main() {
	address := []string{"192.168.145.6:11212", "192.168.145.5:11211"}
	memcache.Instance().SetClient(new(serialization.MsgpackProvider), 200 * time.Millisecond, address...)
	//fmt.Println(instance.Client)
	key := fmt.Sprintf("testkey_%v", 11)
	err := memcache.Instance().SetValue(key, "my value 1", 1000)
	fmt.Println(err)
	key2 := fmt.Sprintf("testkey_%v", 12)
	memcache.Instance().SetValue(key2, "my value 2", 1000)
	memcache.Instance().SetValue("hkh", nil, 1000)
	for index := 0; index < 1; index++ {
		var str = ""
		//err := memcached.Instance().GetValue(key, &str)
		//mapd := make(map[string]string)
		maps, err := memcache.Instance().GetMultiValue([]string{key, key2, "hkh"})
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(str)
			fmt.Println(maps)
			for _, item := range maps {
				fmt.Println(string(item.Value))
			}
			fmt.Println("====================================")
		}
	}
}
