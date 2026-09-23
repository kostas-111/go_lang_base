package tracker

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type UseCase interface {
  Done(in Input, out Output, tracker *Tracker) error
}

type AddUseCase struct{}

func (u AddUseCase) Done(in Input, out Output, tracker *Tracker) error {
    out.Out("enter name:")
    name := in.Get()
    id := uuid.New().String()
    tracker.AddItem(Item{Name: name, ID: id})
    return nil
}

type GetUseCase struct{}

func (u GetUseCase) Done(_ Input, out Output, tracker *Tracker) error {
    for _, item := range tracker.GetItems() {
        out.Out(item.toString())
    }
  return nil
}

type DeleteUseCase struct{}

func (u DeleteUseCase) Done(in Input, out Output, tracker *Tracker) error {
	out.Out("enter id for delete: ")
	id := in.Get()
	index := tracker.FindIndexById(id)

	if index == -1 {
		return ErrNotFound
	}

	deleted := tracker.DeleteItem(index)
	out.Out(fmt.Sprintf("items deleted: %s", deleted.toString()))
	return nil
}

type FindUseCase struct{}

func (u FindUseCase) Done(in Input, out Output, tracker *Tracker) error {
	out.Out("enter item name:")
	name := strings.ToLower(in.Get())
	found := false

	for _, item := range tracker.GetItems() {
		if strings.Contains(strings.ToLower(item.Name), name) {
			out.Out(fmt.Sprintf("found item: %s", item.toString()))
			found = true
		}
	}

	if !found {
		return ErrNotFound
	}
  return nil
}

type UpdateUseCase struct{}

func (u UpdateUseCase) Done(in Input, out Output, tracker *Tracker) error {
	out.Out("enter item ID for update:")


	id := in.Get()
	index := tracker.FindIndexById(id)

	if index == -1 {
		return ErrNotFound
	}

	out.Out("enter new item name:")
	items := tracker.GetItems()
	item := items[index]
	item.Name = in.Get()
	err := tracker.UpdateItem(index, item)
	if (err != nil) {
	  return fmt.Errorf("failed to update item: %w", err)
	}
	out.Out(fmt.Sprintf("item updated: %s", item.toString()))
	return nil
}