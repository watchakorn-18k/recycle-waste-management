package gateways

import "github.com/gofiber/fiber/v2"

func RouteUsers(gateway HTTPGateway, app *fiber.App) {
	api := app.Group("/api/user")

	api.Post("/add_user", gateway.CreateNewUserAccount)
	api.Get("/users", gateway.GetAllUserData)
	api.Put("/update_user", gateway.UpdateUserData)
	api.Delete("/delete_user/:user_id", gateway.DeleteUser)
	api.Get("/get_user", gateway.GetUser)
}

func RouteRecycle(gateway HTTPGateway, app *fiber.App) {
	api := app.Group("/api/recycle-waste")
	api.Get("/get-wastes", gateway.GetRecycleWaste)
	api.Post("/add-waste", gateway.AddRecycleWaste)
}

func RouteCategoryWaste(gateway HTTPGateway, app *fiber.App) {
	api := app.Group("/api/category-waste")
	api.Get("/get-category", gateway.GetCategoryWaste)
}
