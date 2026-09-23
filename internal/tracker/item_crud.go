package tracker

type Tracker struct {
  items []Item
}

func NewTracker() *Tracker {
  return &Tracker{}
}

func (t *Tracker) AddItem(item Item) (Item, error) {
  if t.FindIndexById(item.ID) != -1 {
    return Item{}, ErrAlreadyExists
  }
  t.items = append(t.items, item)
  return item, nil
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

func (t *Tracker) UpdateItem(index int, item Item) error {
  if index < 0 || index >= len(t.items) {
    return  ErrNotFound
  }
  t.items[index] = item
  return nil
}

func (t *Tracker) FindIndexById(id string) int {
	for i, item := range t.items {
		if item.ID == id {
			return i
		}
	}

	return -1
}