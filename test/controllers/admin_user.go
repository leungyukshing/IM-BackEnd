package controllers

import (
	"Test/models"
	"fmt"
)

func (c *MainController) Login() {
	c.TplName = "login.html"
}

func (c *MainController) LoginCheck() {
	var user models.AdminUser
	inputs := c.Input()
	user.UserName = inputs.Get("username")
	user.Password = inputs.Get("password")
	//fmt.Println(inputs)
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