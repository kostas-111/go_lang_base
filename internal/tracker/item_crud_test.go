package tracker

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
      _, err := tracker.AddItem(item)
      require.NoError(t, err)

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
      _, err := tracker.AddItem(item)
      require.NoError(t, err)
      items := tracker.GetItems()
      assert.Len(t, items, 1)
      assert.Equal(t, item, items[0])
  })

  t.Run("error update - not found", func(t *testing.T) {
      t.Parallel()

      tracker := NewTracker()
      item := Item{
          ID:   "1",
          Name: "First Item",
      }

      err := tracker.UpdateItem(0, item)
      assert.ErrorIs(t, err, ErrNotFound)
  })

  t.Run("error add item - duplicate id", func(t *testing.T) {
          t.Parallel()

          tr := NewTracker()
          first := Item{ID: "1", Name: "First Item"}
          second := Item{ID: "1", Name: "Second Item"}

          _, err := tr.AddItem(first)
          assert.NoError(t, err)

          got, err := tr.AddItem(second)

          assert.ErrorIs(t, err, ErrAlreadyExists)
          assert.Equal(t, Item{}, got)

          items := tr.GetItems()
          assert.Len(t, items, 1)
          assert.Equal(t, "First Item", items[0].Name)
      })
}