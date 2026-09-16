package tracker

type Tracker struct {
  items []Item
}

func NewTracker() *Tracker {
  return &Tracker{}
}

func (t *Tracker) AddItem(item Item) {
  t.items = append(t.items, item)
}

func (t *Tracker) GetItems() []Item {
  res := make([]Item, len(t.items))
  copy(res, t.items)
  return res
}

func (t *Tracker) DeleteItem(index int) Item {
    item := t.items[index]
    t.items = append(t.items[:index], t.items[index+1:]...)
    return item
}

func (t *Tracker) UpdateItem(index int, item Item) {
    t.items[index] = item
}

func (t *Tracker) FindIndexById(id string) int {
	for i, item := range t.items {
		if item.ID == id {
			return i
		}
	}

	return -1
}