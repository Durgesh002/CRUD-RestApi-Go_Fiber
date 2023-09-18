package handler

import (
	"fiber-gorm-restapi/database"
	"fiber-gorm-restapi/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllUsers(c *fiber.Ctx) error {
	db := database.DB.Db
	var users []models.User

	db.Find(&users)

	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Users not found", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User found", "data": users})
}

func GetSingleUser(c *fiber.Ctx) error {
	db := database.DB.Db

	id := c.Params("id")

	var user models.User

	//finding singe user
	db.Find(&user, "id = ?", id)

	if user.ID == uuid.Nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User found", "data": user})
}

func CreateUser(c *fiber.Ctx) error {
	db := database.DB.Db
	user := new(models.User)

	err := c.BodyParser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Enter valid input", "data": err})
	}

	err = db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "User created successfully", "data": user})
}

func DeleteUser(c *fiber.Ctx) error {
	db := database.DB.Db

	var user models.User

	id := c.Params("id")
	db.Find(&user, "id = ?", id)

	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found to delete", "data": nil})
	}

	err := db.Delete(&user, "id =  ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "failed to delete user", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User deleted successfully"})
}

func UpdateUser(c *fiber.Ctx) error {
	type UpdateUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	db := database.DB.Db

	var user models.User
	id := c.Params("id")

	db.Find(&user, "id = ?", id)
	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}

	var updateUserData UpdateUser
	err := c.BodyParser(&updateUserData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Enter valid input", "data": err})
	}
	user.Username = updateUserData.Username
	user.Password = updateUserData.Password
	user.Email = updateUserData.Email

	db.Save(&user)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User data updated successfully", "data": user})
}
