package api

type ItemRequest struct {
  ID   string `json:"id"`
  Name string `json:"name"`
}

type GetItemsResponse struct {
  Items []ItemRequest `json:"items"`
}