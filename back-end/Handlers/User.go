package Handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"survey/Utils"
	"survey/Database"
)

func HelloWorld(c *fiber.Ctx) error {
	c.JSON(fiber.Map{
		"hello":"world",
	})
	return nil
}

func Register(c *fiber.Ctx) error {
	form := new(RegisterForm)
	if err := c.BodyParser(form); err != nil {
		return c.Status(400).SendString("Bad Request")
	}

	if !Utils.ValidateEmail(form.Email){
		return c.Status(400).SendString("Bu email geçersiz")
	}

	if !Utils.ValidatePassword(form.Pass) {
		return c.Status(400).SendString("Şifre yeterince güvenli değil")
	}

	if !Utils.ValidateUsername(form.Username) {
		return c.Status(400).SendString("Kullanıcı Adı Boşluk içermemeli")
	}

	if !Database.Register(form.Username, form.Email, form.Pass){
		return c.Status(400).SendString("Bu Email kullanılıyor.")
	}

	login, _ := Database.Login(form.Email,form.Pass)

	c.JSON(fiber.Map{
		"Token": login,
	})
	return nil
}

func Login(c *fiber.Ctx) error {
	form := new(LoginForm)
	if err := c.BodyParser(form); err != nil {
		return c.Status(400).SendString("Bad Request")
	}
	user := Database.FindUserByEmail(form.Email)
	if !user {
		return c.Status(404).SendString("Kullanıcı bulunamadı")
	}

	login, err := Database.Login(form.Email,form.Pass)
	if err != nil{
		return c.Status(400).SendString("E-mail yada şifre yanlış")
	}

	c.JSON(fiber.Map{
		"Token": login,
	})
	return nil
}

func CreateSurvey(c *fiber.Ctx) error {
	form := new(CreateSurveyForm)
	if err := c.BodyParser(form); err != nil {
		return c.Status(400).SendString("Bad Request")
	}

	Token := c.Cookies("Token")
	User, err := Database.FindUserByToken(Token)
	if err != nil {
		return c.Status(404).SendString("Kullanıcı bulunamadı")
	}
	fmt.Println(User)

	err = Database.CreateSurvey(form.Title,form.Description,form.Options)
	fmt.Println(form.Options)
	if err != nil {
		return c.Status(500).SendString("Anket oluşturulamadı")
	}
	
	c.JSON(fiber.Map{
		"message":"Anket oluşturuldu",
	})
	return nil
}

func VoteOption(c *fiber.Ctx) error {
	SurveyToken := c.Params("SurveyToken")
	OptionToken := c.Params("Token")
	UserToken := c.Cookies("Token")
	User, err := Database.FindUserByToken(UserToken)
	if err != nil {
		return c.Status(404).SendString("Kullanıcı bulunamadı")
	}
	fmt.Println(User)
	UserVotes := Database.FindVoteCount(UserToken,SurveyToken);
	if  UserVotes > 0{
		return c.Status(400).SendString("Zaten oy vermişsin")
	}

	Database.GiveVote(UserToken,SurveyToken,OptionToken)
	c.JSON(fiber.Map{
		"message":"Ankete oy verildi",
	})
	return nil
}