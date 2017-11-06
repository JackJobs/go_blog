package controllers

import (
	"MyTest/app/models"
	//"fmt"
	"github.com/revel/revel"
	"strings"
)

type WBlog struct {
	App
}

func (c WBlog) Putup(blog *models.Blog) revel.Result {
	blog.Title = strings.TrimSpace(blog.Title)
	blog.Email = strings.TrimSpace(blog.Email)
	blog.Subject = strings.TrimSpace(blog.Subject)
	blog.Validator(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.WBlog)
	}
	dao, err := models.NewDao()
	if err != nil {
		c.Response.Status = 500
		return c.RenderError(err)
	}
	defer dao.Close()
	err = dao.CreateBlog(blog)
	println("create blog")
	//println(err)
	//fmt.Printf("%v", err)
	if err != nil {
		c.Response.Status = 500
		return c.RenderError(err)
	}
	//return c.Redirect(App.Index)  //报502错误
	return c.Redirect("/")
}
