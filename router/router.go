package router

import (
	"et-practice/controller"
	"et-practice/model"
	"et-practice/response"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func UserRoute(app *fiber.App) {
	app.Post("/user", CreateUser)
	app.Get("/user/:userId", GetUser)
	app.Put("/user/:userId", UpdateUser)
	app.Patch("/user/:userId/:status", UpdateUserStatus)
	app.Delete("/user/:userId", DeleteUser)
}

func CreateUser(c *fiber.Ctx) error {
	var body controller.CreateUserBody

	//validate the request body
	if err := c.BodyParser(&body); err != nil {
		return c.Status(response.BodyParserError).JSON(response.UserResponse{Status: response.BodyParserError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&body); validationErr != nil {
		return c.Status(response.InvalidParameter).JSON(response.UserResponse{Status: response.InvalidParameter, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newUserId, err := controller.CreateUser(body)
	if err != nil {
		return c.Status(response.InvalidDbData).JSON(response.UserResponse{Status: response.InvalidDbData, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(response.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"id": newUserId}})
}

func GetUser(c *fiber.Ctx) error {
	userId := c.Params("userId")
	user, err := controller.GetUser(userId)

	if err != nil {
		return c.Status(response.InvalidDbData).JSON(response.UserResponse{Status: response.InvalidDbData, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(response.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"user": user}})
}

func UpdateUser(c *fiber.Ctx) error {
	userId := c.Params("userId")
	var user model.User

	//validate the request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(response.BodyParserError).JSON(response.UserResponse{Status: response.BodyParserError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(response.InvalidParameter).JSON(response.UserResponse{Status: response.InvalidParameter, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	var updatedUser model.User
	if err := controller.UpdateUser(userId, user, &updatedUser); err != nil {
		return c.Status(response.InvalidDbData).JSON(response.UserResponse{Status: response.InvalidDbData, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(response.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updatedUser}})
}

func UpdateUserStatus(c *fiber.Ctx) error {
	userId := c.Params("userId")
	status, _ := strconv.Atoi(c.Params("status"))

	var updatedUser model.User
	if err := controller.UpdateUserStatus(userId, status, &updatedUser); err != nil {
		return c.Status(response.InvalidDbData).JSON(response.UserResponse{Status: response.InvalidDbData, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(response.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updatedUser}})
}

func DeleteUser(c *fiber.Ctx) error {
	userId := c.Params("userId")

	if err := controller.DeleteUser(userId); err != nil {
		return c.Status(response.InvalidDbData).JSON(response.UserResponse{Status: response.InvalidDbData, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(
		response.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "User successfully deleted!"}},
	)
}
