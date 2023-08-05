package Handlers

import (
	"github.com/gofiber/fiber/v2"
	"survey/Database"
)

func DeleteSurvey(c *fiber.Ctx) error {
	UserToken := c.Cookies("Token")
	SurveyToken := c.Params("Token")
	User, err := Database.FindUserByToken(UserToken)
	if err != nil {
		return c.Status(404).SendString("Kullanıcı bulunamadı")
	}
	if User.Perm != "Admin" {
		return c.Status(401).SendString("Yetersiz yetki")
	}

	Database.DeleteSurvey(SurveyToken)
	c.JSON(fiber.Map{
		"message":"Anket Başarıyla silindi",
	})
	return nil
}