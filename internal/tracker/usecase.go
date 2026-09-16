package tracker

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type UseCase interface {
  Done(in Input, out Output, tracker *Tracker)
}

type AddUseCase struct{}

func (u AddUseCase) Done(in Input, out Output, tracker *Tracker) {
    out.Out("enter name:")
    name := in.Get()
    id := uuid.New().String()
    tracker.AddItem(Item{Name: name, ID: id})
}

type GetUseCase struct{}

func (u GetUseCase) Done(_ Input, out Output, tracker *Tracker) {
    for _, item := range tracker.GetItems() {
        out.Out(item.toString())
    }
}

type DeleteUseCase struct{}

func (u DeleteUseCase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter id for delete: ")
	id := in.Get()
	index := tracker.FindIndexById(id)

	if index == -1 {
		out.Out("item not found")
		return
	}

	deleted := tracker.DeleteItem(index)
	out.Out(fmt.Sprintf("items deleted: %s", deleted.toString()))
}

type FindUseCase struct{}

func (u FindUseCase) Done(in Input, out Output, tracker *Tracker) {
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
		out.Out("items didn't find")
	}
}

type UpdateUseCase struct{}

func (u UpdateUseCase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter item ID for update:")


	id := in.Get()
	index := tracker.FindIndexById(id)

	if index == -1 {
		out.Out("item not found")
		return
	}

	out.Out("enter new item name:")
	items := tracker.GetItems()
	item := items[index]
	item.Name = in.Get()
	tracker.UpdateItem(index, item)
	out.Out(fmt.Sprintf("item updated: %s", item.toString()))
}