package main

import (
	"fmt"
	"job4j.ru/go-lang-base/internal/base"
)

func main() {
	fmt.Println("Hello World")

	  cache := base.NewLruCache(3)

    cache.Put("hello", "word")
    res := cache.Get("hello")

    fmt.Println(res)
}
