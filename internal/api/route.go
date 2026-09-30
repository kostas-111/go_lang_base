package api

import "github.com/gofiber/fiber/v2"

func (s *Server) Route(route fiber.Router) {
  route.Post("/item/", s.CreateItem)
  route.Delete("/item/:id", s.DeleteItem)
  route.Put("item/:id", s.UpdateItem)
  route.Get("/items/", s.GetItems)
  route.Get("/items/search", s.GetItemsByName)
}
