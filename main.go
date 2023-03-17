package main

import (
	"fmt"
	"time"

	"github.com/JunKaiChuang/go_training/pkg/localcache"
)

func main() {
	cache := localcache.New()

	cache.Set("hello", "world")

	fmt.Println(cache.Get("hello"))

	cache.Set("hello", "universe")

	fmt.Println(cache.Get("hello"))

	time.Sleep(2 * time.Second)

	fmt.Println(cache.Get("hello"))

	time.Sleep(30 * time.Second)

	fmt.Println(cache.Get("hello"))

	cache.Set("hello", "cat")

	fmt.Println(cache.Get("hello"))

	fmt.Println("end")
}
