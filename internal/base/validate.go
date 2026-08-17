package base

type ValidateRequest struct {
  UserId      string
  Title       string
  Description string
}

func Validate(req *ValidateRequest) []string {
  res := make([]string, 0)

  if req == nil {
    res = append(res, "nil pointer is used in ValidateRequest")
    return res
  }

  if req.UserId == "" {
    res = append(res, "UserId is required")
  }
  if req.Title == "" {
    res = append(res, "Title is required")
  }
  if req.Description == "" {
    res = append(res, "Description is required")
  }

  return res
}