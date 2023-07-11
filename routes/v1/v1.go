package v1

import (
	"avito-intership-2022/view"

	"github.com/gofiber/fiber/v2"
)

type Router struct {
	App  *fiber.App
	View *view.View
}

type Route struct {
	Group *fiber.Router
	View  *view.View
}

func (r *Router) Routes() {
	v1 := r.App.Group("/v1")
	route := Route{Group: &v1, View: r.View}
	route.Balance()
}
