package Handlers

import (
	"github.com/gofiber/fiber/v2"
	"survey/Database"
)

func FindUser(c *fiber.Ctx) error {
	Token := c.Params("Token")
	User, err := Database.FindUserByToken(Token);
	if err != nil {
		return c.Status(404).SendString("Kullanıcı bulunamadı")
	}

	c.JSON(fiber.Map{
		"data":User,
	})
	return nil
}

func FindSurvey(c *fiber.Ctx) error {
	Token := c.Params("Token")
	Survey , err := Database.FindSurvey(Token)
	if err != nil {
		return c.Status(404).SendString("Anket bulunamadı")
	}

	c.JSON(fiber.Map{
		"data":Survey,
	})
	return nil
}

func FindOptions(c *fiber.Ctx) error {
	Token := c.Params("SurveyToken")
	Options , err := Database.FindOptions(Token)
	if err != nil {
		return c.Status(404).SendString("Seçenekler bulunamadı")
	}

	c.JSON(fiber.Map{
		"data":Options,
	})
	return nil
}

func ListSurveys(c *fiber.Ctx) error {
	Surveys, err := Database.ListSurveys();
	if err != nil {
		return c.Status(404).SendString("Ankketler bulunamadı")
	} 
	c.JSON(fiber.Map{
		"data":Surveys,
	})
	return nil
}