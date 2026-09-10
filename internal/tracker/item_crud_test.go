package tracker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrackerGetItems(t *testing.T) {
  t.Parallel()
  t.Run("check link leak", func(t *testing.T) {
      t.Parallel()

      tracker := NewTracker()
      item := Item{
          ID:   "1",
          Name: "First Item",
      }
      tracker.AddItem(item)

      res := tracker.GetItems()
      res[0].Name = "Second Item"

      assert.Equal(t,
          []Item{item},
          tracker.GetItems(),
      )
  })

  t.Run("check add one", func(t *testing.T) {
      t.Parallel()
      tracker := NewTracker()
      item := Item{ID: "1", Name: "First"}
      tracker.AddItem(item)
      items := tracker.GetItems()
      assert.Len(t, items, 1)
      assert.Equal(t, item, items[0])
  })
}