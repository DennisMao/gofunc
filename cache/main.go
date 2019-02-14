package main

import (
	"goexperience/cache/hashicorp"
	"log"
)

func main() {
	cache, err := hashicorp.NewCache(hashicorp.LRU, 8)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 8; i++ {
		k := string(byte('1' + i))
		log.Printf("Set  key=%s value='ok' \n", k)
		cache.Add(string('1'+i), "ok")
	}

	result, ok := cache.Get("1")
	if !ok {
		panic("Incorrect result")
	}
	log.Printf("Get  key='1' value=%s \n", result)

}
