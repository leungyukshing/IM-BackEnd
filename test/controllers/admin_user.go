package controllers

import (
	"github.com/backend/test/models"
	"fmt"
)

func (c *MainController) Login() {
	c.TplName = "login.html"
}

func (c *MainController) LoginCheck() {
	//proto.unmarshal(c.Ctx.Input.RequestBody, req)
	var user models.AdminUser
	inputs := c.Input()
	user.Username = inputs.Get("username")
	user.Password = inputs.Get("password")
	fmt.Printf("input: %v", inputs)
	fmt.Printf("user: %v", user)
	err := user.ValidateUser()
	if err == nil {
		c.SetSession("name", user)
		c.Redirect("/",301)
	} else {
		fmt.Println(err)
		c.Data["info"] = err
		c.TplName = "error.html"
	}
}

func (c *MainController) Out() {
	c.DelSession("name")
	c.Redirect("/",301)
}