package tracker

type UI struct {
  In Input
  Out Output
  Tracker *Tracker
}

func (u UI) Run() {
  actions := map[string]UseCase{
    "add": AddUseCase{},
    "get": GetUseCase{},
    "delete": DeleteUseCase{},
    "find": FindUseCase{},
    "update": UpdateUseCase{},
  }

  for {
    u.Out.Out("select action")
    u.Out.Out("available: add, get, delete, find, update, exit")
    selected := u.In.Get()

    if selected == "exit" {
      break
    }

    action, ok := actions[selected]
    if !ok {
      u.Out.Out("not found action")
      continue
    }
    action.Done(u.In, u.Out, u.Tracker)
  }
}