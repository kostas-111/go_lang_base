package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func TestLruCache(t *testing.T) {
	t.Run("добавление значений в начало списка", func(t *testing.T) {
		t.Parallel()

		cache := base.NewLruCache(3)

		cache.Put("1", "one")
		cache.Put("2", "two")
		cache.Put("3", "three")

		assert.Equal(t, []string{"3", "2", "1"}, keys(cache))
		assert.Equal(t, "3", cache.Head.Key)
		assert.Equal(t, "1", cache.Tail.Key)
	})

	t.Run("удаление самого давно использованного значения", func(t *testing.T) {
		t.Parallel()

		cache := base.NewLruCache(2)

		cache.Put("1", "one")
		cache.Put("2", "two")
		cache.Put("3", "three")

		assert.Nil(t, cache.Get("1"))
		assert.Equal(t, []string{"3", "2"}, keys(cache))
	})

	t.Run("получение значения переносит его в начало списка", func(t *testing.T) {
		t.Parallel()

		cache := base.NewLruCache(3)
		cache.Put("1", "one")
		cache.Put("2", "two")
		cache.Put("3", "three")

		rsl := cache.Get("1")

		assert.NotNil(t, rsl)
		assert.Equal(t, "one", *rsl)
		assert.Equal(t, []string{"1", "3", "2"}, keys(cache))
		assert.Equal(t, "1", cache.Head.Key)
	})

	t.Run("получение отсутствующего ключа возвращает nil", func(t *testing.T) {
		t.Parallel()

		cache := base.NewLruCache(2)
		cache.Put("1", "one")

		assert.Nil(t, cache.Get("2"))
		assert.Equal(t, []string{"1"}, keys(cache))
	})

	t.Run("обновление существующего значения переносит его в начало списка", func(t *testing.T) {
		t.Parallel()

		cache := base.NewLruCache(2)
		cache.Put("1", "one")
		cache.Put("2", "two")

		cache.Put("1", "new one")

		rsl := cache.Get("1")
		assert.NotNil(t, rsl)
		assert.Equal(t, "new one", *rsl)
		assert.Equal(t, []string{"1", "2"}, keys(cache))
	})

	t.Run("добавление ничего не делает при нулевом размере", func(t *testing.T) {
		t.Parallel()

		cache := base.NewLruCache(0)

		cache.Put("1", "one")

		assert.Nil(t, cache.Get("1"))
		assert.Nil(t, cache.Head)
		assert.Nil(t, cache.Tail)
	})

	t.Run("сложная последовательность операций", func(t *testing.T) {
		t.Parallel()

		cache := base.NewLruCache(3)

		cache.Put("a", "1")
		cache.Put("b", "2")
		cache.Put("c", "3")
		// a -> b -> c

		cache.Get("a")
		// a -> c -> b

		cache.Put("d", "4")
		// d -> a -> c (b вытеснен)

		cache.Get("c")
		// c -> d -> a

		cache.Put("b", "new")
		// b -> c -> d

		assert.Equal(t, []string{"b", "c", "d"}, keys(cache))

		val := cache.Get("b")
		assert.NotNil(t, val)
		assert.Equal(t, "new", *val)

		val = cache.Get("c")
		assert.NotNil(t, val)
		assert.Equal(t, "3", *val)

		val = cache.Get("d")
		assert.NotNil(t, val)
		assert.Equal(t, "4", *val)

		assert.Nil(t, cache.Get("a"))
	})

	t.Run("обновление значения в середине списка", func(t *testing.T) {
		t.Parallel()

		cache := base.NewLruCache(3)
		cache.Put("1", "one")
		cache.Put("2", "two")
		cache.Put("3", "three")
		// 3 -> 2 -> 1

		cache.Put("2", "new two")
		// 2 -> 3 -> 1

		assert.Equal(t, []string{"2", "3", "1"}, keys(cache))
		assert.Equal(t, "new two", cache.Head.Value)
	})

	t.Run("многократное получение одного ключа", func(t *testing.T) {
		t.Parallel()

		cache := base.NewLruCache(3)
		cache.Put("1", "one")
		cache.Put("2", "two")
		cache.Put("3", "three")

		for i := 0; i < 5; i++ {
			cache.Get("1")
		}

		assert.Equal(t, []string{"1", "3", "2"}, keys(cache))
	})
}

func keys(cache *base.LruCache) []string {
	keys := make([]string, 0)

	for node := cache.Head; node != nil; node = node.Next {
		keys = append(keys, node.Key)
	}

	return keys
}
